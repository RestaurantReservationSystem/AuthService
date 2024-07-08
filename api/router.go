package api

import (
	"auth_service/api/handlers"
	"auth_service/api/middleware"
	"github.com/gin-gonic/gin"
)

// @title Voting service
// @version 1.0
// @description Voting service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewGin(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.MiddleWare())
	u := r.Group("/api/user")
	u.POST("/register", h.CreateUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/get-all", h.GetAllUser)
	u.GET("/get-by-id/:id", h.GetUserById)
	u.POST("/login", h.LoginUser)
	//url := ginSwagger.URL("swagger/doc.json")
	////files.Handler, url
	//r.GET("/swagger/*any", ginSwagger.WrapHandler())
	return r
}
