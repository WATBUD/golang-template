package collection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mai.today/core/entity"
	"mai.today/core/entity/basenavstate"
	"mai.today/database/mongodb"
)

type BaseNavStateCollection struct {
	*mongodb.BasicCollection[entity.BaseNavState]
}

func NewBaseNavStateCollection(db *mongo.Database, opts ...*options.CollectionOptions) *BaseNavStateCollection {
	return &BaseNavStateCollection{
		BasicCollection: &mongodb.BasicCollection[entity.BaseNavState]{Collection: db.Collection(basenavstate.Table, opts...)},
	}
}

func (c *BaseNavStateCollection) FindByUserID(ctx context.Context, id string) ([]*entity.BaseNavState, error) {
	filter := bson.M{basenavstate.FieldUserID: id}
	opt := options.Find().SetSort(bson.M{basenavstate.FieldIndex: 1})

	cursor, err := c.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	entities := []*entity.BaseNavState{}
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}
	return entities, nil

}

func (c *BaseNavStateCollection) InsertOne(ctx context.Context, e *entity.BaseNavState) error {
	e.ID = ""
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = e.CreatedAt

	id, err := c.BasicCollection.InsertOne(ctx, e)
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func (c *BaseNavStateCollection) UpdateByID(ctx context.Context, id string, e *entity.BaseNavState) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	e.ID = ""
	result, err := c.UpdateOne(ctx, bson.M{
		"_id": objectID,
	}, bson.M{
		"$set": e,
	})
	e.ID = objectID.Hex()
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}
