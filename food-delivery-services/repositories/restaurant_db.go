package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type restaurantRepositoryRedis struct {
	db *gorm.DB
}

func NewRestaurantRepositoryRedis(db *gorm.DB) RestaurantRepository {
	db.AutoMigrate(&restaurant{})
	mockRestaurantData(db)
	return restaurantRepositoryRedis{db}
}

func (r restaurantRepositoryRedis) GetRestaurants() (restaurants []restaurant, err error) {
	err = r.db.Find(&restaurants).Error
	if err != nil {
		return nil, err
	}

	fmt.Println("Query from database at repository - GetRestaurants()")
	return restaurants, nil
}
