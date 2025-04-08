// cmd/bot/main.go
package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sudomopoy/fileuploader/database"
	"github.com/sudomopoy/fileuploader/internal/handler"
	"github.com/sudomopoy/fileuploader/internal/repository"
	"github.com/sudomopoy/fileuploader/internal/service"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	db, err := database.ConnectDB(os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Panic(err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	fileRepo := repository.NewFileRepository(db)
	channelRepo := repository.NewChannelRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	fileService := service.NewFileService(fileRepo)
	channelService := service.NewChannelService(channelRepo)

	// Create handler
	telegramHandler := handler.NewTelegramHandler(
		bot,
		userService,
		fileService,
		channelService,
	)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	telegramHandler.HandleUpdates(updates)
}
