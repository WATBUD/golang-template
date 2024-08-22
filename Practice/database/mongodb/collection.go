package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection[T] defines a generic interface for interacting with MongoDB collections.
// It provides basic CRUD (Create, Read, Update, Delete) operations for entities of type T,
// along with soft delete functionalities based on ID or a custom "base_id" field.
type Collection[T any] interface {
	// InsertOne inserts a new entity of type T into the collection.
	// It returns the inserted document's ID (as a string) and an error if any occurs.
	InsertOne(ctx context.Context, e *T) (string, error)

	// FindAllByBaseIDs retrieves all entities of type T matching the provided base_id slice.
	// It returns a slice of pointers to the entities and an error if any occurs.
	FindAllByBaseIDs(ctx context.Context, id []*string) ([]*T, error)

	// FindAllByUserID retrieves all entities of type T matching the provided user_id.
	// It returns a slice of pointers to the entities and an error if any occurs.
	FindAllByUserID(ctx context.Context, id string) ([]*T, error)

	// FindByBaseID retrieves an entity of type T from the collection by its custom "base_id".
	// It returns a pointer to the entity and an error if any occurs.
	FindByBaseID(ctx context.Context, id string) (*T, error)

	// FindByID retrieves an entity of type T from the collection by its ID.
	// It returns a pointer to the entity and an error if any occurs.
	FindByID(ctx context.Context, id string) (*T, error)

	// UpdateByID updates an existing entity of type T in the collection based on its ID.
	// It returns true if an entity was matched and updated, false otherwise, along with an error if any occurs.
	UpdateByID(ctx context.Context, id string, e *T) (bool, error)

	// SoftDeleteByID performs a soft delete on an entity identified by its ID.
	// It updates the entity's "deleted_at" field with the current UTC time and returns the number of documents affected (should be 1) and an error if any occurs.
	SoftDeleteByID(ctx context.Context, id string) (int, error)

	// SoftDeleteByBaseID performs a soft delete on an entity identified by a custom "base_id" field.
	// It updates the entity's "deleted_at" field with the current UTC time and returns the number of documents affected (should be 1) and an error if any occurs.
	SoftDeleteByBaseID(ctx context.Context, id string) (int, error)

	// SoftDeleteByBoardID performs a soft delete on an entity identified by a custom "board_id" field.
	// It updates the entity's "deleted_at" field with the current UTC time and returns the number of documents affected (should be 1) and an error if any occurs.
	SoftDeleteByBoardID(ctx context.Context, id string) (int, error)
}

// BasicCollection[T] is a concrete implementation of the Collection interface for MongoDB collections.
// It wraps a *mongo.Collection and provides methods for CRUD operations on entities of type T.
type BasicCollection[T any] struct {
	// Collection embeds a pointer to the underlying mongo.Collection instance.
	*mongo.Collection
}

// InsertOne implements the InsertOne method for BasicCollection.
// See the documentation for Collection.InsertOne for details.
func (c *BasicCollection[T]) InsertOne(ctx context.Context, e *T) (string, error) {
	result, err := c.Collection.InsertOne(ctx, e)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FindAllByBaseIDs retrieves all entities of type T matching the provided base_id slice.
// It returns a slice of pointers to the entities and an error if any occurs.
func (c *BasicCollection[T]) FindAllByBaseIDs(ctx context.Context, ids []string) (entities []*T, err error) {
	filter := bson.M{
		"base_id":    bson.M{"$in": ids},
		"deleted_at": nil,
	}
	opt := options.Find()

	cursor, err := c.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return
}

// FindAllByUserID retrieves all entities of type T matching the provided user_id.
// It returns a slice of pointers to the entities and an error if any occurs.
func (c *BasicCollection[T]) FindAllByUserID(ctx context.Context, id string) (entities []*T, err error) {
	filter := bson.M{
		"user_id":    id,
		"deleted_at": nil,
	}
	opt := options.Find()

	cursor, err := c.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return
}

// FindByBaseID implements the FindByBaseID method for BasicCollection.
// See the documentation for Collection.FindByBaseID for details.
func (c *BasicCollection[T]) FindByBaseID(ctx context.Context, id string) (*T, error) {
	filter := bson.M{"base_id": id}
	opt := options.FindOne()

	var entity T
	if err := c.FindOne(ctx, filter, opt).Decode(&entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindByID implements the FindByID method for BasicCollection.
// See the documentation for Collection.FindByID for details.
func (c *BasicCollection[T]) FindByID(ctx context.Context, id string) (*T, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entity T
	if err := c.FindOne(ctx, bson.M{
		"_id": objectID,
	}, options.FindOne().SetSort(bson.M{})).Decode(&entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

// UpdateByID implements the UpdateByID method for BasicCollection.
// See the documentation for Collection.UpdateByID for details.
func (c *BasicCollection[T]) UpdateByID(ctx context.Context, id string, e *T) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	result, err := c.UpdateOne(ctx, bson.M{
		"_id": objectID,
	}, bson.M{
		"$set": e,
	})
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

// SoftDeleteByID implements the SoftDeleteByID method for BasicCollection.
// It performs a soft delete by marking the entity identified by its ID as deleted
// by setting the "deleted_at" field to the current UTC time.
// See the documentation for Collection.SoftDeleteByID for details.
func (c *BasicCollection[T]) SoftDeleteByID(ctx context.Context, id string) (int, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := c.UpdateOne(ctx, bson.M{
		"_id": objectID,
	}, bson.M{
		"$set": bson.M{"deleted_at": time.Now().UTC()},
	})
	if err != nil {
		return 0, err
	}
	return int(result.MatchedCount), nil
}

// SoftDeleteByBaseID implements the SoftDeleteByBaseID method for BasicCollection.
// It performs a soft delete by marking the entity identified by a custom "base_id" field as deleted
// by setting the "deleted_at" field to the current UTC time.
// See the documentation for Collection.SoftDeleteByBaseID for details.
func (c *BasicCollection[T]) SoftDeleteByBaseID(ctx context.Context, id string) (int, error) {
	filter := bson.M{"base_id": id}

	result, err := c.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{"deleted_at": time.Now().UTC()},
	})
	if err != nil {
		return 0, err
	}
	return int(result.MatchedCount), nil
}

// SoftDeleteByBoardID implements the SoftDeleteByBoardID method for BasicCollection.
// It performs a soft delete by marking the entity identified by a custom "board_id" field as deleted
// by setting the "deleted_at" field to the current UTC time.
// See the documentation for Collection.SoftDeleteByBoardID for details.
func (c *BasicCollection[T]) SoftDeleteByBoardID(ctx context.Context, id string) (int, error) {
	filter := bson.M{"board_id": id}

	result, err := c.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{"deleted_at": time.Now().UTC()},
	})
	if err != nil {
		return 0, err
	}
	return int(result.MatchedCount), nil
}
