package service

import (
	"GoGinToDoList/constants"
	"GoGinToDoList/dto"
	"GoGinToDoList/entity"
	"GoGinToDoList/repository"
	"context"
	"fmt"
	"sync"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
	}

	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

var (
	mu sync.Mutex
)

func (s *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	fmt.Print(flag)
	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		Email:       req.Email,
		Role:        constants.ENUM_ROLE_USER,
	}
	fmt.Print(user)

	userReq, err := s.userRepo.RegisterUser(ctx, nil, user)

	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:          userReq.ID.String(),
		Name:        userReq.Name,
		PhoneNumber: userReq.PhoneNumber,
		Role:        userReq.Role,
		Email:       userReq.Email,
	}, nil

}
