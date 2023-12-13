package pgsqlservice

import (
	"chatservice/model"
	"log"
)

func CreatePassword(password *model.Password, user *model.User) (*model.Password, error) {
	var id int
	err := pgdb.QueryRow("INSERT INTO spasswords (EncryptedData) VALUES ($1)", password.EncryptedData).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to create password: %v", err)
		return &model.Password{}, err
	}

	var id2 int
	err = pgdb.QueryRow("INSERT INTO suserpasswords (UserID, PasswordID) VALUES ($1, $2)", user.ID, id).Scan(&id2)
	if err != nil {
		return &model.Password{}, err
	}
	return &model.Password{ID: id, EncryptedData: password.EncryptedData}, nil
}

func GetPasswordByUserID(userID int) (*model.Password, error) {
	password := &model.Password{}
	err := pgdb.QueryRow("SELECT * FROM spasswords WHERE ID IN (SELECT PasswordID FROM suserpasswords WHERE UserID = $1)", userID).Scan(&password.ID, &password.EncryptedData)
	if err != nil {
		return nil, err
	}
	return password, nil
}

func UpdatePassword(password *model.Password) (*model.Password, error) {
	_, err := pgdb.Exec("UPDATE spasswords SET EncryptedData = $1 WHERE ID = $2", password.EncryptedData, password.ID)
	if err != nil {
		return &model.Password{}, err
	}
	return password, nil
}

func DeletePassword(password *model.Password) error {
	_, err := pgdb.Exec("DELETE FROM spasswords WHERE ID = $1", password.ID)
	if err != nil {
		return err
	}
	return nil
}
