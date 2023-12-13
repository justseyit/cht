package authservice

import (
	"chatservice/model"
	pgsqlservice "chatservice/service/pgsql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("SERVER_SECRET_KEY")

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func UserLogin(email string, password string) (*model.User, error) {
	user, err := pgsqlservice.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	passwordFromDB, err := pgsqlservice.GetPasswordByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB.EncryptedData), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserRegister(user *model.User, password string) (*model.User, error) {
	user, err := pgsqlservice.CreateUser(user)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
		return nil, err
	}
	passwordHash := getHash([]byte(password))
	_, err = pgsqlservice.CreatePassword(&model.Password{EncryptedData: passwordHash}, user)
	if err != nil {
		log.Fatalf("Failed to create password: %v", err)
		return nil, err
	}
	return user, nil
}
