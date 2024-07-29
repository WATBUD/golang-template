package repository

import (
	"chat_room_mod/core/domain/chatroom"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoChatroomRepository struct {
	chatroomCollection *mongo.Collection
	messageCollection  *mongo.Collection
}

func NewMongoChatroomRepository(client *mongo.Client) chatroom.ChatroomRepository {
	chatroomCollection := client.Database("mai_dev").Collection("chatrooms")
	messageCollection := client.Database("mai_dev").Collection("messages")
	return &MongoChatroomRepository{
		chatroomCollection: chatroomCollection,
		messageCollection:  messageCollection,
	}
}

func (r *MongoChatroomRepository) CreateChatroom(chatroom *chatroom.Entity_Chatroom) (*chatroom.Entity_Chatroom, error) {
	newID := primitive.NewObjectID()
	chatroom.ID = newID.Hex()
	result, err := r.chatroomCollection.InsertOne(context.TODO(), chatroom)
	if err != nil {
		fmt.Printf("Error inserting document: %v\n", err)
		return nil, err
	}
	fmt.Printf("result: %v\n", result)
	return chatroom, nil
}

func (r *MongoChatroomRepository) CheckChatroomId(id string) (*chatroom.Entity_Chatroom, error) {
	var _chatroom chatroom.Entity_Chatroom
	filter := bson.M{"_id": id}
	err := r.chatroomCollection.FindOne(context.TODO(), filter).Decode(&_chatroom)
	if err != nil {
		return nil, err
	}
	return &_chatroom, nil
}

func (r *MongoChatroomRepository) GetChatrooms() ([]*chatroom.Entity_Chatroom, error) {
	filter := bson.D{{}}
	options := options.Find().SetSort(bson.D{{Key: "last_msg_time", Value: -1}})
	cur, err := r.chatroomCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var __chatrooms []*chatroom.Entity_Chatroom
	for cur.Next(context.TODO()) {
		var _chatroom chatroom.Entity_Chatroom
		err := cur.Decode(&_chatroom)
		if err != nil {
			return nil, err
		}
		__chatrooms = append(__chatrooms, &_chatroom)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return __chatrooms, nil
}

func (r *MongoChatroomRepository) UpdateChatroom(chatroom *chatroom.Entity_Chatroom) error {
	filter := bson.M{"_id": chatroom.ID}
	update := bson.M{"$set": chatroom}
	_, err := r.chatroomCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoChatroomRepository) SendMessage(message *chatroom.Entity_Message) (*chatroom.Entity_Message, error) {

	newID := primitive.NewObjectID()
	message.ID = newID.Hex()
	result, err := r.messageCollection.InsertOne(context.TODO(), message)
	if err != nil {
		fmt.Printf("Error inserting document: %v\n", err)
		return nil, err
	}
	fmt.Printf("result: %v\n", result)
	return message, nil
}

func (r *MongoChatroomRepository) GetMessages(chatroomID string) ([]*chatroom.Entity_Message, error) {

	filter := bson.M{"chatroom_id": chatroomID}
	options := options.Find().SetSort(bson.D{{Key: "ctime", Value: -1}})
	cur, err := r.messageCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var messages []*chatroom.Entity_Message
	for cur.Next(context.TODO()) {
		var _message chatroom.Entity_Message
		err := cur.Decode(&_message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &_message)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MongoChatroomRepository) RemoveUserFromChatroom(chatroomID string, userID string) error {
	filter := bson.M{"_id": chatroomID}
	update := bson.M{
		"$pull": bson.M{"members": userID},
	}

	_, err := r.chatroomCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
