package api

import (
	"chatservice/model"
	pb "chatservice/proto"
	chatservice "chatservice/service/chat"
	"context"
	"log"
)

//chatServer is used to implement ChatService in proto generated .go files.

type chatServiceServer struct {
	Chats chan pb.StreamByUser
	pb.ChatServiceServer
}

// SubscribeByUser is a method of the ChatServiceServer interface.
// It is used to subscribe a user to the chat service.
func (s *chatServiceServer) SubscribeByUser(user *pb.User, stream pb.ChatService_SubscribeByUserServer) error {
	//Find the chat groups that the user is in, from the database.
	chats, err := chatservice.GetChatsByUserID(int(user.Id))
	if err != nil {
		log.Fatalf("Failed to get chats by user id: %v", err)
		return err
	}

	var streamByUser pb.StreamByUser

	//Create a stream for each chat group.
	for _, chat := range chats {
		var filledChat pb.FilledChat
		messages, err := chatservice.GetMessagesByChatID(int(chat.ID))
		if err != nil {
			log.Fatalf("Failed to get messages by chat id: %v", err)
			return err
		}
		//Create a stream for each chat group.
		for _, message := range messages {
			filledChat.Messages = append(filledChat.Messages, &pb.Message{
				Id:              int32(message.ID),
				MessageText:     message.MessageText,
				QuotedMessageID: int32(message.QuotedMsgID),
			})
		}

		streamByUser.MsgStream.FilledChats = append(streamByUser.MsgStream.FilledChats, &filledChat)

		//Send the stream to the user.
		if err := stream.Send(&streamByUser); err != nil {
			log.Fatalf("Failed to send stream to user: %v", err)
			return err
		}

		//Add the stream to the list of streams.
		s.Chats <- streamByUser
	}

	//Listen for new messages. If a new message is received, send it to the user through the stream and add it to the database.
	/*for {
		//Receive a new message from the stream.
		err := stream.RecvMsg(pb.Empty{})
		if err != nil {
			log.Fatalf("Failed to receive message from stream: %v", err)
			return err
		}

		//Send the message to the user.
		if err := stream.Send(&streamByUser); err != nil {
			log.Fatalf("Failed to send message to user: %v", err)
			return err
		}

		//Add the message to the database.
		_, err = chatservice.SendMessageToChat(int(streamByUser.MsgStream.FilledChats[0].Messages[0].QuotedMessageID), streamByUser.MsgStream.FilledChats[0].Messages[0].MessageText)
		if err != nil {
			log.Fatalf("Failed to send message to database: %v", err)
			return err
		}
	}*/

	return nil

}

// SendMessageToChat is a method of the ChatServiceServer interface.
// It is used to send a message to a chat.
func (s *chatServiceServer) SendMessageToChat(ctx context.Context, req *pb.RequestSendMessage) (*pb.ResponseSendMessage, error) {
	msg, err := chatservice.SendMessageToChat(int(req.ChatID), req.Message.MessageText)
	if err != nil {
		return &pb.ResponseSendMessage{
			MessageSent: &pb.Message{},
			Message:     err.Error(),
			Success:     false,
		}, err
	}
	return &pb.ResponseSendMessage{
		MessageSent: &pb.Message{
			Id:              int32(msg.ID),
			MessageText:     msg.MessageText,
			QuotedMessageID: int32(msg.QuotedMsgID),
		},
		Message: "",
		Success: true,
	}, nil
}

// CreateChat is a method of the ChatServiceServer interface.
// It is used to create a chat.
func (s *chatServiceServer) CreateChat(ctx context.Context, req *pb.RequestCreateChat) (*pb.ResponseCreateChat, error) {
	chat, err := chatservice.CreateChat(&model.Chat{
		ChatName:        req.Chat.ChatName,
		ChatLanguage:    req.Chat.ChatLanguage,
		ChatDescription: req.Chat.ChatDescription,
		ChatImage:       req.Chat.ChatImage,
		ChatCreatedAt:   req.Chat.ChatCreatedAt,
	})
	if err != nil {
		return &pb.ResponseCreateChat{
			Chat:    &pb.Chat{},
			Message: err.Error(),
			Success: false,
		}, err
	}
	return &pb.ResponseCreateChat{
		Chat: &pb.Chat{
			Id:              int32(chat.ID),
			ChatName:        chat.ChatName,
			ChatLanguage:    chat.ChatLanguage,
			ChatDescription: chat.ChatDescription,
			ChatImage:       chat.ChatImage,
			ChatCreatedAt:   chat.ChatCreatedAt,
		},
		Message: "",
		Success: true,
	}, nil
}

