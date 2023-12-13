package model

type StreamByUser struct {
	StreamID int64 `json:"stream_id"`
	Stream   struct {
		FilledChats []struct {
			Messages Message `json:"messages"`
		} `json:"filled_chats"`
	} `json:"stream"`
}
