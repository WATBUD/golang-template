package user

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	userRepo IUserRepo
}

func NewUser(mongo *mongo.Client) *User {
	return &User{
		userRepo: NewUserRepo(mongo),
	}
}

func (u *User) FindUserByFireBaseUID(fireBaseUID string) *UserModel {
	userModel := u.userRepo.FindUserByFireBaseUID(fireBaseUID)
	return userModel
}

func (u *User) CreateUser(firebaseUID string) error {
	//TODO
	picture := &Picture{
		Path:       "https://www.dora-world.com.tw/dist/images/character_1.png",
		UpdateTime: time.Now().UTC(),
	}
	newModel := UserModel{
		UserInfo: UserInfo{
			FirebaseUID: firebaseUID,
			Nickname:    "小強",
			Avatar:      *picture,
			Status:      RegisterProcess,
			CreatedTime: time.Now().UTC(),
		},
		UserBase: UserBase{
			BaseID: []string{"test1", "test2"},
		},
	}
	return u.userRepo.CreateUser(newModel)
}
