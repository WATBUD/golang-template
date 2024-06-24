package user

import "time"

type Status string

const (
	IsRegister Status = "is_register"
	Available  Status = "available"
	Delete     Status = "delete"
	Disable    Status = "disable"
)

type Picture struct {
	Path       string    `json:"path" bson:"path"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

type UserModel struct {
	UserInfo UserInfo `bson:"user_info" json:"user_info"`
	UserBase UserBase `bson:"user_base" json:"user_base"`
}

type UserInfo struct {
	UserID             string    `bson:"user_id" json:"user_id"`
	Nickname           string    `bson:"nickname" json:"nickname"` //TODO
	Avatar             Picture   `bson:"avatar" json:"avatar"`     //TODO
	FirebaseUID        []string  `bson:"firebase_uid" json:"firebase_uid"`
	Status             Status    `bson:"status" json:"status"` //TODO
	LastUpdatedTime    time.Time `bson:"last_updated_time,omitempty" json:"last_updated_time,omitempty"`
	LastInteractionDay time.Time `bson:"last_interaction_day,omitempty" json:"last_interaction_day,omitempty"`
	CreatedTime        time.Time `bson:"created_time" json:"created_time"`
}

type UserBase struct {
	base_id []string `bson:"base_id" json:"base_id"` //TODO
}
