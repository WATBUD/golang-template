package repository

import (
	"context"
	"fmt"
	"folder_API/internal/entities"
	"folder_API/internal/usecases"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFolderRepository struct {
	collection *mongo.Collection
}

func NewMongoFolderRepository(client *mongo.Client) usecases.FolderRepository {
	collection := client.Database("mai_dev").Collection("folders")
	return &MongoFolderRepository{collection: collection}
}

func (r *MongoFolderRepository) Create(ctx context.Context, folder *entities.Folder) error {
	// Generate a new ObjectID if not already set
	// if folder.BaseID.IsZero() {
	// 	folder.BaseID = primitive.NewObjectID()
	// }
	if folder.BaseID == "" {
		return fmt.Errorf("BaseID is required and cannot be empty")
	}

	now := time.Now()
	folder.CreatedAt = now
	folder.UpdatedAt = now
	// Check for duplicate combination of parentIndex and folderIndex
	filter := bson.M{
		"parentIndex": folder.ParentIndex,
		"folderIndex": folder.FolderIndex,
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("combination of parentIndex %d and folderIndex %d already exists", folder.ParentIndex, folder.FolderIndex)
	}

	_, err = r.collection.InsertOne(ctx, folder)
	return err
}

func (r *MongoFolderRepository) FindAll(ctx context.Context) ([]*entities.Folder, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var folders []*entities.Folder
	if err := cursor.All(ctx, &folders); err != nil {
		return nil, err
	}
	return folders, nil
}

func (r *MongoFolderRepository) FindByID(ctx context.Context, id string) (*entities.Folder, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var folder entities.Folder
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&folder); err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *MongoFolderRepository) Update(ctx context.Context, folder *entities.Folder) error {
	objID, err := primitive.ObjectIDFromHex(folder.BaseID)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": folder})
	return err
}

func (r *MongoFolderRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *MongoFolderRepository) UpdateIndex(ctx context.Context, id string, folderIndex int) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"folderIndex": folderIndex}})
	return err
}

func (r *MongoFolderRepository) UpdateParent(ctx context.Context, id string, parentIndex int) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"parentIndex": parentIndex}})
	return err
}
