package services

import (
	"context"
	"encoding/json"
	"fmt"
	"food-delivery-service/consts"
	"food-delivery-service/repositories"
	"food-delivery-service/utils"
)

type menuService struct {
	menuRepo repositories.MenuRepository
}

func NewMenuService(menuRepo repositories.MenuRepository) MenuService {
	return menuService{menuRepo}
}

func (s menuService) GetMenu(restaurantId string) (menuWrapper *MenuWrapper, err error) {
	key := fmt.Sprintf("service::GetMenu::restaurantId::%v", restaurantId)
	redisClient := utils.GetRedisClient()

	// Redis Get
	if menuJson, err := redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(menuJson), &menuWrapper) == nil {
			// fmt.Println("Cache hit from redis at service - GetMenu()", menuWrapper)
			fmt.Println("view menu from cached")
			return menuWrapper, nil
		}
	}

	menusData, err := s.menuRepo.GetMenu(restaurantId)
	if err != nil {
		return nil, err
	}
	// if len(menusData) == 0 {
	// 	return nil, errors.New("data not found")
	// }

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
		redisClient.Set(context.Background(), key, string(data), consts.CACHE_TIME)
	}

	// fmt.Println("Query from database at service - GetMenu()", menuWrapperDB)
	fmt.Println("view menu from database")

	return &menuWrapperDB, nil
}
