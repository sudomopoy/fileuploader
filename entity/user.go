// entity/user.go
package entity

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    TelegramID    int64  `gorm:"uniqueIndex"`
    FirstName     string
    LastName      string
    Username      string
    IsAdmin       bool
    AdminPassword string
    IsBlocked     bool
}

// entity/file.go
package entity

import (
    "time"
    "gorm.io/gorm"
)

type File struct {
    gorm.Model
    FileID       string
    ChannelMsgID int
    UserID       uint
    Caption      string
    FileURL      string
}

// entity/channel.go
package entity

type Channel struct {
    ChannelID int64 `gorm:"primaryKey"`
}