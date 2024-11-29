package services

import (
	"context"
	"encoding/json"
	"fmt"
	"food-delivery-service/consts"
	"food-delivery-service/repositories"
	"food-delivery-service/utils"
)

type riderService struct {
	riderRepo repositories.RiderRepository
}

func NewRiderService(riderRepo repositories.RiderRepository) RiderService {
	return riderService{riderRepo}
}

func (s riderService) GetRiders() (riderWrapper *RiderWrapper, err error) {
	key := "service::GetRiders"
	redisClient := utils.GetRedisClient()

	// Redis Get
	if riderJson, err := redisClient.Get(context.Background(), key).Result(); err == nil {
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
		redisClient.Set(context.Background(), key, string(data), consts.CACHE_TIME)
	}

	// fmt.Println("Query from database at service - GetRiders()", riderWrapperDB)
	fmt.Println("view rider from database")

	return &riderWrapperDB, nil
}
