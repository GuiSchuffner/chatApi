package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ChatUser struct {
	Name        string
	Email       string
	Conn        *websocket.Conn
	RoomManager *ChatRoomManager
	Send        chan []byte
}

func NewChatUser(name string, email string, conn *websocket.Conn, hub *ChatRoomManager) *ChatUser {
	return &ChatUser{
		Name:        name,
		Email:       email,
		Conn:        conn,
		RoomManager: hub,
		Send:        make(chan []byte),
	}
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func (user *ChatUser) ReadMessages() {
	defer func() {
		fmt.Println("Closing connection in readMessages")
		user.RoomManager.Unregister <- user
	}()
	user.Conn.SetReadLimit(maxMessageSize)
	user.Conn.SetReadDeadline(time.Now().Add(pongWait))
	user.Conn.SetPongHandler(user.pongHandler)
	for {
		messageType, message, err := user.Conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		chatMessage := ChatMessage{
			Message:   string(bytes.TrimSpace(message)),
			UserName:  user.Name,
			UserEmail: user.Email,
			Time:      time.Now().UTC().Format(time.RFC3339),
		}
		messageToSend, err := json.Marshal(chatMessage)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		user.RoomManager.Broadcast <- messageToSend
	}
}

func (user *ChatUser) WriteMessages() {
	defer func() {
		fmt.Println("Closing connection in writeMessages")
		user.RoomManager.Unregister <- user
	}()
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case message, ok := <-user.Send:
			if !ok {
				user.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			user.Conn.WriteMessage(websocket.TextMessage, message)
		case <-ticker.C:
			user.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := user.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (user *ChatUser) pongHandler(pongMsg string) error {
	return user.Conn.SetReadDeadline(time.Now().Add(pongWait))
}
