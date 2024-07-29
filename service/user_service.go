package service

import (
	"GoGinToDoList/constants"
	"GoGinToDoList/dto"
	"GoGinToDoList/entity"
	"GoGinToDoList/helpers"
	"GoGinToDoList/repository"
	"context"
	"fmt"
	"sync"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
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

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	check, flag, err := s.userRepo.CheckEmail(ctx, nil, req.Email)

	if err != nil || !flag {
		return dto.UserLoginResponse{}, dto.ErrEmailNotFound
	}

	checkPassword, err := helpers.CheckPassword(check.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	token := s.jwtService.GenerateToken(check.ID.String(), check.Role)

	return dto.UserLoginResponse{
		Token: token,
		Role:  check.Role,
	}, nil
}
