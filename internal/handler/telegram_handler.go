// handler/telegram_handler.go
package handler

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sudomopoy/fileuploader/internal/entity"
	"github.com/sudomopoy/fileuploader/internal/service"
)

type TelegramHandler struct {
	bot         *tgbotapi.BotAPI
	userService service.UserService
}

func NewTelegramHandler(
	bot *tgbotapi.BotAPI,
	userService service.UserService,
) *TelegramHandler {
	return &TelegramHandler{
		bot:         bot,
		userService: userService,
	}
}

func (h *TelegramHandler) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		user, _ := h.userService.HandleUser(
			update.Message.From.ID,
			update.Message.From.FirstName,
			update.Message.From.LastName,
			update.Message.From.UserName,
		)

		if user.IsBlocked {
			continue
		}

		if update.Message.IsCommand() {
			h.handleCommand(update.Message, user)
			continue
		}

		if update.Message.Document != nil || update.Message.Photo != nil {
			h.handleFileUpload(update.Message, user)
			continue
		}

		if strings.HasPrefix(update.Message.Text, "https://t.me/yourbot?start=") {
			h.handleFileAccess(update.Message, user)
		}
	}
}

func (h *TelegramHandler) handleFileUpload(msg *tgbotapi.Message, user *entity.User) {
	// Handle file upload logic
	// Forward to private channel
	// Save to database
	// Send link to user
}

func (h *TelegramHandler) handleFileAccess(msg *tgbotapi.Message, user *entity.User) {
	// Extract file ID from link
	// Retrieve file from channel
	// Forward to user
}

func (h *TelegramHandler) handleCommand(msg *tgbotapi.Message, user *entity.User) {
	// switch msg.Command() {
	// case "setchannel":
	// 	h.handleSetChannelCommand(msg, user)
	// case "block":
	// 	h.handleBlockCommand(msg, user)
	// case "delete":
	// 	h.handleDeleteCommand(msg, user)
	// 	// Add other commands...
	// }
}

// Implement command handlers...
