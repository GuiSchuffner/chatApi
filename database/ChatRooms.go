package database

import "github.com/GuiSchuffner/chatApi/models"

var ChatRooms map[string]*models.ChatRoomManager
var roomTopics = [8]string{"programing", "games", "music", "movies", "sports", "politics", "news", "random"}

func InitChatRooms() {
	for _, topic := range roomTopics {
		chatRoom := models.NewChatRoomManager(topic)
		ChatRooms[topic] = chatRoom
		go chatRoom.Run()
	}
}
