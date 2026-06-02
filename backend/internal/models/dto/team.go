package dto

type CreateTeamDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Slug        string `json:"slug" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=500"`
	TeamType    string `json:"team_type" validate:"omitempty,oneof=open private"`
}

type UpdateTeamDTO struct {
	Description        *string `json:"description"`
	AvatarURL          *string `json:"avatar_url"`
	BannerURL          *string `json:"banner_url"`
	TeamType           *string `json:"team_type" validate:"omitempty,oneof=open private"`
	OwnerRoleName      *string `json:"owner_role_name"`
	ModeratorRoleName  *string `json:"moderator_role_name"`
	MemberRoleName     *string `json:"member_role_name"`
}

type UpdateTeamMemberDTO struct {
	Role           *string `json:"role" validate:"omitempty,oneof=owner moderator member"`
	CustomRoleName *string `json:"custom_role_name" validate:"omitempty,max=100"`
}

type TeamStats struct {
	MemberCount      int `json:"member_count"`
	SubscribersCount int `json:"subscribers_count"`
}

type TeamRoleNames struct {
	Owner     string `json:"owner"`
	Moderator string `json:"moderator"`
	Member    string `json:"member"`
}

type TeamDTO struct {
	ID          int           `json:"id"`
	OwnerID     int           `json:"owner_id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	AvatarURL   string        `json:"avatar_url"`
	BannerURL   string        `json:"banner_url"`
	TeamType    string        `json:"team_type"`
	RoleNames   TeamRoleNames `json:"role_names"`
	CreatedAt   string        `json:"created_at"`
	Stats       TeamStats     `json:"stats"`
	IsMember    bool          `json:"is_member"`
	IsSubscriber bool          `json:"is_subscriber"`
	HasRequested bool          `json:"has_requested"`
}

type TeamMemberUserDTO struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Picture   string       `json:"picture"`
	ActiveTag string       `json:"active_tag"`
	UserLevel UserLevelDTO `json:"user_level"`
}

type TeamMemberDTO struct {
	User     TeamMemberUserDTO `json:"user"`
	Role     string            `json:"role"`
	RoleName string            `json:"role_name"`
	JoinedAt string            `json:"joined_at"`
}

type TeamSubscriberDTO struct {
	User         TeamMemberUserDTO `json:"user"`
	SubscribedAt string            `json:"subscribed_at"`
}

type TeamJoinRequestDTO struct {
	User      TeamMemberUserDTO `json:"user"`
	CreatedAt string            `json:"created_at"`
}
