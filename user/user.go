package user

import (
	"time"

	"github.com/lithammer/shortuuid/v4"
	"go.mongodb.org/mongo-driver/mongo"
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
			UserID:      shortuuid.New(),
			Nickname:    "小強",
			Avatar:      *picture,
			FirebaseUID: []string{firebaseUID},
			Status:      IsRegister,
			CreatedTime: time.Now().UTC(),
		},
		UserBase: UserBase{
			base_id: []string{"test1", "test2"},
		},
	}
	return u.userRepo.CreateUser(newModel)
}
