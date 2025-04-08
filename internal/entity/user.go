// entity/user.go
package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID    int64 `gorm:"uniqueIndex"`
	FirstName     string
	LastName      string
	Username      string
	IsAdmin       bool
	AdminPassword string
	IsBlocked     bool
}

type File struct {
	gorm.Model
	FileID       string
	ChannelMsgID int
	UserID       uint
	Caption      string
	FileURL      string
}

type Channel struct {
	ChannelID int64 `gorm:"primaryKey"`
}
