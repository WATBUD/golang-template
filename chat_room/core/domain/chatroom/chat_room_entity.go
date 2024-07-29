package chatroom

import (
	"time"
)

type Entity_Chatroom struct {
	ID          string `bson:"_id"`
	BaseID      string `bson:"base_id"`
	BoardID     string `bson:"board_id"`
	Title       string `bson:"title"`
	UnreadCount int    `bson:"unread_count"`
	//NotificationOn bool      `bson:"notification_on"`
	LastMsg       string    `bson:"last_msg"`
	LastMsgTime   time.Time `bson:"last_msg_time"`
	LastMsgSender string    `bson:"last_msg_sender"`
	Avatar        string    `bson:"avatar"`
	CreatorID     string    `bson:"creator_id"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
	Members       []string  `bson:"members"`
}

type Entity_Message struct {
	ID         string    `bson:"_id"`
	ChatroomID string    `bson:"chatroom_id"`
	SenderID   string    `bson:"sender_id"`
	Type       string    `bson:"type"`
	Data       string    `bson:"data"`
	CTime      time.Time `bson:"ctime"`
	UTime      time.Time `bson:"utime"`
}
