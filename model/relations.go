package model

type MessageInChat struct {
	ID        int `json:"id"`
	MessageID int `json:"message_id"`
	ChatID    int `json:"chat_id"`
}

type MessageOfUser struct {
	ID        int `json:"id"`
	MessageID int `json:"message_id"`
	UserID    int `json:"user_id"`
}

type UserInChat struct {
	ID     int `json:"id"`
	ChatID int `json:"chat_id"`
	UserID int `json:"user_id"`
}

type ReactionToMessage struct {
	ID            int    `json:"id"`
	MessageID     int    `json:"message_id"`
	ReactedUserID int    `json:"reacted_user_id"`
	Reaction      string `json:"reaction"`
}

type ChatAdmin struct {
	ID           int `json:"id"`
	UserInChatID int `json:"user_in_chat_id"`
}

type PasswordUser struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	PasswordID int `json:"password_id"`
}
