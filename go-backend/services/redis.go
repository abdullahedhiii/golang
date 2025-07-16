package caching

import (
	"context"
	"fmt"
	"time"

	"github.com/abdullahedhiii/go-backend/models"
	redis "github.com/go-redis/redis/v8"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx = context.Background()
)

func GetBook(id string) (map[string]string, error) {
	cacheKey := fmt.Sprintf("book:%s", id)
	res, err := client.HGetAll(ctx, cacheKey).Result()
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return res, nil
}

func AddBook(book models.Book) {
	cacheKey := fmt.Sprintf("book:%s", book.ID)
	_, err := client.HSet(ctx, cacheKey, map[string]interface{}{
		"Title":  book.Title,
		"Author": book.Author,
	}).Result()
	if err == nil {
		client.Expire(ctx, cacheKey, 10*time.Minute)
	}
}
