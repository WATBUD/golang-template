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

func (r *MongoFolderRepository) Create(ctx context.Context, folder *entities.Folder) error {
	// Generate a new ObjectID if not already set
	if folder.ID.IsZero() {
		folder.ID = primitive.NewObjectID()
	}

	// Check for duplicate combination of parentIndex and index
	filter := bson.M{
		"parentIndex": folder.ParentIndex,
		"index":       folder.Index,
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("combination of parentIndex %d and index %d already exists", folder.ParentIndex, folder.Index)
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
	objID, err := primitive.ObjectIDFromHex(folder.ID.Hex())
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

func (r *MongoFolderRepository) UpdateIndex(ctx context.Context, id string, index int) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"index": index}})
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
