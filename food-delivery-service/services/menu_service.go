package services

import (
	"context"
	"encoding/json"
	"fmt"
	"food-delivery-services/consts"
	"food-delivery-services/repositories"

	"github.com/go-redis/redis/v8"
)

type menuService struct {
	menuRepo    repositories.MenuRepository
	redisClient *redis.Client
}

func NewMenuService(menuRepo repositories.MenuRepository, redisClient *redis.Client) MenuService {
	return menuService{menuRepo, redisClient}
}

func (s menuService) GetMenus(restaurantId string) (menuWrapper *MenuWrapper, err error) {
	key := fmt.Sprintf("service::GetMenus::restaurantId::%v", restaurantId)

	// Redis Get
	if menuJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(menuJson), &menuWrapper) == nil {
			// fmt.Println("Cache hit from redis at service - GetMenus()", menuWrapper)
			fmt.Println("view menu from cached")
			return menuWrapper, nil
		}
	}

	menusData, err := s.menuRepo.GetMenus(restaurantId)
	if err != nil {
		return nil, err
	}

	menus := []Menu{}
	for _, menu := range menusData {
		menus = append(menus, Menu{
			ID:          menu.ID,
			Name:        menu.Name,
			Price:       menu.Price,
			Description: menu.Description,
		})
	}
	menuWrapperDB := MenuWrapper{
		RestaurantId: restaurantId,
		Menus:        menus,
	}

	// Redis SET
	if data, err := json.Marshal(menuWrapperDB); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), consts.CACHE_TIME)
	}

	// fmt.Println("Query from database at service - GetMenus()", menuWrapperDB)
	fmt.Println("view menu from database")

	return &menuWrapperDB, nil
}
