package repository

import (
	"chat_room_mod/core/domain/chatroom"
	"context"
	"errors"
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

func (r *MongoChatroomRepository) GetChatroomByID(id string) (*chatroom.Entity_Chatroom, error) {
	var _chatroom chatroom.Entity_Chatroom
	filter := bson.M{"_id": id}
	err := r.chatroomCollection.FindOne(context.TODO(), filter).Decode(&_chatroom)
	if err != nil {
		return nil, err
	}
	return &_chatroom, nil
}

func (r *MongoChatroomRepository) GetChatrooms(userID string) ([]*chatroom.Entity_Chatroom, error) {
	// Filter to search for chatrooms where the PersonalConfig array contains an entry with the matching userID
	filter := bson.M{"personal_config.user_id": userID}

	// Set the sorting option based on "last_msg_time" in descending order
	options := options.Find().SetSort(bson.D{{Key: "last_msg_time", Value: -1}})

	// Perform the MongoDB find operation
	cur, err := r.chatroomCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	// Initialize a slice to hold the chatrooms
	var __chatrooms []*chatroom.Entity_Chatroom

	// Iterate through the cursor and decode each chatroom into the slice
	for cur.Next(context.TODO()) {
		var _chatroom chatroom.Entity_Chatroom
		err := cur.Decode(&_chatroom)
		if err != nil {
			return nil, err
		}
		__chatrooms = append(__chatrooms, &_chatroom)
	}

	// Check for any errors during the cursor iteration
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Return the slice of chatrooms and a nil error to indicate success
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

func (r *MongoChatroomRepository) GetChatroomPersonalConfig(chatroomID string) ([]*chatroom.ChatroomPersonalConfig, error) {
	filter := bson.M{"_id": chatroomID}

	var result struct {
		PersonalConfig []chatroom.ChatroomPersonalConfig `bson:"personal_config"`
	}

	err := r.chatroomCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("chatroom not found")
		}
		return nil, fmt.Errorf("failed to retrieve chatroom members: %w", err)
	}

	var members []*chatroom.ChatroomPersonalConfig
	for _, member := range result.PersonalConfig {
		members = append(members, &member)
	}

	return members, nil
}

func (r *MongoChatroomRepository) RemoveUserFromChatroom(dto chatroom.DTO_AddOrRemoveChatRoomUserRequest) error {
	filter := bson.M{
		"_id":                     dto.ChatroomID,
		"personal_config.user_id": dto.UserID,
	}

	update := bson.M{
		"$pull": bson.M{
			"personal_config": bson.M{
				"user_id": dto.UserID,
			},
		},
	}
	// Perform the update operation
	_, err := r.chatroomCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to remove user: %w", err) // Return formatted error if the update operation fails
	}

	// Successfully removed the user from the chatroom
	return nil
}

func (r *MongoChatroomRepository) SetUserChatRoomConfig(dto chatroom.DTO_SetUserChatRoomConfig) error {
	filter := bson.M{
		"_id":                     dto.ChatroomID,
		"personal_config.user_id": dto.UserID,
	}

	update := bson.M{
		"$set": bson.M{
			"personal_config.$." + dto.ParamName: dto.ParamValue,
		},
	}
	_, err := r.chatroomCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to mute user: %w", err)
	}

	return nil
}

func (r *MongoChatroomRepository) UserMutedChatRoom(dto chatroom.DTO_MuteUserRequest) error {
	filter := bson.M{
		"_id":                     dto.ChatroomID,
		"personal_config.user_id": dto.UserID,
	}

	update := bson.M{
		"$set": bson.M{
			"personal_config.$.is_muted": dto.Mute,
		},
	}
	_, err := r.chatroomCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to mute user: %w", err)
	}

	return nil
}

func (r *MongoChatroomRepository) AddUserToChatroom(chatroomID string, personal_config []*chatroom.ChatroomPersonalConfig) error {
	filter := bson.M{"_id": chatroomID}

	update := bson.M{
		"$set": bson.M{
			"personal_config": personal_config,
		},
	}

	_, err := r.chatroomCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update chatroom: %w", err)
	}

	return nil
}
