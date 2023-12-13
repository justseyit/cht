package pgsqlservice

import (
	"chatservice/model"
	"log"
)

func CreateUser(user *model.User) (*model.User, error) {
	var id int
	err := pgdb.QueryRow("INSERT INTO suser(UserName, UserEmail) VALUES($1, $2) RETURNING ID", user.UserName, user.UserEmail).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
		return &model.User{}, err
	}
	return &model.User{ID: id, UserName: user.UserName, UserEmail: user.UserEmail}, nil
}

func GetUserByID(id int) (*model.User, error) {
	user := &model.User{}
	err := pgdb.QueryRow("SELECT * FROM suser WHERE ID = $1", id).Scan(&user.ID, &user.UserName, &user.UserEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := pgdb.QueryRow("SELECT * FROM suser WHERE UserEmail = $1", email).Scan(&user.ID, &user.UserName, &user.UserEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}
