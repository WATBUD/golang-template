package chatroom

import (
	"time"
)

type Entity_Chatroom struct {
	ID             string                    `json:"_id" bson:"_id"`
	BaseID         string                    `json:"base_id" bson:"base_id"`
	BoardID        string                    `json:"board_id" bson:"board_id"`
	Title          string                    `json:"title" bson:"title"`
	LastMsg        string                    `json:"last_msg" bson:"last_msg"`
	LastMsgTime    time.Time                 `json:"last_msg_time" bson:"last_msg_time"`
	LastMsgSender  string                    `json:"last_msg_sender" bson:"last_msg_sender"`
	Avatar         string                    `json:"avatar" bson:"avatar"`
	CreatorID      string                    `json:"creator_id" bson:"creator_id"`
	CreatedAt      time.Time                 `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at" bson:"updated_at"`
	PersonalConfig []*ChatroomPersonalConfig `json:"personal_config" bson:"personal_config"`
	Type           string                    `json:"type" bson:"type"` // "private" or "board"
}

type ChatroomPersonalConfig struct {
	LastReadTime   time.Time `json:"last_read_time" bson:"last_read_time"`
	UnreadCount    int       `json:"unread_count" bson:"unread_count"`
	NotificationOn bool      `json:"notification_on" bson:"notification_on"`
	UserID         string    `json:"userId" bson:"user_id"`
	IsMuted        bool      `json:"isMuted" bson:"is_muted"`
	IsHidden       bool      `json:"isHidden" bson:"is_hidden"`
	IsPinned       bool      `json:"isPinned" bson:"is_pinned"`
}

type Entity_Message struct {
	ID         string    `json:"_id" bson:"_id"`
	ChatroomID string    `json:"chatroom_id" bson:"chatroom_id"`
	SenderID   string    `json:"sender_id" bson:"sender_id"`
	Type       string    `json:"type" bson:"type"`
	Data       string    `json:"data" bson:"data"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}
