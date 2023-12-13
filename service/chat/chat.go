package chatservice

import (
	"chatservice/model"
	pgsqlservice "chatservice/service/pgsql"
)

func CreateChat(chat *model.Chat) (*model.Chat, error) {
	return pgsqlservice.CreateChat(*chat)
}

func UpdateChat(chat *model.Chat) (*model.Chat, error) {
	return pgsqlservice.UpdateChat(*chat)
}

func GetChatByID(id int) (*model.Chat, error) {
	return pgsqlservice.GetChatByID(id)
}

func GetChatsByUserID(id int) ([]*model.Chat, error) {
	return pgsqlservice.GetChatsByUserID(id)
}

func AddUserToChat(userInChat *model.UserInChat) (*model.UserInChat, error) {
	return pgsqlservice.AddUserToChat(*userInChat)
}

func RemoveUserFromChat(userInChatID *int) (*model.UserInChat, error) {
	return pgsqlservice.RemoveUserFromChat(*userInChatID)
}

func MakeUserAdmin(userInChatID *int) (*model.UserInChat, error) {
	return pgsqlservice.MakeUserAdmin(*userInChatID)
}

func RemoveUserAdmin(userInChatID *int) (*model.UserInChat, error) {
	return pgsqlservice.RemoveUserAdmin(*userInChatID)
}

func GetUsersInChat(chatID int) ([]*model.User, error) {
	return pgsqlservice.GetUsersByChatID(chatID)
}

func ReactToMessage(messageReaction *model.ReactionToMessage) (*model.ReactionToMessage, error) {
	return pgsqlservice.ReactMessage(*messageReaction)
}

func UnreactMessage(messageReaction *model.ReactionToMessage) (*model.ReactionToMessage, error) {
	return pgsqlservice.UnreactMessage(*messageReaction)
}

func GetReactionsByMessageID(messageID int) ([]*model.ReactionToMessage, error) {
	return pgsqlservice.GetReactionsByMessageID(messageID)
}

func GetUsersByReactedUsersByMessage(messageID int) ([]*model.User, error) {
	return pgsqlservice.GetUsersByReactionToMessage(messageID)
}

func SendMessageToChat(chatID int, message string) (*model.Message, error) {
	return pgsqlservice.SendMessageToChat(chatID, message)
}

func GetMessagesByChatID(chatID int) ([]*model.Message, error) {
	return pgsqlservice.GetMessagesByChatID(chatID)
}
