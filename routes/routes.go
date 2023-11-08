package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	AuthRoutes(r.Group("/auth"))
	UserRoutes(r.Group("/user"))
	r.Run()
}
