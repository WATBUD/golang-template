package rbacShell2

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/open-policy-agent/opa/rego"
	"go.mongodb.org/mongo-driver/bson"
	"mai.today/database"
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	mongoClient, err := database.NewMongoClient()
	if err != nil {
		panic(err)
	}

	// // 連接到MongoDB
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	rolesCollection := mongoClient.Database("rbac_db").Collection("roles")
	usersCollection := mongoClient.Database("rbac_db").Collection("users")

	// 從MongoDB讀取角色和權限
	var rolesData []bson.M
	cursor, err := rolesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &rolesData); err != nil {
		log.Fatal(err)
	}

	roles := make(map[string][]string)
	for _, role := range rolesData {
		roleName := role["role"].(string)
		permissions := role["permissions"].(bson.A)
		for _, perm := range permissions {
			roles[roleName] = append(roles[roleName], perm.(string))
		}
	}

	// 從MongoDB讀取用戶及其角色
	var usersData []bson.M
	cursor, err = usersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &usersData); err != nil {
		log.Fatal(err)
	}

	users := make(map[string][]string)
	for _, user := range usersData {
		userName := user["user"].(string)
		userRoles := user["roles"].(bson.A)
		for _, role := range userRoles {
			users[userName] = append(users[userName], role.(string))
		}
	}

	// 定義輸入數據
	input := map[string]interface{}{
		"user":        "alice",
		"action":      "create_channel",
		"roles":       users,
		"permissions": roles,
	}

	// 加載Rego策略並準備查詢
	query, err := rego.New(
		rego.Query("data.rbac.allow"),
		rego.Load([]string{"policy.rego"}, nil), // 使用Rego策略文件
	).PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 評估Rego策略
	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	// 輸出結果
	if len(rs) > 0 && rs[0].Expressions[0].Value.(bool) {
		fmt.Println("Access granted") // 如果策略允許該操作，輸出 "Access granted"
	} else {
		fmt.Println("Access denied") // 否則，輸出 "Access denied"
	}
}
