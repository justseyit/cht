package api

import (
	"chatservice/model"
	pb "chatservice/proto"
	authservice "chatservice/service/auth"
	"context"
)

//chatServer is used to implement ChatService in proto generated .go files.

type authserviceServer struct {
	pb.AuthServiceServer
}

func Login(ctx context.Context, req *pb.RequestUserLogin) (*pb.User, error) {
	user, err := authservice.UserLogin(req.UserEmail, req.UserPassword)
	if err != nil {
		return &pb.User{}, err
	}
	return &pb.User{
		Id:        int32(user.ID),
		Username:  user.UserName,
		UserEmail: user.UserEmail,
	}, nil
}

func Register(ctx context.Context, req *pb.RequestUserRegister) (*pb.User, error) {
	user, err := authservice.UserRegister(&model.User{
		UserName:  req.Username,
		UserEmail: req.UserEmail,
	}, req.UserPassword)
	if err != nil {
		return &pb.User{}, err
	}
	return &pb.User{
		Id:        int32(user.ID),
		Username:  user.UserName,
		UserEmail: user.UserEmail,
	}, nil
}

func Logout(ctx context.Context, req *pb.RequestUserLogout) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func NewAuthServiceServer() pb.AuthServiceServer {
	return &authserviceServer{}
}
