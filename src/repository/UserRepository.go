package repository

import (
	"errors"
	"gorm.io/gorm"
	"vietha/src/entity"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		connection: db,
	}
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.connection.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("User không tồn tại")
	}
	return &user, err
}

func (r *userRepository) Save(user *entity.User) error {
	return r.connection.Create(user).Error
}
