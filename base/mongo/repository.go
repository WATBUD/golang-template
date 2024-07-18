package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mai.today/base/mongo/collection"
	"mai.today/core/entity"
	"mai.today/database/mongodb"
)

type Client interface {
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Database(opts ...*options.DatabaseOptions) Database
}

type Database interface {
	BaseCollection(opts ...*options.CollectionOptions) *collection.BaseCollection
	InfoCollection(opts ...*options.CollectionOptions) *collection.BaseInfoCollection
	MemberCollection(opts ...*options.CollectionOptions) *collection.BaseMemberCollection
	NavStateCollection(opts ...*options.CollectionOptions) *collection.BaseNavStateCollection
}

type BaseRepository struct {
	Client
}

func NewBaseRepository(c *mongo.Client) *BaseRepository {
	client := &mongodb.DefaultDatabaseClient[Database]{
		Client: c,
		CreateDatabase: func(db *mongo.Database) Database {
			return &BaseDatabase{db}
		},
	}

	return &BaseRepository{
		Client: client,
	}
}

func (r BaseRepository) FindMemberByBaseID(ctx context.Context, id string) (entity *entity.BaseMember, err error) {
	return r.Client.Database().MemberCollection().FindByBaseID(ctx, id)
}

func (r BaseRepository) FindNavStatesByUserID(ctx context.Context, userID string) ([]*entity.BaseNavState, error) {
	return r.Client.Database().NavStateCollection().FindByUserID(ctx, userID)
}

func (r BaseRepository) InsertOne(ctx context.Context, member *entity.BaseMember, info *entity.BaseInfo, navState *entity.BaseNavState) (err error) {
	session, err := r.Client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		db := r.Client.Database()

		base := &entity.Base{}
		err := db.BaseCollection().InsertOne(ctx, base)
		if err != nil {
			return nil, err
		}

		info.BaseID = base.ID
		err = db.InfoCollection().InsertOne(ctx, info)
		if err != nil {
			return nil, err
		}

		member.BaseID = base.ID
		err = db.MemberCollection().InsertOne(ctx, member)
		if err != nil {
			return nil, err
		}

		navState.BaseID = base.ID
		err = db.NavStateCollection().InsertOne(ctx, navState)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (r BaseRepository) SoftDeleteByID(ctx context.Context, id string) (err error) {
	session, err := r.Client.StartSession()
	if err != nil {
		return err
	}

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		db := r.Client.Database()

		_, err = db.BaseCollection().SoftDeleteByID(ctx, id)
		if err != nil {
			return nil, err
		}

		_, err = db.InfoCollection().SoftDeleteByBaseID(ctx, id)
		if err != nil {
			return nil, err
		}

		_, err = db.MemberCollection().SoftDeleteByBaseID(ctx, id)
		if err != nil {
			return nil, err
		}

		_, err = db.NavStateCollection().SoftDeleteByBaseID(ctx, id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (r BaseRepository) UpdateInfoByBaseID(ctx context.Context, info *entity.BaseInfo) (bool, error) {
	return r.Client.Database().InfoCollection().UpdateByBaseID(ctx, info)
}

func (r BaseRepository) UpdateNavStates(ctx context.Context, items []*entity.BaseNavState) (err error) {
	session, err := r.Client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		c := r.Client.Database().NavStateCollection()

		for _, item := range items {
			if _, err := c.UpdateByID(ctx, item.ID, item); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})
	return err
}
