package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// Sample entity for testing
type Sample struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BaseID    string             `bson:"base_id,omitempty"`
	Name      string             `bson:"name"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}

func TestBasicCollection_InsertOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful insert", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		entity := &Sample{Name: "Test"}
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		id, err := basicCollection.InsertOne(context.Background(), entity)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	mt.Run("insert failure", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		entity := &Sample{Name: "Test"}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		id, err := basicCollection.InsertOne(context.Background(), entity)
		assert.Error(t, err)
		assert.Empty(t, id)
	})
}

func TestBasicCollection_FindByBaseID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful find", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "DBName.CollectionName", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "base_id", Value: "base1"},
			{Key: "name", Value: "Test"},
		}))

		result, err := basicCollection.FindByBaseID(context.Background(), "base1")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "base1", result.BaseID)
		assert.Equal(t, "Test", result.Name)
	})

	mt.Run("find failure", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "DBName.CollectionName", mtest.FirstBatch))

		result, err := basicCollection.FindByBaseID(context.Background(), "base1")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestBasicCollection_FindByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful find", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		objectID := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "DBName.CollectionName", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: objectID},
			{Key: "name", Value: "Test"},
		}))

		result, err := basicCollection.FindByID(context.Background(), objectID.Hex())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, objectID, result.ID)
		assert.Equal(t, "Test", result.Name)
	})

	mt.Run("find failure", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		objectID := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "DBName.CollectionName", mtest.FirstBatch))

		result, err := basicCollection.FindByID(context.Background(), objectID.Hex())
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestBasicCollection_UpdateByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful update", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		objectID := primitive.NewObjectID()
		entity := &Sample{Name: "Updated Test"}
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "n", Value: 1}}...))

		matched, err := basicCollection.UpdateByID(context.Background(), objectID.Hex(), entity)
		assert.NoError(t, err)
		assert.True(t, matched)
	})

	mt.Run("update failure", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		objectID := primitive.NewObjectID()
		entity := &Sample{Name: "Updated Test"}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		matched, err := basicCollection.UpdateByID(context.Background(), objectID.Hex(), entity)
		assert.Error(t, err)
		assert.False(t, matched)
	})
}

func TestBasicCollection_SoftDeleteByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful soft delete", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		objectID := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "n", Value: 1}}...))

		deletedCount, err := basicCollection.SoftDeleteByID(context.Background(), objectID.Hex())
		assert.NoError(t, err)
		assert.Equal(t, 1, deletedCount)
	})

	mt.Run("soft delete failure", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		objectID := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		deletedCount, err := basicCollection.SoftDeleteByID(context.Background(), objectID.Hex())
		assert.Error(t, err)
		assert.Equal(t, 0, deletedCount)
	})
}

func TestBasicCollection_SoftDeleteByBaseID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful soft delete", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		baseID := "base1"
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "n", Value: 1}}...))

		deletedCount, err := basicCollection.SoftDeleteByBaseID(context.Background(), baseID)
		assert.NoError(t, err)
		assert.Equal(t, 1, deletedCount)
	})

	mt.Run("soft delete failure", func(mt *mtest.T) {
		collection := mt.Coll
		basicCollection := &BasicCollection[Sample]{Collection: collection}

		baseID := "base1"
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		deletedCount, err := basicCollection.SoftDeleteByBaseID(context.Background(), baseID)
		assert.Error(t, err)
		assert.Equal(t, 0, deletedCount)
	})
}
