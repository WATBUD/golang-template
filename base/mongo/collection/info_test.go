package collection

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"mai.today/core/entity"
)

func TestBaseInfoCollection_InsertOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful insert", func(mt *mtest.T) {
		collection := NewBaseInfoCollection(mt.Client.Database("test_db"))
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		defaultID := "test id"
		e := &entity.BaseInfo{
			ID: defaultID,
		}

		err := collection.InsertOne(context.Background(), e)
		assert.NoError(t, err)
		assert.NotEqual(t, defaultID, e.ID)
		assert.NotEmpty(t, e.ID)
		assert.NotZero(t, e.CreatedAt)
		assert.NotZero(t, e.UpdatedAt)
	})

	mt.Run("insert failure", func(mt *mtest.T) {
		collection := NewBaseInfoCollection(mt.Client.Database("test_db"))
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))
		e := &entity.BaseInfo{}

		err := collection.InsertOne(context.Background(), e)
		assert.Error(t, err)
		assert.Empty(t, e.ID)
		assert.NotZero(t, e.CreatedAt)
		assert.NotZero(t, e.UpdatedAt)
	})
}

func TestBaseInfoCollection_UpdateByBaseID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful update", func(mt *mtest.T) {
		collection := NewBaseInfoCollection(mt.DB)
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "n", Value: 1}}...))

		e := &entity.BaseInfo{
			BaseID: "test_base_id",
		}
		updated, err := collection.UpdateByBaseID(context.Background(), e)
		assert.NoError(t, err)
		assert.True(t, updated)
		assert.NotZero(t, e.CreatedAt)
		assert.NotZero(t, e.UpdatedAt)
	})

	mt.Run("update failure", func(mt *mtest.T) {
		collection := NewBaseInfoCollection(mt.DB)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		e := &entity.BaseInfo{
			BaseID: "test_base_id",
		}
		updated, err := collection.UpdateByBaseID(context.Background(), e)
		assert.Error(t, err)
		assert.False(t, updated)
	})
}
