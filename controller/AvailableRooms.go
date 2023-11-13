package controller

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/models"
	"github.com/gin-gonic/gin"
)

type availableRoomsResponse struct {
	AvailableRooms []*chatRoomRespose
}

type chatRoomRespose struct {
	RoomName       string `json:"roomName"`
	ConnectedUsers int    `json:"connectedUsers"`
}

func AvailableRooms(c *gin.Context) {
	rooms := database.ChatRooms

	var roomsResponse []*chatRoomRespose

	for _, room := range rooms {
		roomsResponse = append(roomsResponse, chatRoomManagerToResponse(room))
	}

	c.JSON(http.StatusOK, availableRoomsResponse{roomsResponse})
}

func chatRoomManagerToResponse(roomManager *models.ChatRoomManager) *chatRoomRespose {
	return &chatRoomRespose{
		RoomName:       roomManager.Name,
		ConnectedUsers: len(roomManager.Users),
	}
}
