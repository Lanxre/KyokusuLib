package sse

import (
	"log/slog"
	"sync"

	"github.com/lanxre/kyokusulib/internal/models/dto"
)

type Client chan dto.Notification

type NotificationHub struct {
	mu     sync.RWMutex
	clients map[int64]map[Client]struct{}
	log    *slog.Logger
}

func NewNotificationHub(log *slog.Logger) *NotificationHub {
	return &NotificationHub{
		clients: make(map[int64]map[Client]struct{}),
		log:    log,
	}
}

func (h *NotificationHub) Subscribe(userID int64) Client {
	h.mu.Lock()
	defer h.mu.Unlock()

	client := make(Client, 64)

	if _, exists := h.clients[userID]; !exists {
		h.clients[userID] = make(map[Client]struct{})
	}

	h.clients[userID][client] = struct{}{}

	return client
}

func (h *NotificationHub) Unsubscribe(userID int64, client Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	userClients, ok := h.clients[userID]
	if !ok {
		return
	}

	delete(userClients, client)
	close(client)

	if len(userClients) == 0 {
		delete(h.clients, userID)
	}
}

func (h *NotificationHub) Publish(
	userID int64,
	notification dto.Notification,
) {
	h.mu.RLock()
	clients, ok := h.clients[userID]
	if !ok {
		h.mu.RUnlock()
		return
	}

	snapshot := make([]Client, 0, len(clients))
	for client := range clients {
		snapshot = append(snapshot, client)
	}
	h.mu.RUnlock()

	for _, client := range snapshot {
		select {
		case client <- notification:
		default:
			h.log.Warn("notification buffer overflow",
				slog.Int64("user_id", userID),
				slog.Int64("notification_id", notification.ID),
			)
		}
	}
}

func (h *NotificationHub) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, clients := range h.clients {
		for client := range clients {
			close(client)
		}
	}
	h.clients = make(map[int64]map[Client]struct{})
}

func (h *NotificationHub) PublishMany(
	userIDs []int64,
	notification dto.Notification,
) {
	for _, userID := range userIDs {
		h.Publish(userID, notification)
	}
}