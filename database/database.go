// database/database.go
package database

import (
	"github.com/sudomopoy/fileuploader/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate entities
	err = db.AutoMigrate(
		&entity.User{},
		&entity.File{},
		&entity.Channel{},
	)

	return db, err
}
