package postgres

import (
	pb "auth_service/genproto"
	"auth_service/storage"
	"reflect"
	"testing"
)

func UserRepo(t *testing.T) *UserRepository {
	db, err := storage.ConnectionDb()
	if err != nil {
		t.Error("ERROR : ", err)
		return nil
	}
	userRepo := NewUserRepository(db)
	return userRepo
}

func TestRegister(t *testing.T) {
	userRepo := UserRepo(t)

	userId := pb.RegisterUserRequest{UserName: "Bekzod", Password: "1111", Email: "example@gmail.com"}

	user, err := userRepo.Register(&userId)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := &pb.Void{}
	if !reflect.DeepEqual(user, case1) {
		t.Error("Result : ", user, "Expected : ", case1)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	userRepo := UserRepo(t)

	update := pb.UpdatedUserRequest{Id: "ae4de3a7-01e1-4893-b3a9-95d03d5a7b5c", UserName: "Sanjarbek", Password: "1111", Email: "sanjarbek@gmail.com"}

	user, err := userRepo.UpdateUser(&update)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := &pb.Void{}

	if !reflect.DeepEqual(user, case1) {
		t.Error("Result : ", user, "Expected : ", case1)
		return
	}
}

func TestDeleteUser(t *testing.T) {
	userRepo := UserRepo(t)

	id := pb.IdRequest{Id: "10b3e850-be43-4938-bf4a-562419d4b567"}

	user, err := userRepo.DeleteUser(&id)
	if err != nil {
		t.Error("Error : ", err)
	}

	case1 := &pb.Void{}

	if !reflect.DeepEqual(user, case1) {
		t.Error("Result : ", user, "Expected : ", case1)
		return
	}
}

func TestGetUserById(t *testing.T) {
	userRepo := UserRepo(t)
	if userRepo == nil {
		return
	}

	id := pb.IdRequest{Id: "ae4de3a7-01e1-4893-b3a9-95d03d5a7b5c"}

	user, err := userRepo.GetUserById(&id)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	expectedUser := &pb.UserResponse{
		UserName: "Sanjarbek",
		Password: "1111",
		Email:    "sanjarbek@gmail.com",
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Result : %+v, Expected : %+v", user, expectedUser)
	}
}

func TestGetAllUser(t *testing.T) {
	userRepo := UserRepo(t)
	if userRepo == nil {
		return
	}

	request := pb.GetAllUserRequest{Filter: &pb.Filter{Limit: int64(1)}}

	users, err := userRepo.GeAllUser(&request)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	expectedUser := &pb.UserResponse{
		UserName: "Sanjarbek",
		Password: "1111",
		Email:    "sanjarbek@gmail.com",
	}

	expectedUsers := &pb.GetAllUsers{
		Users: []*pb.UserResponse{expectedUser},
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Result : %+v, Expected : %+v", users, expectedUsers)
	}
}

func TestLogin(t *testing.T) {
	userRepo := UserRepo(t)

	request := pb.LoginRequest{UserName: "Sanjarbek"}

	resp, err := userRepo.Login(&request)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := &pb.LoginResponse{
		UserName: "Sanjarbek",
		Password: "1111",
		Email:    "sanjarbek@gmail.com",
	}

	if !reflect.DeepEqual(resp, case1) {
		t.Error("Result : ", resp, "Expected : ", case1)
		return
	}
}
