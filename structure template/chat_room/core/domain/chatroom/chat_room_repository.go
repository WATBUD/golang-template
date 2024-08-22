package chatroom

type ChatroomRepository interface {
	CreateChatroom(chatroom *Entity_Chatroom) (*Entity_Chatroom, error)
	GetChatroomByID(id string) (*Entity_Chatroom, error)
	GetChatrooms(userID string) ([]*Entity_Chatroom, error)
	UpdateChatroom(chatroom *Entity_Chatroom) error
	SendMessage(message *Entity_Message) (*Entity_Message, error)
	GetMessages(chatroomID string) ([]*Entity_Message, error)
	RemoveUserFromChatroom(dto DTO_AddOrRemoveChatRoomUserRequest) error
	AddUserToChatroom(chatroomID string, personal_config []*ChatroomPersonalConfig) error
	SetUserChatRoomConfig(dto DTO_SetUserChatRoomConfig) error
	GetChatroomPersonalConfig(chatroomID string) ([]*ChatroomPersonalConfig, error)
}
