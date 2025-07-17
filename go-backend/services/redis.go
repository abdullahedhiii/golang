package caching

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	ctx = context.Background()
)

func GET(prefix, id string) (map[string]string, error) {
	cacheKey := fmt.Sprintf("%s:%s", prefix, id)
	res, err := client.HGetAll(ctx, cacheKey).Result()
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return res, nil
}

func PUT(prefix string, id string, data interface{}, ttl time.Duration) error {
	cacheKey := fmt.Sprintf("%s:%s", prefix, id)

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("SetToCache expects a struct, got %s", v.Kind())
	}

	fields := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		value := v.Field(i).Interface()

		fields[name] = value
	}

	if _, err := client.HSet(ctx, cacheKey, fields).Result(); err != nil {
		return err
	}
	return client.Expire(ctx, cacheKey, ttl).Err()
}
