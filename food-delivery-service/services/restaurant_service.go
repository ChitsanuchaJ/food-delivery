package services

import (
	"context"
	"encoding/json"
	"fmt"
	"food-delivery-service/consts"
	"food-delivery-service/repositories"
	"food-delivery-service/utils"
)

type restaurantService struct {
	restaurantRepo repositories.RestaurantRepository
}

func NewRestaurantService(restaurantRepo repositories.RestaurantRepository) RestaurantService {
	return restaurantService{restaurantRepo}
}

func (s restaurantService) GetRestaurants() (restaurantWrapper *RestaurantWrapper, err error) {
	key := "service::GetRestaurants"
	redisClient := utils.GetRedisClient()

	// Redis Get
	if restaurantJson, err := redisClient.Get(context.Background(), key).Result(); err == nil {
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
		redisClient.Set(context.Background(), key, string(data), consts.CACHE_TIME)
	}

	// fmt.Println("Query from database at service - GetRestaurants()", restaurantWrapperDB)
	fmt.Println("view restaurant from database")

	return &restaurantWrapperDB, nil
}

func (s restaurantService) AcceptOrder(acceptOrderReq AcceptOrderRequest) (acceptOrderResp *AcceptOrderResponse, err error) {
	restaurantData, err := s.restaurantRepo.GetRestaurantByID(acceptOrderReq.RestaurantID)
	if err != nil {
		return nil, err
	}

	acceptOrderCoreReq := AcceptOrderCoreRequest{
		OrderID:        acceptOrderReq.OrderID,
		RestaurantID:   acceptOrderReq.RestaurantID,
		RestaurantName: restaurantData.Name,
	}

	fmt.Println("AcceptOrderCoreReq:", acceptOrderCoreReq)

	_, err = utils.GetRestyClient().R().
		SetBody(acceptOrderCoreReq).
		SetResult(&acceptOrderResp).
		Post("/restaurant/order/accept")

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return acceptOrderResp, nil
}
