package api

import (
	"auth_service/api/handlers"
	"auth_service/api/middleware"
	pb "auth_service/genproto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	//"github.com/gin-gonic/gin"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Voting service
// @version 1.0
// @description Voting service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewGin(userC *grpc.ClientConn) *gin.Engine {
	user := pb.NewUserServiceClient(userC)
	h := handlers.NewHandler(user)
	r := gin.Default()
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.MiddleWare())
	u := r.Group("/api/user")
	u.POST("/register", h.CreateUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/get-all", h.GetAllUser)
	u.GET("/get-by-id/:id", h.GetUserById)
	u.POST("/login", h.LoginUser)

	return r
}
