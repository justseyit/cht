package model

type Password struct {
	ID            int    `json:"id"`
	EncryptedData string `json:"encrypted_data"`
}
