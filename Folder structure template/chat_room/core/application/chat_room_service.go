package application

import (
	"chat_room_mod/core/domain/chatroom"
	"time"
)

type chatroomService struct {
	repo chatroom.ChatroomRepository
}

func NewChatroomService(repo chatroom.ChatroomRepository) *chatroomService {
	return &chatroomService{repo: repo}
}

func (s *chatroomService) CreateChatroom(chatroom *chatroom.Entity_Chatroom) (*chatroom.Entity_Chatroom, error) {
	return s.repo.CreateChatroom(chatroom)
}

func (s *chatroomService) CheckChatroomId(id string) (*chatroom.Entity_Chatroom, error) {
	return s.repo.CheckChatroomId(id)
}

func (s *chatroomService) GetChatrooms() ([]*chatroom.Entity_Chatroom, error) {
	return s.repo.GetChatrooms()
}

func (s *chatroomService) UpdateChatroom(chatroom *chatroom.Entity_Chatroom) error {
	return s.repo.UpdateChatroom(chatroom)
}

func (s *chatroomService) SendMessage(chatroomID, senderID, messageType, data string, ctime, utime time.Time) (*chatroom.Entity_Message, error) {
	message := &chatroom.Entity_Message{
		ChatroomID: chatroomID,
		SenderID:   senderID,
		Type:       messageType,
		Data:       data,
		CTime:      ctime,
		UTime:      utime,
	}

	// Save the message
	savedMessage, err := s.repo.SendMessage(message)
	if err != nil {
		return nil, err
	}

	// Update the chatroom with the latest message details
	chatroom, err := s.repo.CheckChatroomId(chatroomID)
	if err != nil {
		return nil, err
	}

	chatroom.LastMsg = data
	chatroom.LastMsgTime = ctime
	chatroom.LastMsgSender = senderID

	if err := s.repo.UpdateChatroom(chatroom); err != nil {
		return nil, err
	}

	return savedMessage, nil
}

func (s *chatroomService) GetMessages(chatroomID string) ([]*chatroom.Entity_Message, error) {
	return s.repo.GetMessages(chatroomID)
}

func (s *chatroomService) RemoveUserFromChatroom(chatroomID string, userID string) error {
	return s.repo.RemoveUserFromChatroom(chatroomID, userID)
}
