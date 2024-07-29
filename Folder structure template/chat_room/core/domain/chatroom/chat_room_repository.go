package chatroom

type ChatroomRepository interface {
	CreateChatroom(chatroom *Entity_Chatroom) (*Entity_Chatroom, error)
	CheckChatroomId(id string) (*Entity_Chatroom, error)
	GetChatrooms() ([]*Entity_Chatroom, error)
	UpdateChatroom(chatroom *Entity_Chatroom) error
	SendMessage(message *Entity_Message) (*Entity_Message, error)
	GetMessages(chatroomID string) ([]*Entity_Message, error)
	RemoveUserFromChatroom(chatroomID string, userID string) error
}
