package service

import (
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserById(userId int) (*dto.GetUserDTO, error) {
	userDb, err := s.Repo.GetByID(userId)

	if err != nil {
		return nil, err
	}

	return  &dto.GetUserDTO{
		ID:       userDb.ID,
		Name:     userDb.Name,
		Picture:  userDb.Picture,
		Role:     userDb.Role,
		Status:   userDb.Status,
		About:    userDb.About,
		Birthday: userDb.Birthday,
		Gender:   string(userDb.Gender),
		IsPublic: userDb.IsPublic,
		Banner:   userDb.Banner,

	}, err
}