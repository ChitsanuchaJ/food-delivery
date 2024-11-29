package utils

import (
	"sync"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB
	restyClient  *resty.Client
	redisClient  *redis.Client
	producer     sarama.SyncProducer
	restyOnce    sync.Once
	redisOnce    sync.Once
	dbOnce       sync.Once
	producerOnce sync.Once
)

// Lazy initialize
func InitRestyClient() {
	restyOnce.Do(func() {
		restyClient = resty.New().
			SetBaseURL("http://localhost:8001").
			SetHeader("Content-Type", "application/json")
	})
}

func GetRestyClient() *resty.Client {
	if restyClient == nil {
		InitRestyClient()
	}
	return restyClient
}

//////////////////////////////////////////////////////

func InitDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/food-delivery")

	var err error
	dbOnce.Do(
		func() {
			db, err = gorm.Open(dial, &gorm.Config{})
			if err != nil {
				panic(err)
			}
		})

	return db
}

//////////////////////////////////////////////////////

func InitRedis() {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	})
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitRedis()
	}
	return redisClient
}

//////////////////////////////////////////////////////

func InitProducer() sarama.SyncProducer {
	var err error
	producerOnce.Do(
		func() {
			producer, err = sarama.NewSyncProducer([]string{"localhost:9093"}, nil)
			if err != nil {
				panic(err)
			}
		})

	return producer
}
