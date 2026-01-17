package dto

type CreateTeamDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Slug        string `json:"slug" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=500"`
}

type UpdateTeamDTO struct {
	Description *string `json:"description"`
	AvatarURL   *string `json:"avatar_url"`
}