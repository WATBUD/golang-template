package user

import "time"

type Status string

const (
	RegisterProcess Status = "in_registration_process"
	Normal          Status = "normal"
	Delete          Status = "delete"
	Disable         Status = "disable"
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
	FirebaseUID        string    `bson:"firebase_uid" json:"firebase_uid"`
	Nickname           string    `bson:"nickname" json:"nickname"`
	Avatar             Picture   `bson:"avatar" json:"avatar"`
	Status             Status    `bson:"status" json:"status"`
	LastUpdatedTime    time.Time `bson:"last_updated_time,omitempty" json:"last_updated_time,omitempty"`
	LastInteractionDay time.Time `bson:"last_interaction_day,omitempty" json:"last_interaction_day,omitempty"`
	CreatedTime        time.Time `bson:"created_time" json:"created_time"`
}

type UserBase struct {
	BaseID []string `bson:"base_id" json:"base_id"`
}
