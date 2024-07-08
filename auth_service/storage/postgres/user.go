package postgres

import (
	pb "auth_service/genproto"
	"auth_service/help"
	"database/sql"
	"time"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) Register(request *pb.RegisterUserRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into users(user_name,password,email)values ($1,$2,$3,$4)", request.UserName, request.Password, request.Email)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo *UserRepository) UpdateUser(request *pb.UpdatedUserRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update users set user_name=$1,password=$2,email=$3  ,updated_at=$4 where id=$5 and deleted_at is null", request.UserName, request.Password, request.Email, time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo *UserRepository) DeleteUser(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  users set  deleted_at=current_timestamp where  deleted_at is null and id=$1", request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo *UserRepository) GetUserById(request *pb.IdRequest) (*pb.UserResponse, error) {
	var userResponse pb.UserResponse
	err := repo.Db.QueryRow("select user_name,password,email,created_at,updated_at,deleted_at from users where id=$1 and deleted_at is null", request.Id).Scan(&userResponse.UserName, &userResponse.Password, &userResponse.Email, &userResponse.CreatedAt, &userResponse.UpdatedAt, &userResponse.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &userResponse, err

}
func (repo *UserRepository) GeAllUser(request *pb.GetAllUserRequest) (*pb.GetAllUsers, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(request.UserName) > 0 {
		params["user_name"] = request.UserName
		filter += " and user_name = :user_name "
	}
	if len(request.Password) > 0 {
		params["password"] = request.Password
		filter += " and password:=password"
	}
	if len(request.Email) > 0 {
		params["email"] = request.Email
		filter += "and email:=email"
	}
	if request.Filter.Limit > 0 {
		params["limit"] = request.Filter.Limit
		filter += "and limit:=limit"
	}
	if request.Filter.Offset > 0 {
		params["offset"] = request.Filter.Offset
		filter += "and offset:=offset"
	}

	query := "select user_name,password ,email,created_at,updated_at,deleted_at from users  where  deleted_at is null"

	query = query + filter + limit + offset
	query, arr = help.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var users []*pb.UserResponse
	for rows.Next() {
		var userResponse pb.UserResponse
		err := rows.Scan(&userResponse.UserName, &userResponse.Password, &userResponse.Email, &userResponse.CreatedAt, &userResponse.UpdatedAt, &userResponse.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &userResponse)
	}
	return &pb.GetAllUsers{Users: users}, nil
}
func (repo *UserRepository) Login(request *pb.LoginRequest) (*pb.LoginResponse, error) {
	var loginUser pb.LoginResponse
	err := repo.Db.QueryRow("select user_name,password,email from  users where user_name=$1", request.UserName).Scan(&loginUser.UserName, &loginUser.Password, &loginUser.Email)
	if err != nil {
		return nil, err
	}
	return &loginUser, nil
}
