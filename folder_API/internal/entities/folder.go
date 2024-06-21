package entities

import (
	"time"
)

type Folder struct {
	BaseID      string    `bson:"base_id" json:"base_id"`
	Name        string    `bson:"name" json:"name"`
	Color       string    `bson:"color" json:"color"`
	FolderIndex int       `bson:"folderIndex" json:"folderIndex"`
	ParentIndex int       `bson:"parentIndex" json:"parentIndex"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}
