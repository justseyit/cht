package pgsqlservice

import "chatservice/model"

func SendMessageToChat(chatID int, message string) (*model.Message, error) {
	var id int
	err := pgdb.QueryRow("INSERT INTO smessage (MessageText, QuotedMessageID) VALUES ($1, $2)", chatID, message).Scan(&id)
	if err != nil {
		return &model.Message{}, err
	}

	err2 := pgdb.QueryRow(("INSERT INTO smessageinchat (ChatID, MessageID) VALUES ($1, $2)"), chatID, id).Err()
	if err2 != nil {
		return &model.Message{}, err2
	}

	return &model.Message{ID: id, MessageText: message, QuotedMsgID: chatID}, nil
}

func GetMessageByID(id int) (*model.Message, error) {
	message := &model.Message{}
	err := pgdb.QueryRow("SELECT * FROM smessage WHERE ID = $1", id).Scan(&message.ID, &message.MessageText, &message.QuotedMsgID)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func GetMessagesByChatID(chatID int) ([]*model.Message, error) {
	var messages []*model.Message
	rows, err := pgdb.Query("SELECT * FROM smessage WHERE ID IN (SELECT MessageID FROM smessageinchat WHERE ChatID = $1)", chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		message := &model.Message{}
		err := rows.Scan(&message.ID, &message.MessageText, &message.QuotedMsgID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func GetChatsByUserID(userID int) ([]*model.Chat, error) {
	var chats []*model.Chat
	rows, err := pgdb.Query("SELECT * FROM schat WHERE ID IN (SELECT ChatID FROM suserinchat WHERE UserID = $1)", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		chat := &model.Chat{}
		err := rows.Scan(&chat.ID, &chat.ChatName)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func ReactMessage(reaction model.ReactionToMessage) (*model.ReactionToMessage, error) {
	var id int
	err := pgdb.QueryRow("INSERT INTO sreactiontomessage (MessageID, ReactedUserID, Reaction) VALUES ($1, $2, $3)", reaction.MessageID, reaction.ReactedUserID, reaction.Reaction).Scan(&id)
	if err != nil {
		return &model.ReactionToMessage{}, err
	}
	return &model.ReactionToMessage{ID: id, MessageID: reaction.MessageID, ReactedUserID: reaction.ReactedUserID, Reaction: reaction.Reaction}, nil
}

func UnreactMessage(reaction model.ReactionToMessage) (*model.ReactionToMessage, error) {
	var id int
	err := pgdb.QueryRow("DELETE FROM sreactiontomessage WHERE MessageID = $1 AND ReactedUserID = $2 AND Reaction = $3", reaction.MessageID, reaction.ReactedUserID, reaction.Reaction).Scan(&id)
	if err != nil {
		return &model.ReactionToMessage{}, err
	}
	return &model.ReactionToMessage{ID: id, MessageID: reaction.MessageID, ReactedUserID: reaction.ReactedUserID, Reaction: reaction.Reaction}, nil
}

func GetUsersByReactionToMessage(messageID int) ([]*model.User, error) {
	var users []*model.User
	rows, err := pgdb.Query("SELECT * FROM suser WHERE ID IN (SELECT ReactedUserID FROM sreactiontomessage WHERE MessageID = $1)", messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.UserName, &user.UserEmail)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func MakeUserAdmin(userInChat int) (*model.UserInChat, error) {
	var userInChatt model.UserInChat
	err0 := pgdb.QueryRow("SELECT * FROM suserinchat WHERE ID = $1", userInChat).Scan(&userInChatt.ID)
	if err0 != nil {
		return &model.UserInChat{}, err0
	}

	err := pgdb.QueryRow("UPDATE suserinchat SET IsAdmin = true WHERE ID = $1", userInChat).Scan(&userInChatt.ID)
	if err != nil {
		return &model.UserInChat{}, err
	}

	return &model.UserInChat{
		ID:     userInChatt.ID,
		UserID: userInChatt.UserID,
		ChatID: userInChatt.ChatID,
	}, nil
}

func RemoveUserAdmin(userInChat int) (*model.UserInChat, error) {
	var userInChatt model.UserInChat
	err0 := pgdb.QueryRow("SELECT * FROM suserinchat WHERE ID = $1", userInChat).Scan(&userInChatt.ID)
	if err0 != nil {
		return &model.UserInChat{}, err0
	}

	err := pgdb.QueryRow("UPDATE suserinchat SET IsAdmin = false WHERE ID = $1", userInChat).Scan(&userInChatt.ID)
	if err != nil {
		return &model.UserInChat{}, err
	}

	return &model.UserInChat{
		ID:     userInChatt.ID,
		UserID: userInChatt.UserID,
		ChatID: userInChatt.ChatID,
	}, nil
}

func RemoveUserFromChat(userInChat int) (*model.UserInChat, error) {
	var userInChatt model.UserInChat
	err0 := pgdb.QueryRow("SELECT * FROM suserinchat WHERE ID = $1", userInChat).Scan(&userInChatt.ID)
	if err0 != nil {
		return &model.UserInChat{}, err0
	}

	err := pgdb.QueryRow("DELETE FROM suserinchat WHERE ID = $1", userInChat).Scan(&userInChatt.ID)
	if err != nil {
		return &model.UserInChat{}, err
	}

	return &model.UserInChat{
		ID:     userInChatt.ID,
		UserID: userInChatt.UserID,
		ChatID: userInChatt.ChatID,
	}, nil
}

func AddUserToChat(userInChat model.UserInChat) (*model.UserInChat, error) {
	var id int
	err := pgdb.QueryRow("INSERT INTO suserinchat (ChatID, UserID) VALUES ($1, $2)", userInChat.ChatID, userInChat.UserID).Scan(&id)
	if err != nil {
		return &model.UserInChat{}, err
	}
	return &model.UserInChat{ID: id, ChatID: userInChat.ChatID, UserID: userInChat.UserID}, nil
}

func CreateChat(chat model.Chat) (*model.Chat, error) {
	var id int
	err := pgdb.QueryRow("INSERT INTO schat (ChatName) VALUES ($1)", chat.ChatName).Scan(&id)
	if err != nil {
		return &model.Chat{}, err
	}
	return &model.Chat{ID: id, ChatName: chat.ChatName}, nil
}

func GetChatByID(id int) (*model.Chat, error) {
	chat := &model.Chat{}
	err := pgdb.QueryRow("SELECT * FROM schat WHERE ID = $1", id).Scan(&chat.ID, &chat.ChatName)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func UpdateChat(chat model.Chat) (*model.Chat, error) {
	_, err := pgdb.Exec("UPDATE schat SET ChatName = $1 WHERE ID = $2", chat.ChatName, chat.ID)
	if err != nil {
		return &model.Chat{}, err
	}
	return &model.Chat{ID: chat.ID, ChatName: chat.ChatName}, nil
}

func GetUsersByChatID(chatID int) ([]*model.User, error) {
	var users []*model.User
	rows, err := pgdb.Query("SELECT * FROM suser WHERE ID IN (SELECT UserID FROM suserinchat WHERE ChatID = $1)", chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.UserName, &user.UserEmail)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByMessageID(messageID int) (*model.User, error) {
	user := &model.User{}
	err := pgdb.QueryRow("SELECT * FROM suser WHERE ID IN (SELECT UserID FROM smessageofuser WHERE MessageID = $1)", messageID).Scan(&user.ID, &user.UserName, &user.UserEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetReactionsByMessageID(messageID int) ([]*model.ReactionToMessage, error) {
	var reactions []*model.ReactionToMessage
	rows, err := pgdb.Query("SELECT * FROM sreactiontomessage WHERE MessageID = $1", messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		reaction := &model.ReactionToMessage{}
		err := rows.Scan(&reaction.ID, &reaction.MessageID, &reaction.ReactedUserID, &reaction.Reaction)
		if err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}
	return reactions, nil
}

func GetMessagesByChat(chatID int) ([]*model.Message, error) {
	var messages []*model.Message
	rows, err := pgdb.Query("SELECT * FROM smessage WHERE ID IN (SELECT MessageID FROM smessageinchat WHERE ChatID = $1)", chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		message := &model.Message{}
		err := rows.Scan(&message.ID, &message.MessageText, &message.QuotedMsgID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
