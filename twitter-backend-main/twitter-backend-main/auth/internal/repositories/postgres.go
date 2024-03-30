package repositories

import (
	"context"

	"github.com/leoantony72/twitter-backend/auth/internal/ports"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	//"redis"

)

var (
	sadadsadsa 


)


var ctx = context.Background()

type UserPostgresRepo struct {
	db    *gorm.DB
	redis iRepo
}

type iRepo interface {
	Exists(ctx context.Context, keys ...string) *redis.IntCmd //1 TDIS.Exists
}

func NewUserPostgresRepo(db *gorm.DB, redis iRepo) ports.UserRepository { //iRepo TEMP

	return &UserPostgresRepo{
		db:    db,
		redis: redis,
	}
	
}


UserPostgresRepo.redis.DoesKeyExist22222

func DoesKeyExist22222(u iRepo, key string) bool {
	exists := u.Exists(ctx, key).Val()
	return exists != 0
}
