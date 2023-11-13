package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	IsSuccessfully bool        `json:"isSuccessfully"`
	Message        string      `json:"message"`
	UserData       apiUserData `json:"userData"`
}

func UserData(c *gin.Context) {
	userData, err := getUserData(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, response{false, "Invalid Data", apiUserData{}})
		return
	}

	c.JSON(http.StatusOK, response{true, "Success", userData})
}
