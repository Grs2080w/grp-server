package redis

// "github.com/Grs2080w/grp_server/core/db/redis"

import (
	"context"
	"errors"
	"time"

	c "github.com/Grs2080w/grp_server/core/config"
	"github.com/redis/go-redis/v9"
)

type ClientR struct {
	Client *redis.Client
	Ctx context.Context
}

var Client ClientR

func Init_Client_Redis() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     c.GetValueEnv("REDIS_URL"),
		Username: c.GetValueEnv("REDIS_USERNAME"),
		Password: c.GetValueEnv("REDIS_PASSWORD"),
		DB:       c.GetValueEnvInt("REDIS_DB"),
	})

	Client.Client = client
	Client.Ctx = ctx
}

func R_set(key string, value any, tll int) {
	Client.Client.Set(Client.Ctx, key, value, time.Duration(tll)*time.Second)
}


func R_set_json(key string, value any, tll int) {
	Client.Client.JSONMSet(Client.Ctx, key, value, time.Duration(tll)*time.Second)
}


func R_get(key string) (string, error){

	result, err := Client.Client.Get(Client.Ctx, key).Result()

	if err != nil {
		return "", errors.New("error on get key redis")
	}

	return result, nil
}

func R_del(key string) {
	Client.Client.Del(Client.Ctx, key)
}


