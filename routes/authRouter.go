package routes

import (
	"github.com/GuiSchuffner/chatApi/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authRoute := r.Group("/auth")
	authRoute.POST("/register", controller.Register)
	authRoute.POST("/login", controller.Login)
}
