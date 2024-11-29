package main

import (
	"fmt"
	"food-delivery-service/handlers"
	"food-delivery-service/producer"
	"food-delivery-service/repositories"
	"food-delivery-service/services"
	"food-delivery-service/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	utils.InitRestyClient()
	utils.InitRedis()

	db := utils.InitDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to get database object: %v\n", err)
	}
	defer sqlDB.Close()

	kafkaProducer := utils.InitProducer()
	defer kafkaProducer.Close()

	redisClient := utils.GetRedisClient()
	defer redisClient.Close()

	restaurantRepo := repositories.NewRestaurantRepositoryRedis(db)
	restaurantService := services.NewRestaurantService(restaurantRepo)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantService)

	riderRepo := repositories.NewRiderRepositoryRedis(db)
	riderService := services.NewRiderService(riderRepo)
	riderHandler := handlers.NewRiderHandler(riderService)

	menuRepo := repositories.NewMenuRepositoryRedis(db)
	menuService := services.NewMenuService(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuService)

	eventProducer := producer.NewEventProducer(kafkaProducer)
	orderService := services.NewOrderService(eventProducer, restaurantRepo, menuRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	app := echo.New()
	app.Use(middleware.Recover())
	// app.Use(middleware.Logger())

	app.GET("/restaurant", restaurantHandler.GetRestaurants)
	app.POST("/restaurant/order/accept", restaurantHandler.AcceptOrder)

	app.GET("/rider", riderHandler.GetRiders)
	app.POST("/rider/order/pickup", riderHandler.PickUpOrder)
	app.POST("/rider/order/deliver", riderHandler.DeliverOrder)

	app.GET("/menu/:id", menuHandler.GetMenusByID)
	app.POST("/order", orderHandler.PlaceOrder)

	app.Logger.Fatal(app.Start(":8000"))
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
