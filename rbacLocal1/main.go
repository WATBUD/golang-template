package main

import (
	"context"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	// 創建一個背景上下文
	ctx := context.Background()
	// 加載Rego策略並準備查詢
	query, err := rego.New(
		rego.Query("data.rbac.allow"),           // 定義查詢，檢查策略中的`allow`規則
		rego.Load([]string{"policy.rego"}, nil), // 加載Rego策略文件
	).PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err) // 如果加載或準備查詢時發生錯誤，記錄並退出程序
	}

	// 定義測試數據輸入
	input := map[string]interface{}{
		"user":   "alice",          // 測試用戶為 "alice"
		"action": "create_channel", // 測試操作為 "create_channel"
	}

	// 評估Rego策略
	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err) // 如果評估策略時發生錯誤，記錄並退出程序
	}

	// 輸出結果
	if len(rs) > 0 && rs[0].Expressions[0].Value.(bool) {
		fmt.Println("Access granted") // 如果策略允許該操作，輸出 "Access granted"
	} else {
		fmt.Println("Access denied") // 否則，輸出 "Access denied"
	}
}
