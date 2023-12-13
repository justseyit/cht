package model

type Message struct {
	ID          int    `json:"id"`
	MessageText string `json:"message_text"`
	QuotedMsgID int    `json:"quoted_msg_id"`
}
