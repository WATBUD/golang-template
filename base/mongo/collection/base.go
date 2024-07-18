package collection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mai.today/core/entity"
	"mai.today/core/entity/base"
	"mai.today/database/mongodb"
)

type BaseCollection struct {
	*mongodb.BasicCollection[entity.Base]
}

func NewBaseCollection(db *mongo.Database, opts ...*options.CollectionOptions) *BaseCollection {
	return &BaseCollection{
		BasicCollection: &mongodb.BasicCollection[entity.Base]{Collection: db.Collection(base.Table, opts...)},
	}
}

func (c *BaseCollection) InsertOne(ctx context.Context, e *entity.Base) error {
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
