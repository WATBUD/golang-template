package collection

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"mai.today/core/entity"
)

func TestBaseNavStateCollection_FindByUserID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("find by user id", func(mt *mtest.T) {
		collection := NewBaseNavStateCollection(mt.DB)
		id := primitive.NewObjectID()
		expected := []*entity.BaseNavState{
			{ID: id.Hex(), UserID: "user1", Index: 10},
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test_db.base_nav_states", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: id},
			{Key: "UserID", Value: expected[0].UserID},
			{Key: "Index", Value: expected[0].Index},
		}),
			mtest.CreateCursorResponse(0, "test_db.base_nav_states", mtest.NextBatch),
		)

		actual, err := collection.FindByUserID(context.Background(), "user1")
		assert.NoError(t, err)
		assert.Equal(t, expected[0].ID, actual[0].ID)
		assert.Equal(t, expected[0].UserID, actual[0].UserID)
		assert.Equal(t, expected[0].Index, actual[0].Index)
	})

	mt.Run("find by user id no result", func(mt *mtest.T) {
		collection := NewBaseNavStateCollection(mt.DB)
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "test_db.base_nav_state", mtest.FirstBatch))

		actual, err := collection.FindByUserID(context.Background(), "user2")
		assert.NoError(t, err)
		assert.Empty(t, actual)
	})
}

func TestBaseNavStateCollection_InsertOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful insert", func(mt *mtest.T) {
		collection := NewBaseNavStateCollection(mt.Client.Database("test_db"))
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		defaultID := "test id"
		e := &entity.BaseNavState{
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
		collection := NewBaseNavStateCollection(mt.Client.Database("test_db"))
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))
		e := &entity.BaseNavState{}

		err := collection.InsertOne(context.Background(), e)
		assert.Error(t, err)
		assert.Empty(t, e.ID)
		assert.NotZero(t, e.CreatedAt)
		assert.NotZero(t, e.UpdatedAt)
	})
}

func TestBaseNavStateCollection_UpdateByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful update", func(mt *mtest.T) {
		collection := NewBaseNavStateCollection(mt.DB)
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "n", Value: 1}}...))

		e := &entity.BaseNavState{
			UserID: "user1",
			Index:  2,
		}
		updated, err := collection.UpdateByID(context.Background(), "507f191e810c19729de860ea", e)
		assert.NoError(t, err)
		assert.True(t, updated)
	})

	mt.Run("update failure", func(mt *mtest.T) {
		collection := NewBaseNavStateCollection(mt.DB)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		e := &entity.BaseNavState{
			UserID: "user1",
			Index:  2,
		}
		updated, err := collection.UpdateByID(context.Background(), "507f191e810c19729de860ea", e)
		assert.Error(t, err)
		assert.False(t, updated)
	})
}
