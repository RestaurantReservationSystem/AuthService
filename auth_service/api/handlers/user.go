package handlers

import (
	"auth_service/api/token"
	pb "auth_service/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"strconv"
	"strings"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
func IsValidLimit(limit string) (int, error) {
	if limit == "" {
		limit += "0"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		return 0, err
	}
	return limit1, nil
}
func IsValidOffset(offset string) (int, error) {
	if offset == "" {
		offset += "0"
	}
	offset1, err := strconv.Atoi(offset)
	if err != nil {
		return 0, err
	}
	return offset1, nil
}

// CreateUser 		handles the creation of a new user
// @Summary 		Create User
// @Description 	Create page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Create"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/user/register [post]

func (h *Handler) CreateUser(gn *gin.Context) {
	userName := gn.Query("user_name")
	password := gn.Query("password")
	email := gn.Query("email")
	if len(userName) < 4 {
		BadRequest(gn, fmt.Errorf("error ->user_name is not vaildate"))
	}
	if len(password) < 7 {
		BadRequest(gn, fmt.Errorf("error ->password is not vaildate"))
	}
	if len(email) < 7 && !strings.Contains(email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("error ->email is no vaildate"))
	}

	request := pb.RegisterUserRequest{
		UserName: userName,
		Password: password,
		Email:    email,
	}
	_, err := h.UserService.CreateUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	Created(gn, fmt.Errorf("succes that created"))

}

// UpdateUser 		handles the creation of a new user
// @Summary			Update User
// @Description 	Update page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id 		path   string     true   "User ID"
// @Param   		Update  body   pb.UpdatedUserRequest     true   "Update"
// @Success 		200   {string} string    "Update Successful"
// @Failure 		401   {string} string    "Error while created"
// @Router 			/api/user/update{id} [put]

func (h *Handler) UpdateUser(gn *gin.Context) {
	userName := gn.Query("user_name")
	password := gn.Query("password")
	email := gn.Query("email")
	if len(userName) < 4 {
		BadRequest(gn, fmt.Errorf("error ->user_name is not vaildate"))
	}
	if len(password) < 7 {
		BadRequest(gn, fmt.Errorf("error ->password is not vaildate"))
	}
	if len(email) < 7 && !strings.Contains(email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("error ->email is no vaildate"))
	}
	id := gn.Param("id")

	if isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("is id is not validate"))
	}
	request := pb.UpdatedUserRequest{
		UserName: userName,
		Password: password,
		Email:    email,
		Id:       id,
	}
	_, err := h.UserService.UpdateUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, fmt.Errorf("succes that Updated"))

}

// DeleteUser 		handles the creation of a new User
// @Summary			Delete User
// @Description 	Delete page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id   path     string   true   "User ID"
// @Success 		200 {string}  string   "Delete Successful"
// @Failure 		401 {string}  string   "Error while Deleted"
// @Router 			/api/user/delete{id} [delete]
func (h *Handler) DeleteUser(gn *gin.Context) {
	id := gn.Param("id")

	if isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("is id is not validate"))
	}
	request := pb.IdRequest{
		Id: id,
	}
	_, err := h.UserService.DeleteUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, fmt.Errorf("succes that Deleted"))

}

// GetByIdUser 		handles the creation of a new User
// @Summary 		GetById User
// @Description 	GetById page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id   path      string   true    "User ID"
// @Success 		200 {object}   pb.User  "GetById Successful"
// @Failure 		401 {string}   string   "Error while GetByIdd"
// @Router 			/api/user/get-by-id/{id} [get]

func (h *Handler) GetUserById(gn *gin.Context) {
	id := gn.Param("id")
	if isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("is id is not validate"))
	}
	request := pb.IdRequest{
		Id: id,
	}
	response, err := h.UserService.GetByIdUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, response)

}

// GetAllUser godoc
// @Summary Get all users
// @Description Get all users with filtering options
// @Tags users
// @Accept json
// @Produce json
// @Param user_name query string false "User Name"
// @Param password query string false "Password"
// @Param email query string false "Email"
// @Param id path string true "User ID"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} pb.GetAllUserResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (h *Handler) GetAllUser(gn *gin.Context) {
	userName := gn.Query("user_name")
	password := gn.Query("password")
	email := gn.Query("email")
	if len(userName) < 4 {
		BadRequest(gn, fmt.Errorf("error ->user_name is not vaildate"))
	}
	if len(password) < 7 {
		BadRequest(gn, fmt.Errorf("error ->password is not vaildate"))
	}
	if len(email) < 7 && !strings.Contains(email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("error ->email is no vaildate"))
	}
	id := gn.Param("id")

	if isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("is id is not validate"))
	}
	limit := gn.Query("limit")
	limit1, err := IsValidLimit(limit)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	offset := gn.Query("offset")
	offset1, err := IsValidOffset(offset)
	if err != nil {
		BadRequest(gn, err)
	}

	request := pb.GetAllUserRequest{
		UserName: userName,
		Password: password,
		Email:    email,
		Filter:   &pb.Filter{Limit: int64((limit1)), Offset: int64(offset1)},
	}
	response, err := h.UserService.GetAllUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, response)

}

// GetByIdUser 		handles the creation of a new User
// @Summary 		/LoginUser
// @Description 	LoginUser page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body  pb.LoginRequest    true     "Create"
// @Success 		200 {object}   pb.LoginResponse "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/api/user/login [post]
func (h *Handler) LoginUser(gn *gin.Context) {
	userName := gn.Query("user_name")
	if len(userName) < 4 {
		BadRequest(gn, fmt.Errorf("error ->user_name is not vaildate"))
	}

	request := pb.LoginRequest{
		UserName: userName,
	}
	response, err := h.UserService.LoginUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	t := token.GENERATEJWTToken(response)
	gn.JSON(200, t)

}
