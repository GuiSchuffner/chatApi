package models

type ChatRoomManager struct {
	Name       string
	Users      map[*ChatUser]bool
	Register   chan *ChatUser
	Unregister chan *ChatUser
	Broadcast  chan []byte
}

func NewChatRoomManager(name string) *ChatRoomManager {
	return &ChatRoomManager{
		Name:       name,
		Users:      make(map[*ChatUser]bool),
		Register:   make(chan *ChatUser, 5),
		Unregister: make(chan *ChatUser, 5),
		Broadcast:  make(chan []byte, 5),
	}
}

func (manager *ChatRoomManager) Run() {
	for {
		select {
		case user := <-manager.Register:
			manager.Users[user] = true
		case user := <-manager.Unregister:
			if _, ok := manager.Users[user]; ok {
				delete(manager.Users, user)
				close(user.Send)
				user.Conn.Close()
			}
		case message := <-manager.Broadcast:
			for user := range manager.Users {
				select {
				case user.Send <- message:
				default:
					close(user.Send)
					delete(manager.Users, user)
					user.Conn.Close()
				}
			}
		}
	}
}
