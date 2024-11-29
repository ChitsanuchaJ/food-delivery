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

func (r restaurantRepositoryRedis) GetRestaurantByID(id string) (restaurant *restaurant, err error) {
	err = r.db.First(&restaurant, id).Error
	if err != nil {
		fmt.Println("Data not found, id:", id, "is not existing.")
		return nil, err
	}

	fmt.Println("Query from database at repository - GetRestaurantByID() id:", id)
	return restaurant, nil
}
