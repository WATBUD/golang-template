package entities

import (
	"time"
)

// ID          string    `bson:"_id,omitempty" json:"id,omitempty"` // 唯一标识符
type Folder struct {
	BaseID      string    `bson:"base_id"`
	BaseIndex   int       `bson:"BaseIndex"`
	Name        string    `bson:"name"`
	Color       string    `bson:"color"`
	FolderIndex int       `bson:"folderIndex"`
	ParentID    *string   `bson:"parent_id,omitempty"`
	ChildIDs    []string  `bson:"child_ids,omitempty"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}
