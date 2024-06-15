package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type UserRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) *UserRepo {
	return &UserRepo{
		client: client,
	}
}

func (r *UserRepo) Database(opts ...*options.CollectionOptions) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	return r.client.Database("mai_dev").Collection("users", opts...), ctx, cancel
}

func (r *UserRepo) FindUserByFireBaseUID(fireBaseUID string) *UserModel {
	var user *UserModel
	filter := bson.M{"user_info.firebase_uid": fireBaseUID}
	db, ctx, cancel := r.Database()
	defer cancel()
	result := db.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		switch {
		case strings.Contains(err.Error(), "no documents in result"):
			return nil
		default:
			fmt.Println(err)
		}
	}

	err := result.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.UserInfo.Nickname)
	return user
}

func (r *UserRepo) CreateUser(user UserModel) error {
	db, ctx, cancel := r.Database()
	defer cancel()
	_, err := db.InsertOne(ctx, user)
	return err
}
