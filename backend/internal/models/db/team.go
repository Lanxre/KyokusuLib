package db

import "time"

type PublisherTeam struct {
	ID                int       `json:"id"`
	OwnerID           int       `json:"owner_id"`
	Name              string    `json:"name"`
	Slug              string    `json:"slug"`
	Description       string    `json:"description"`
	AvatarURL         string    `json:"avatar_url"`
	BannerURL         string    `json:"banner_url"`
	OwnerRoleName     string    `json:"owner_role_name"`
	ModeratorRoleName string    `json:"moderator_role_name"`
	MemberRoleName    string    `json:"member_role_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	MemberCount       int       `json:"member_count"`
	SubscribersCount  int       `json:"subscribers_count"`
	IsMember          bool      `json:"is_member"`
	IsSubscriber      bool      `json:"is_subscriber"`
}

type TeamMemberUser struct {
	User       User      `json:"user"`
	Role       string    `json:"role"`
	CustomRole string    `json:"custom_role"`
	JoinedAt   time.Time `json:"joined_at"`
}

type TeamSubscriberUser struct {
	User         User      `json:"user"`
	SubscribedAt time.Time `json:"subscribed_at"`
}
