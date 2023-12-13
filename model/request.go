package model

type RequestUserRegister struct {
	UserName            string `json:"user_name"`
	UserEmail           string `json:"user_email"`
	UserPassword        string `json:"user_password"`
	UserPasswordConfirm string `json:"user_password_confirm"`
}

type RequestUserLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type RequestUserLogout struct {
	UserID int `json:"user_id"`
}

type RequestSendMessage struct {
	Message Message `json:"message"`
	UserID  int     `json:"user_id"`
	ChatID  int     `json:"chat_id"`
}

type RequestCreateChat struct {
	Chat      Chat `json:"chat"`
	CreatorID int  `json:"creator_id"`
}

type RequestAddUserToChat struct {
	UserID int `json:"user_id"`
	ChatID int `json:"chat_id"`
}

type RequestReactMessage struct {
	MessageID         int               `json:"message_id"`
	ReactionToMessage ReactionToMessage `json:"reaction_to_message"`
}

type RequestGetMessages struct {
	ChatID int `json:"chat_id"`
}

type RequestGetChats struct {
	UserID int `json:"user_id"`
}

type RequestGetUsers struct {
	ChatID int `json:"chat_id"`
}

type RequestRemoveUserFromChat struct {
	UserInChatID int `json:"user_in_chat_id"`
}

type RequestMakeUserAdmin struct {
	UserInChatID int `json:"user_in_chat_id"`
}

type RequestRemoveUserAdmin struct {
	UserInChatID int `json:"user_in_chat_id"`
}

type RequestUpdateChat struct {
	Chat Chat `json:"chat"`
}
