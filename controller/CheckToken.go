package controller

import (
	"fmt"
	"net/http"

	"github.com/GuiSchuffner/chatApi/utils/token"
	"github.com/gin-gonic/gin"
)

type input struct {
	Token string `json:"token"`
}

func CheckUserLogin(c *gin.Context) {
	var input input
	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err = token.VerifyToken(input.Token)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Valid token"})
}
