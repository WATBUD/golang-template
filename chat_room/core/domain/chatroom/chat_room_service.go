package chatroom

import (
	"time"
)

type ChatroomService interface {
	CreateChatroom(chatroom *Entity_Chatroom) (*Entity_Chatroom, error)
	CheckChatroomId(id string) (*Entity_Chatroom, error)
	GetChatrooms() ([]*Entity_Chatroom, error)
	UpdateChatroom(chatroom *Entity_Chatroom) error
	SendMessage(chatroomID, senderID, messageType, data string, ctime, utime time.Time) (*Entity_Message, error)
	GetMessages(chatroomID string) ([]*Entity_Message, error)
	RemoveUserFromChatroom(chatroomID string, userID string) error
}
