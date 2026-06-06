package sse

import (
	"io"
	"net/http"
	"strconv"
)

type SSEWriter struct {
	w       http.ResponseWriter
	flusher http.Flusher
}

func NewSSEWriter(w http.ResponseWriter) (*SSEWriter, bool) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return nil, false
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	
	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	return &SSEWriter{w: w, flusher: flusher}, true
}

func (s *SSEWriter) SendEvent(event string, id int64, data []byte) error {
	if id > 0 {
		if _, err := io.WriteString(s.w, "id: "+strconv.FormatInt(id, 10)+"\n"); err != nil {
			return err
		}
	}
	if event != "" {
		if _, err := io.WriteString(s.w, "event: "+event+"\n"); err != nil {
			return err
		}
	}
	if _, err := s.w.Write([]byte("data: ")); err != nil {
		return err
	}
	if _, err := s.w.Write(data); err != nil {
		return err
	}
	if _, err := s.w.Write([]byte("\n\n")); err != nil {
		return err
	}
	s.flusher.Flush()
	return nil
}

func (s *SSEWriter) SendHeartbeat() error {
	if _, err := io.WriteString(s.w, ":heartbeat\n\n"); err != nil {
		return err
	}
	s.flusher.Flush()
	return nil
}
