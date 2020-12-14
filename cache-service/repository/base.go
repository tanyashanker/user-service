package repository

import (
	"cache-service/common"
	"cache-service/models"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBRepo struct {
	MongoClient *mongo.Client
	RedisClient *redis.Client
}

var DBClient *DBRepo

//Creating connection with MongoDB
func (dc *DBRepo) MongoConnect(config *models.Secret) error {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := fmt.Sprintf("mongodb+srv://%s", config.Host)

	client := options.Client().
		SetAppName(config.DBName).
		SetMaxConnIdleTime(time.Microsecond * 100000).
		SetAuth(options.Credential{
			Username: config.UserName,
			Password: config.Password,
		}).
		ApplyURI(uri)
	log.Println(uri)

	dc.MongoClient, err = mongo.Connect(ctx, client)
	if err != nil {
		log.Fatalf("Unable to connect MongoDB %v", err)
		return err
	}
	log.Printf("Mongodb started at %s PORT", config.Port)
	return nil
}

//Creating connection with Redis Cache
func (dc *DBRepo) RedisConnect(config *models.Secret) (*redis.Client, error) {
	var err error

	addr := fmt.Sprintf("%v:%v", config.Host, config.Port)
	dbName, _ := strconv.Atoi(config.DBName)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		DB:           dbName,
		ReadTimeout:  time.Duration(common.REDIS_READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(common.REDIS_WRITE_TIMEOUT) * time.Second,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatalf("Unable to connect Redis %v", err)
		return nil, err
	}
	log.Printf("connected to redis cache....................")
	return client, nil
}
