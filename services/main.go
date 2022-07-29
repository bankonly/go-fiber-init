package services

import (
	"context"
	"gofiber/configs"

	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceInstances struct {
	DBClient  *mongo.Client
	DBContext context.Context
	DB        *mongo.Database
	Redis     *redis.Client
}

func InitServiceInstances() ServiceInstances {
	return ServiceInstances{
		DBClient:  configs.DBClient,
		DBContext: configs.DBContext,
		DB:        configs.DB,
		Redis:     configs.RedisCli,
	}
}
