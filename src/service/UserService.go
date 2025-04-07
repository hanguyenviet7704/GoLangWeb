package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"vietha/src/entity"
	"vietha/src/repository"
)

type UserService interface {
	LoginUser(email string, password string) bool
	RegisterUser(username string, email string, password string) (*entity.User, error)
	GetAllUsers() (*[]entity.User, error)
	GetUserById(id int) (*entity.User, error)
	CreateOrUpdateUser(user *entity.User, roleName []string) error
	DeleteUser(id int) error
	GetRoleFromUser(id int) (*[]entity.Roles, error)
	AssignRolesToUser(id int, roleNames []string) error
}
type userService struct {
	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewUserService(userRepo repository.UserRepository, db *gorm.DB) UserService {
	return &userService{
		userRepository: userRepo,
		db:             db,
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
	user := &entity.User{
		Name:      name,
		Email:     email,
		Password:  password,
		Is_active: true,
	}
	var userRole entity.Roles
	service.db.Where("name = ?", "ROLE_USER").First(&userRole)
	user.Roles = []entity.Roles{userRole}
	// Lưu user vào database
	err := service.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (service *userService) GetAllUsers() (*[]entity.User, error) {
	users, err := service.userRepository.FindAll()
	return users, err
}
func (service *userService) GetUserById(id int) (*entity.User, error) {
	user, err := service.userRepository.FindUserById(id)
	return user, err
}
func (service *userService) CreateOrUpdateUser(user *entity.User, roleNames []string) error {
	var roles []entity.Roles
	if err := service.db.Where("name IN ?", roleNames).Find(&roles).Error; err != nil {
		return fmt.Errorf("error querying roles: %w", err)
	}
	if len(roles) == 0 {
		return errors.New("Role Not Found")
	}
	user.Roles = roles
	var existingUser entity.User
	if err := service.db.First(&existingUser, user.Id).Error; err != nil {
		// Nếu không tìm thấy => tạo mới
		user.CreatedAt = time.Now()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := service.userRepository.Create(user); err != nil {
				return errors.New("Cannot Create User")
			}
		}
	} else {
		// User đã tồn tại => cập nhật
		user.CreatedAt = existingUser.CreatedAt
		if err := service.userRepository.Save(user); err != nil {
			return errors.New("Cannot update user")
		}
	}
	return nil
}
func (service *userService) DeleteUser(id int) error {
	return service.userRepository.Delete(&entity.User{Id: id})
}
func (service *userService) GetRoleFromUser(id int) (*[]entity.Roles, error) {
	_, err := service.userRepository.FindUserById(id)
	if err != nil {
		return nil, errors.New("User Not Found")
	}
	return service.userRepository.FindRoleFromUser(id)
}
func (service *userService) AssignRolesToUser(id int, roleNames []string) error {
	roles := []entity.Roles{}
	if err := service.db.Where("name IN ?", roleNames).Find(&roles).Error; err != nil {
		return fmt.Errorf("Can not find role in database")
	}
	err := service.userRepository.CreateRoleFromUser(id, &roles)
	if err != nil {
		return err
	}
	return nil
}
