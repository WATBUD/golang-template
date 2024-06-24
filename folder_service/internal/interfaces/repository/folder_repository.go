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

func (r *MongoFolderRepository) CreateFolder(ctx context.Context, folder *entities.Folder) error {
	if folder.BaseID == "" {
		return fmt.Errorf("BaseID is required and cannot be empty")
	}
	now := time.Now()
	folder.CreatedAt = now
	folder.UpdatedAt = now
	if folder.ChildIDs == nil {
		folder.ChildIDs = []string{}
	}
	// 插入新文件夾
	insertResult, err := r.collection.InsertOne(ctx, folder)
	if err != nil {
		return fmt.Errorf("failed to insert folder: %w", err)
	}

	// 檢查是否提供了 ParentID
	if folder.ParentID != nil && *folder.ParentID != "" {
		newFolderID := insertResult.InsertedID

		// 更新父文件夾的 ChildIDs
		filter := bson.M{"_id": folder.ParentID}
		update := bson.M{
			"$push": bson.M{
				"ChildIDs": newFolderID,
			},
			"$set": bson.M{
				"UpdatedAt": now,
			},
		}
		_, err := r.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return fmt.Errorf("failed to update parent folder's ChildIDs: %w", err)
		}
	}

	return nil
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
