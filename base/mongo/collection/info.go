package collection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mai.today/core/entity"
	"mai.today/core/entity/baseinfo"
	"mai.today/database/mongodb"
)

type BaseInfoCollection struct {
	*mongodb.BasicCollection[entity.BaseInfo]
}

func NewBaseInfoCollection(db *mongo.Database, opts ...*options.CollectionOptions) *BaseInfoCollection {
	return &BaseInfoCollection{
		BasicCollection: &mongodb.BasicCollection[entity.BaseInfo]{Collection: db.Collection(baseinfo.Table, opts...)},
	}
}

func (c *BaseInfoCollection) InsertOne(ctx context.Context, e *entity.BaseInfo) error {
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

func (c *BaseInfoCollection) UpdateByBaseID(ctx context.Context, e *entity.BaseInfo) (bool, error) {
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = e.CreatedAt

	result, err := c.UpdateOne(ctx, bson.M{
		"base_id": e.BaseID,
	}, bson.M{
		"$set": e,
	})
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}
