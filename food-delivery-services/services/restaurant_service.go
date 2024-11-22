package services

import (
	"context"
	"encoding/json"
	"fmt"
	"food-delivery-services/consts"
	"food-delivery-services/repositories"

	"github.com/go-redis/redis/v8"
)

type restaurantService struct {
	restaurantRepo repositories.RestaurantRepository
	redisClient    *redis.Client
}

func NewRestaurantService(restaurantRepo repositories.RestaurantRepository, redisClient *redis.Client) RestaurantService {
	return restaurantService{restaurantRepo, redisClient}
}

func (s restaurantService) GetRestaurants() (restaurantWrapper *RestaurantWrapper, err error) {
	key := "service::GetRestaurants"

	// Redis Get
	if restaurantJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(restaurantJson), &restaurantWrapper) == nil {
			// fmt.Println("Cache hit from redis at service - GetRestaurants()", restaurantWrapper)
			fmt.Println("view restaurant from cached")
			return restaurantWrapper, nil
		}
	}

	restaurantsData, err := s.restaurantRepo.GetRestaurants()
	if err != nil {
		return nil, err
	}

	restaurants := []Restaurant{}
	for _, restaurant := range restaurantsData {
		restaurants = append(restaurants, Restaurant{
			ID:   restaurant.ID,
			Name: restaurant.Name,
		})
	}
	restaurantWrapperDB := RestaurantWrapper{restaurants}

	// Redis SET
	if data, err := json.Marshal(restaurantWrapperDB); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), consts.CACHE_TIME)
	}

	// fmt.Println("Query from database at service - GetRestaurants()", restaurantWrapperDB)
	fmt.Println("view restaurant from database")

	return &restaurantWrapperDB, nil
}
