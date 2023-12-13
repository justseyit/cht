package repository

import (
	models "chatservice/model"
)

var Values = struct {
	CurrentUser models.User `json:"user"`
}{}
