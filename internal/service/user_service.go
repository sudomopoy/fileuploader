// service/user_service.go
package service

import (
	"github.com/sudomopoy/fileuploader/internal/entity"
	"github.com/sudomopoy/fileuploader/internal/repository"
)

type UserService interface {
	HandleUser(telegramID int64, firstName, lastName, username string) (*entity.User, error)
	// SetAdmin(telegramID int64, password string) error
	// BlockUser(telegramID int64) error
	// UnblockUser(telegramID int64) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) HandleUser(telegramID int64, firstName, lastName, username string) (*entity.User, error) {
	user, err := s.userRepo.FindByTelegramID(telegramID)

	if err == nil {
		// Update user info if changed
		if user.FirstName != firstName || user.LastName != lastName || user.Username != username {
			user.FirstName = firstName
			user.LastName = lastName
			user.Username = username
			return user, s.userRepo.Update(user)
		}
		return user, nil
	}

	// Create new user
	newUser := &entity.User{
		TelegramID: telegramID,
		FirstName:  firstName,
		LastName:   lastName,
		Username:   username,
	}
	return newUser, s.userRepo.Create(newUser)
}

// Implement other service methods...
