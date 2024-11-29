package services

type RestaurantWrapper struct {
	Restaurants []Restaurant `json:"restaurant"`
}

type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//////////////////////////////////////////////////////

type AcceptOrderRequest struct {
	OrderID      string `json:"order_id"`
	RestaurantID string `json:"restaurant_id"`
}

type AcceptOrderResponse struct {
	Status string `json:"status"`
}

// Forward request to core service with additional field
type AcceptOrderCoreRequest struct {
	OrderID        string `json:"order_id"`
	RestaurantID   string `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
}

type RestaurantService interface {
	GetRestaurants() (*RestaurantWrapper, error)
	AcceptOrder(AcceptOrderRequest) (*AcceptOrderResponse, error)
}
