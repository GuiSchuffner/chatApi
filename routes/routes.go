package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	AuthRoutes(r)
	UserRoutes(r)
	r.Run()
}
