package services

type AcceptOrderRequest struct {
	OrderID        string `json:"order_id"`
	RestaurantID   string `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
}

type AcceptOrderResponse struct {
	Status string `json:"status"`
}

type RestaurantService interface {
	AcceptOrder(AcceptOrderRequest) (*AcceptOrderResponse, error)
}
