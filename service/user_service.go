package service

import (
	pb "auth_service/genproto"
	"auth_service/storage/postgres"
	"context"
)

type UserService struct {
	UserRepo *postgres.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (service *UserService) CreateUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.Void, error) {
	return service.UserRepo.Register(in)
}
func (service *UserService) DeleteUser(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.UserRepo.DeleteUser(in)
}
func (service *UserService) UpdateUser(ctx context.Context, in *pb.UpdatedUserRequest) (*pb.Void, error) {
	return service.UserRepo.UpdateUser(in)
}
func (service *UserService) GetByIdUser(ctx context.Context, in *pb.IdRequest) (*pb.UserResponse, error) {
	return service.UserRepo.GetUserById(in)
}
func (service *UserService) GetAllUser(ctx context.Context, in *pb.GetAllUserRequest) (*pb.GetAllUsers, error) {
	return service.UserRepo.GeAllUser(in)
}
func (service *UserService) LoginUser(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return service.UserRepo.Login(in)
}
func (service *UserService) GenerateToken(ctx context.Context, in *pb.LoginResponse) (*pb.AccessToken, error) {
	jwtToken, err := service.UserRepo.GENERATEJWTToken(in)
	if err != nil {
		return nil, err
	}

	return &pb.AccessToken{Token: jwtToken.AccessToken}, nil
}
