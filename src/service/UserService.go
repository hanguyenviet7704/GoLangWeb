package service

import (
	"errors"
	"vietha/src/entity"
	"vietha/src/repository"
)

type UserService interface {
	LoginUser(email string, password string) bool
	RegisterUser(username string, email string, password string) (*entity.User, error)
}
type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
func (service *userService) LoginUser(email string, password string) bool {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return false
	}
	if user.Password != password {
		return false
	}
	return true
}
func (service *userService) RegisterUser(name, email, password string) (*entity.User, error) {
	// Kiểm tra email đã tồn tại chưa
	existingUser, _ := service.userRepository.FindByEmail(email)
	if existingUser != nil { // Nếu user tồn tại thì báo lỗi
		return nil, errors.New("Email đã tồn tại")
	}
	// Tạo user mới
	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	// Lưu user vào database
	err := service.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
