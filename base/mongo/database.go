package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mai.today/base/mongo/collection"
)

// mockgen -source=base/service.go -destination=base/test/mock_service.go -package=test

// BaseDatabase represents a wrapper around a mongo.Database instance, providing access to
// commonly used collections and simplifying database interactions.
type BaseDatabase struct {
	// Database embeds the underlying mongo.Database instance.
	*mongo.Database
}

// BaseCollection returns a new instance of BaseCollection for the base collection
// within the wrapped database. You can provide optional CollectionOptions for further configuration.
func (db *BaseDatabase) BaseCollection(opts ...*options.CollectionOptions) *collection.BaseCollection {
	return collection.NewBaseCollection(db.Database, opts...)
}

// InfoCollection returns a new instance of BaseInfoCollection for the "info" collection
// within the wrapped database. You can provide optional CollectionOptions for further configuration.
func (db *BaseDatabase) InfoCollection(opts ...*options.CollectionOptions) *collection.BaseInfoCollection {
	return collection.NewBaseInfoCollection(db.Database, opts...)
}

// MemberCollection returns a new instance of BaseMemberCollection for the "members" collection
// within the wrapped database. You can provide optional CollectionOptions for further configuration.
func (db *BaseDatabase) MemberCollection(opts ...*options.CollectionOptions) *collection.BaseMemberCollection {
	return collection.NewBaseMemberCollection(db.Database, opts...)
}

// NavStateCollection returns a new instance of BaseNavStateCollection for the "nav_states" collection
// within the wrapped database. You can provide optional CollectionOptions for further configuration.
func (db *BaseDatabase) NavStateCollection(opts ...*options.CollectionOptions) *collection.BaseNavStateCollection {
	return collection.NewBaseNavStateCollection(db.Database, opts...)
}
