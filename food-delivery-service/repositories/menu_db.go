package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type menuRepositoryRedis struct {
	db *gorm.DB
}

func NewMenuRepositoryRedis(db *gorm.DB) MenuRepository {
	db.AutoMigrate(&Menu{})
	mockMenuData(db)
	return menuRepositoryRedis{db}
}

func (r menuRepositoryRedis) GetMenu(restaurantId string) (menus []Menu, err error) {
	err = r.db.Where("restaurant_id = ?", restaurantId).Find(&menus).Error
	if err != nil {
		return nil, err
	}

	fmt.Println("Query from database at repository - GetMenu()")
	return menus, nil
}
