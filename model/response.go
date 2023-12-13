package model

type ResponseUserRegister struct {
	User    User   `json:"user"`
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}

type ResponseUserLogin struct {
	User    User   `json:"user"`
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}

type ResponseUserLogout struct {
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}

type ResponseGetUsers struct {
	Users   []User `json:"users"`
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}

type ResponseGetChats struct {
	Chats   []Chat `json:"chats"`
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}

type ResponseGetMessages struct {
	Messages []Message `json:"messages"`
	Message  string    `json:"message"`
	Succes   bool      `json:"succes"`
}

type ResponseSendMessage struct {
	MessageSent Message `json:"message_sent"`
	Message     string  `json:"message"`
	Succes      bool    `json:"succes"`
}

type ResponseCreateChat struct {
	Chat    Chat   `json:"chat"`
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}

type ResponseAddUserToChat struct {
	UserInChat UserInChat `json:"user_in_chat"`
	Message    string     `json:"message"`
	Succes     bool       `json:"succes"`
}

type ResponseReactMessage struct {
	ReactionToMessage ReactionToMessage `json:"reaction_to_message"`
	Message           string            `json:"message"`
	Succes            bool              `json:"succes"`
}

type ResponseRemoveUserFromChat struct {
	UserInChat UserInChat `json:"user_in_chat"`
	Message    string     `json:"message"`
	Succes     bool       `json:"succes"`
}

type ResponseMakeUserAdmin struct {
	UserInChat UserInChat `json:"user_in_chat"`
	Message    string     `json:"message"`
	Succes     bool       `json:"succes"`
}

type ResponseRemoveUserAdmin struct {
	UserInChat UserInChat `json:"user_in_chat"`
	Message    string     `json:"message"`
	Succes     bool       `json:"succes"`
}

type ResponseUpdateChat struct {
	Chat    Chat   `json:"chat"`
	Message string `json:"message"`
	Succes  bool   `json:"succes"`
}
