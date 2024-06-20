package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Folder struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Color       string             `bson:"color" json:"color"`
	Index       int                `bson:"index" json:"index"`
	ParentIndex int                `bson:"parentIndex" json:"parentIndex"`
}
