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

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Create handler
	telegramHandler := handler.NewTelegramHandler(
		bot,
		userService,
	)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	telegramHandler.HandleUpdates(updates)
}
