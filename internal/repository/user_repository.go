// repository/user_repository.go
package repository

import (
	"github.com/sudomopoy/fileuploader/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindByTelegramID(telegramID int64) (*entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) FindByTelegramID(telegramID int64) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("telegram_id = ?", telegramID).First(&user).Error
	return &user, err
}

// Implement other repositories similarly...
