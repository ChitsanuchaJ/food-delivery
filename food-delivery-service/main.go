package main

import (
	"fmt"
	"food-delivery-services/handlers"
	"food-delivery-services/producer"
	"food-delivery-services/repositories"
	"food-delivery-services/services"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to get database object: %v\n", err)
	}
	defer sqlDB.Close()

	kafkaProducer := initProducer()
	defer kafkaProducer.Close()

	redisClient := initRedis()
	defer redisClient.Close()

	restaurantRepo := repositories.NewRestaurantRepositoryRedis(db)
	restaurantService := services.NewRestaurantService(restaurantRepo, redisClient)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantService)

	riderRepo := repositories.NewRiderRepositoryRedis(db)
	riderService := services.NewRiderService(riderRepo, redisClient)
	riderHandler := handlers.NewRiderHandler(riderService)

	menuRepo := repositories.NewMenuRepositoryRedis(db)
	menuService := services.NewMenuService(menuRepo, redisClient)
	menuHandler := handlers.NewMenuHandler(menuService)

	eventProducer := producer.NewEventProducer(kafkaProducer)
	orderService := services.NewOrderService(eventProducer, restaurantService, menuService)
	orderHandler := handlers.NewOrderHandler(orderService)

	e := echo.New()

	e.GET("/restaurant", restaurantHandler.GetRestaurants)
	e.GET("/rider", riderHandler.GetRiders)
	e.GET("/menu/:id", menuHandler.GetMenusByID)
	e.POST("/order", orderHandler.PlaceOrder)

	e.Logger.Fatal(e.Start(":8000"))
}

func initProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9093"}, nil)
	if err != nil {
		panic(err)
	}
	return producer
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/food-delivery")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
// 	customers := []Customer{}

// 	h.DB.Find(&customers)

// 	return c.JSON(http.StatusOK, customers)
// }

// func (h *CustomerHandler) GetCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) SaveCustomer(c echo.Context) error {
// 	customer := Customer{}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := h.DB.Delete(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }
