package services

import "example.com/m/internal/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetInfoUserService() string {
	return us.userRepository.GetInfoUser()
}
