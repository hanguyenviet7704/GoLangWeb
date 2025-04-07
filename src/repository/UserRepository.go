package repository

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"vietha/src/entity"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	FindAll() (*[]entity.User, error)
	FindUserById(id int) (*entity.User, error)
	Save(user *entity.User) error
	Create(user *entity.User) error
	Delete(user *entity.User) error
	FindRoleFromUser(id int) (*[]entity.Roles, error)
	CreateRoleFromUser(id int, roles *[]entity.Roles) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("User không tồn tại")
	}
	return &user, err
}
func (r *userRepository) FindAll() (*[]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("Roles").Find(&users).Error
	return &users, err
}
func (r *userRepository) FindUserById(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}
func (r *userRepository) Save(user *entity.User) error {
	err := r.db.Save(user).Error
	return err
}
func (r *userRepository) Create(user *entity.User) error {
	err := r.db.Create(user).Error
	return err
}
func (r *userRepository) Delete(user *entity.User) error {
	if err := r.db.Model(user).Association("Roles").Clear(); err != nil {
		return err
	}
	if err := r.db.Model(user).Association("Permissions").Clear(); err != nil {
		return err
	}
	if err := r.db.Where("user_id = ?", user.Id).Delete(&entity.Tokens{}).Error; err != nil {
		return err
	}
	err := r.db.Delete(user).Error
	return err
}
func (r *userRepository) FindRoleFromUser(id int) (*[]entity.Roles, error) {
	var user entity.User
	err := r.db.Preload("Roles").Find(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user.Roles, nil
}

func (r *userRepository) CreateRoleFromUser(id int, roles *[]entity.Roles) error {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return errors.New("Can't find user with id " + strconv.Itoa(id))
	}
	err = r.db.Model(&user).Association("Roles").Append(roles)
	if err != nil {
		return errors.New("Can't create role for user with id " + strconv.Itoa(id))
	}
	return nil
}
