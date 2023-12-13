package model

type Chat struct {
	ID              int    `json:"id"`
	ChatName        string `json:"chat_name"`
	ChatLanguage    string `json:"chat_language"`
	ChatDescription string `json:"chat_description"`
	ChatImage       string `json:"chat_image"`
	ChatCreatedAt   string `json:"chat_created_at"`
}
