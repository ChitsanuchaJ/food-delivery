package services

import (
	"context"
	"encoding/json"
	"fmt"
	"food-delivery-services/consts"
	"food-delivery-services/repositories"

	"github.com/go-redis/redis/v8"
)

type riderService struct {
	riderRepo   repositories.RiderRepository
	redisClient *redis.Client
}

func NewRiderService(riderRepo repositories.RiderRepository, redisClient *redis.Client) RiderService {
	return riderService{riderRepo, redisClient}
}

func (s riderService) GetRiders() (riderWrapper *RiderWrapper, err error) {
	key := "service::GetRiders"

	// Redis Get
	if riderJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(riderJson), &riderWrapper) == nil {
			// fmt.Println("Cache hit from redis at service - GetRiders()", riderWrapper)
			fmt.Println("view rider from cached")
			return riderWrapper, nil
		}
	}

	ridersData, err := s.riderRepo.GetRiders()
	if err != nil {
		return nil, err
	}

	riders := []Rider{}
	for _, rider := range ridersData {
		riders = append(riders, Rider{
			ID:   rider.ID,
			Name: rider.Name,
		})
	}
	riderWrapperDB := RiderWrapper{riders}

	// Redis SET
	if data, err := json.Marshal(riderWrapperDB); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), consts.CACHE_TIME)
	}

	// fmt.Println("Query from database at service - GetRiders()", riderWrapperDB)
	fmt.Println("view rider from database")

	return &riderWrapperDB, nil
}
