package application

import (
	"chat_room_mod/core/domain/chatroom"
	"errors"
	"time"
)

type ChatroomUsecase interface {
	CreateChatroom(title string, base_id string) (*chatroom.Entity_Chatroom, error)
	CheckChatroomId(id string) (*chatroom.Entity_Chatroom, error)
	GetChatrooms() ([]*chatroom.Entity_Chatroom, error)
	UpdateChatroom(chatroom *chatroom.Entity_Chatroom) error
	SendMessage(chatroomID, senderID, messageType, data string, ctime, utime time.Time) (*chatroom.Entity_Message, error)
	GetMessages(chatroomID string) ([]*chatroom.Entity_Message, error)
	RemoveUserFromChatroom(chatroomID string, userID string) error
}

type chatroomUsecase struct {
	service *chatroomService
}

func NewChatroomUsecase(service *chatroomService) ChatroomUsecase {
	return &chatroomUsecase{service: service}
}

func (u *chatroomUsecase) CreateChatroom(title string, base_id string) (*chatroom.Entity_Chatroom, error) {
	chatroom := &chatroom.Entity_Chatroom{
		Title:       title,
		UnreadCount: 0,
		//MemberNumber: 1,
		BaseID: base_id,
	}
	createdChatroom, err := u.service.CreateChatroom(chatroom)
	if err != nil {
		return nil, err
	}
	return createdChatroom, nil
}

func (u *chatroomUsecase) CheckChatroomId(id string) (*chatroom.Entity_Chatroom, error) {
	return u.service.CheckChatroomId(id)
}

func (u *chatroomUsecase) GetChatrooms() ([]*chatroom.Entity_Chatroom, error) {
	return u.service.GetChatrooms()
}

func (u *chatroomUsecase) UpdateChatroom(chatroom *chatroom.Entity_Chatroom) error {
	return u.service.UpdateChatroom(chatroom)
}

func (u *chatroomUsecase) SendMessage(chatroomID, senderID, messageType, data string, ctime, utime time.Time) (*chatroom.Entity_Message, error) {
	chatroom, err := u.service.CheckChatroomId(chatroomID)
	if err != nil {
		return nil, errors.New("chatroom does not exist")
	}

	if chatroom == nil {
		return nil, errors.New("chatroom does not exist")
	}

	return u.service.SendMessage(chatroomID, senderID, messageType, data, ctime, utime)
}

func (u *chatroomUsecase) GetMessages(chatroomID string) ([]*chatroom.Entity_Message, error) {
	return u.service.GetMessages(chatroomID)
}

func (u *chatroomUsecase) RemoveUserFromChatroom(chatroomID string, userID string) error {
	return u.service.RemoveUserFromChatroom(chatroomID, userID)
}