// UpdateChat is a method of the ChatServiceServer interface.
// It is used to update a chat.
func (s *chatServiceServer) UpdateChat(ctx context.Context, req *pb.RequestUpdateChat) (*pb.ResponseUpdateChat, error) {
	chat, err := chatservice.UpdateChat(&model.Chat{
		ID:              int(req.Chat.Id),
		ChatName:        req.Chat.ChatName,
		ChatLanguage:    req.Chat.ChatLanguage,
		ChatDescription: req.Chat.ChatDescription,
		ChatImage:       req.Chat.ChatImage,
		ChatCreatedAt:   req.Chat.ChatCreatedAt,
	})
	if err != nil {
		return &pb.ResponseUpdateChat{
			Chat:    &pb.Chat{},
			Message: err.Error(),
			Success: false,
		}, err
	}
	return &pb.ResponseUpdateChat{
		Chat: &pb.Chat{
			Id:              int32(chat.ID),
			ChatName:        chat.ChatName,
			ChatLanguage:    chat.ChatLanguage,
			ChatDescription: chat.ChatDescription,
			ChatImage:       chat.ChatImage,
			ChatCreatedAt:   chat.ChatCreatedAt,
		},
		Message: "",
		Success: true,
	}, nil
}

// GetChatsByUser is a method of the ChatServiceServer interface.
// It is used to get all chats of a user.
func (s *chatServiceServer) GetChatsByUser(ctx context.Context, req *pb.RequestGetChatsByUser) (*pb.ResponseGetChatsByUser, error) {
	chats, err := chatservice.GetChatsByUserID(int(req.UserID))
	if err != nil {
		return &pb.ResponseGetChatsByUser{
			Chats:   []*pb.Chat{},
			Message: err.Error(),
			Success: false,
		}, err
	}
	var chatss []*pb.Chat
	for _, chat := range chats {
		chatss = append(chatss, &pb.Chat{
			Id:              int32(chat.ID),
			ChatName:        chat.ChatName,
			ChatLanguage:    chat.ChatLanguage,
			ChatDescription: chat.ChatDescription,
			ChatImage:       chat.ChatImage,
			ChatCreatedAt:   chat.ChatCreatedAt,
		})
	}
	return &pb.ResponseGetChatsByUser{
		Chats:   chatss,
		Message: "",
		Success: true,
	}, nil
}

// GetUsersByChat is a method of the ChatServiceServer interface.
// It is used to get all users of a chat.
func (s *chatServiceServer) GetUsersByChat(ctx context.Context, req *pb.RequestGetUsersByChat) (*pb.ResponseGetUsersByChat, error) {
	users, err := chatservice.GetUsersInChat(int(req.ChatID))
	if err != nil {
		return &pb.ResponseGetUsersByChat{
			Users:   []*pb.User{},
			Message: err.Error(),
			Success: false,
		}, err
	}
	var usrs []*pb.User
	for _, user := range users {
		usrs = append(usrs, &pb.User{
			Id:        int32(user.ID),
			Username:  user.UserName,
			UserEmail: user.UserEmail,
		})
	}
	return &pb.ResponseGetUsersByChat{
		Users:   usrs,
		Message: "",
		Success: true,
	}, nil
}

// GetMessagesByChat is a method of the ChatServiceServer interface.
// It is used to get all messages of a chat.
func (s *chatServiceServer) GetMessagesByChat(ctx context.Context, req *pb.RequestGetMessagesByChat) (*pb.ResponseGetMessagesByChat, error) {
	messages, err := chatservice.GetMessagesByChatID(int(req.ChatID))
	if err != nil {
		return &pb.ResponseGetMessagesByChat{
			Messages: []*pb.Message{},
			Message:  err.Error(),
			Success:  false,
		}, err
	}
	var msgs []*pb.Message
	for _, message := range messages {
		msgs = append(msgs, &pb.Message{
			Id:              int32(message.ID),
			MessageText:     message.MessageText,
			QuotedMessageID: int32(message.QuotedMsgID),
		})
	}
	return &pb.ResponseGetMessagesByChat{
		Messages: msgs,
		Message:  "",
		Success:  true,
	}, nil
}

// AddUserToChat is a method of the ChatServiceServer interface.
// It is used to add a user to a chat.
func (s *chatServiceServer) AddUserToChat(ctx context.Context, req *pb.RequestAddUserToChat) (*pb.ResponseAddUserToChat, error) {
	user, err := chatservice.AddUserToChat(&model.UserInChat{
		UserID: int(req.UserID),
		ChatID: int(req.ChatID),
	})
	if err != nil {
		return &pb.ResponseAddUserToChat{
			UserInChat: &pb.UserInChat{},
			Message:    err.Error(),
			Success:    false,
		}, err
	}
	return &pb.ResponseAddUserToChat{
		UserInChat: &pb.UserInChat{
			UserID: int32(user.UserID),
			ChatID: int32(user.ChatID),
		},
		Message: "",
		Success: true,
	}, nil
}

