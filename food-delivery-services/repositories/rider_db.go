package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type riderRepositoryRedis struct {
	db *gorm.DB
}

func NewRiderRepositoryRedis(db *gorm.DB) RiderRepository {
	db.AutoMigrate(&rider{})
	mockRiderData(db)
	return riderRepositoryRedis{db}
}

func (r riderRepositoryRedis) GetRiders() (riders []rider, err error) {
	err = r.db.Find(&riders).Error
	if err != nil {
		return nil, err
	}

	fmt.Println("Query from database at repository - GetRiders()")
	return riders, nil
}
