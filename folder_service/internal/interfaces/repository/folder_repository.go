package repository

import (
	"context"
	"fmt"
	"folder_API/internal/entities"
	"folder_API/internal/usecases"

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

func (r *MongoFolderRepository) CreateFolder(ctx context.Context, folder *entities.Folder) (*entities.Folder, error) {
	// Insert the new folder
	insertResult, err := r.collection.InsertOne(ctx, folder)
	if err != nil {
		return nil, fmt.Errorf("failed to insert folder: %w", err)
	}
	// Retrieve the inserted document
	var insertedFolder entities.Folder
	err = r.collection.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&insertedFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve inserted folder: %w", err)
	}
	// newFolderID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	// folder.ID = newFolderID
	return &insertedFolder, nil
}

func (r *MongoFolderRepository) GetFolders(ctx context.Context) ([]*entities.Folder, error) {
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

func (r *MongoFolderRepository) UpdateFolderData(ctx context.Context, folder *entities.Folder) error {
	objectID, err := primitive.ObjectIDFromHex(folder.ID)
	if err != nil {
		return fmt.Errorf("invalid parent_id: %w", err)
	}
	filter := bson.M{"_id": objectID, "base_id": folder.BaseID}
	update := bson.M{
		"$set": bson.M{
			"parent_id":  folder.ParentID,
			"position":   folder.Position,
			"data":       folder.Data,
			"updated_at": folder.UpdatedAt,
		},
	}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update folder: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no matching document found for ID: %s and BaseID: %s", folder.ID, folder.BaseID)
	}
	return nil
}

func (r *MongoFolderRepository) DeleteFolder(ctx context.Context, folder *entities.Folder) error {
	objectID, err := primitive.ObjectIDFromHex(folder.ID)
	if err != nil {
		return fmt.Errorf("invalid parent_id: %w", err)
	}
	// Find the folder to be deleted based on the _id
	var folderToDelete entities.Folder
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&folderToDelete)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("folder not found")
		}
		return fmt.Errorf("failed to find folder: %w", err)
	}

	// Collect ObjectIDs that need to be deleted
	objIDs := []primitive.ObjectID{objectID} // Add the main folder's ObjectID to the delete list

	// Convert ChildIDs to ObjectIDs and add them to the delete list
	// for _, childID := range folderToDelete.ChildIDs {
	// 	childObjID, err := primitive.ObjectIDFromHex(childID)
	// 	if err != nil {
	// 		return fmt.Errorf("invalid child ID: %w", err)
	// 	}
	// 	objIDs = append(objIDs, childObjID)
	// }

	// Delete all collected ObjectIDs
	_, err = r.collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": objIDs}})
	return err
}

func (r *MongoFolderRepository) UpdateFolderParentID(ctx context.Context, objectID string, parentID string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"base_id": objectID},
		bson.M{"$set": bson.M{"parent_id": parentID}},
	)
	return err
}

func (r *MongoFolderRepository) AddChildIDToParent(ctx context.Context, parentID string, childID string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"base_id": parentID},
		bson.M{"$addToSet": bson.M{"child_ids": childID}},
	)
	return err
}

func (r *MongoFolderRepository) PositionExists(ctx context.Context, position float64) (bool, error) {
	filter := bson.M{"position": position}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
