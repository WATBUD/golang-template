package collection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mai.today/core/entity"
	"mai.today/core/entity/basemember"
	"mai.today/database/mongodb"
)

type BaseMemberCollection struct {
	*mongodb.BasicCollection[entity.BaseMember]
}

func NewBaseMemberCollection(db *mongo.Database, opts ...*options.CollectionOptions) *BaseMemberCollection {
	return &BaseMemberCollection{
		BasicCollection: &mongodb.BasicCollection[entity.BaseMember]{Collection: db.Collection(basemember.Table, opts...)},
	}
}

func (c *BaseMemberCollection) InsertOne(ctx context.Context, e *entity.BaseMember) error {
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