// RemoveUserFromChat is a method of the ChatServiceServer interface.
// It is used to remove a user from a chat.
func (s *chatServiceServer) RemoveUserFromChat(ctx context.Context, req *pb.RequestRemoveUserFromChat) (*pb.ResponseRemoveUserFromChat, error) {
	id := int(req.UserInChatID)
	userInChat, err := chatservice.RemoveUserFromChat(&id)

	if err != nil {
		return &pb.ResponseRemoveUserFromChat{
			UserInChat: &pb.UserInChat{},
			Message:    err.Error(),
			Success:    false,
		}, err
	}
	return &pb.ResponseRemoveUserFromChat{
		UserInChat: &pb.UserInChat{
			UserID: int32(userInChat.UserID),
			ChatID: int32(userInChat.ChatID),
		},
		Message: "",
		Success: true,
	}, nil
}

// MakeUserAdmin is a method of the ChatServiceServer interface.
// It is used to make a user an admin of a chat.
func (s *chatServiceServer) MakeUserAdmin(ctx context.Context, req *pb.RequestMakeUserAdmin) (*pb.ResponseMakeUserAdmin, error) {
	id := int(req.UserInChatID)
	user, err := chatservice.MakeUserAdmin(&id)
	if err != nil {
		return &pb.ResponseMakeUserAdmin{
			UserInChat: &pb.UserInChat{},
			Message:    err.Error(),
			Success:    false,
		}, err
	}
	return &pb.ResponseMakeUserAdmin{
		UserInChat: &pb.UserInChat{
			UserID: int32(user.UserID),
			ChatID: int32(user.ChatID),
		},
		Message: "",
		Success: true,
	}, nil
}

// RemoveUserAdmin is a method of the ChatServiceServer interface.
// It is used to remove a user from the admin list of a chat.
func (s *chatServiceServer) RemoveUserAdmin(ctx context.Context, req *pb.RequestRemoveUserAdmin) (*pb.ResponseRemoveUserAdmin, error) {
	id := int(req.UserInChatID)
	user, err := chatservice.RemoveUserAdmin(&id)
	if err != nil {
		return &pb.ResponseRemoveUserAdmin{
			UserInChat: &pb.UserInChat{},
			Message:    err.Error(),
			Success:    false,
		}, err
	}
	return &pb.ResponseRemoveUserAdmin{
		UserInChat: &pb.UserInChat{
			UserID: int32(user.UserID),
			ChatID: int32(user.ChatID),
		},
		Message: "",
		Success: true,
	}, nil
}

// RequestReactAMessage is a method of the ChatServiceServer interface.
// It is used to react a message.
func (s *chatServiceServer) RequestReactAMessage(ctx context.Context, req *pb.RequestReactMessage) (*pb.ResponseReactMessage, error) {
	reactionToMessage, err := chatservice.ReactToMessage(&model.ReactionToMessage{
		ReactedUserID: int(req.ReactionToMessage.ReactedUserID),
		MessageID:     int(req.MessageID),
		Reaction:      req.ReactionToMessage.Reaction,
	})
	if err != nil {
		return &pb.ResponseReactMessage{
			ReactionToMessage: &pb.ReactionToMessage{},
			Message:           err.Error(),
			Success:           false,
		}, err
	}
	return &pb.ResponseReactMessage{
		ReactionToMessage: &pb.ReactionToMessage{
			Id:            int32(reactionToMessage.ID),
			ReactedUserID: int32(reactionToMessage.ReactedUserID),
			MessageID:     int32(reactionToMessage.MessageID),
			Reaction:      reactionToMessage.Reaction,
		},
		Message: "",
		Success: true,
	}, nil
}

// RequestUnreactAMessage is a method of the ChatServiceServer interface.
// It is used to unreact a message.
func (s *chatServiceServer) RequestUnreactAMessage(ctx context.Context, req *pb.RequestUnreactMessage) (*pb.ResponseUnreactMessage, error) {
	reactionToMessage, err := chatservice.UnreactMessage(&model.ReactionToMessage{
		ReactedUserID: int(req.ReactionToMessage.ReactedUserID),
		MessageID:     int(req.MessageID),
		Reaction:      req.ReactionToMessage.Reaction,
	})
	if err != nil {
		return &pb.ResponseUnreactMessage{
			ReactionToMessage: &pb.ReactionToMessage{},
			Message:           err.Error(),
			Success:           false,
		}, err
	}
	return &pb.ResponseUnreactMessage{
		ReactionToMessage: &pb.ReactionToMessage{
			Id:            int32(reactionToMessage.ID),
			ReactedUserID: int32(reactionToMessage.ReactedUserID),
			MessageID:     int32(reactionToMessage.MessageID),
			Reaction:      reactionToMessage.Reaction,
		},
		Message: "",
		Success: true,
	}, nil
}

// Disconnect is a method of the ChatServiceServer interface.
// It is used to disconnect a user from the chat service.
func (s *chatServiceServer) Disconnect(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

// NewChatServiceServer creates a new chatServiceServer.
func NewChatServiceServer() pb.ChatServiceServer {
	return &chatServiceServer{}
}
