package handlers

import "auth_service/genproto"

type Handler struct {
	UserService genproto.UserServiceClient
}

func NewHandler(userHandler genproto.UserServiceClient) *Handler {
	return &Handler{UserService: userHandler}
}
