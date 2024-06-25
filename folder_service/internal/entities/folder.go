package entities

import (
	"time"
)

// omitempty 標籤用於JSON 或 BSON 編碼器，
// struct的值為零值 對於指針類型為 nil，於數字類型為 0，
// 對於字符串類型為空字符串，該字段應該被省略。
// MongoDB 中，無法直接設定預設值
//方法：程式自己檢查處理/Triggers/validation rules/aggregation pipeline

type Folder struct {
	ID        string      `bson:"_id,omitempty"`
	BaseID    string      `bson:"base_id"`
	Name      string      `bson:"name"`
	Color     string      `bson:"color"`
	ParentID  string      `bson:"parent_id,omitempty"`
	ChildIDs  []string    `bson:"child_ids"`
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
	Data      interface{} `bson:"data"`
}
