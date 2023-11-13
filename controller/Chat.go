package controller

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Chat(c *gin.Context) {
	userData, err := getUserData(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, response{false, "Invalid Data", apiUserData{}})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response{false, "Internal server error", apiUserData{}})
		return
	}

	room := database.ChatRooms["programmig"]

	chatUser := models.ChatUser{
		Name:        userData.Name,
		Email:       userData.Email,
		Conn:        conn,
		RoomManager: room,
		Send:        make(chan []byte, 256),
	}

	room.Register <- &chatUser

	go chatUser.ReadMessages()
	go chatUser.WriteMessages()
}
