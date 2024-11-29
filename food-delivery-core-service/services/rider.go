package services

type PickUpOrderRequest struct {
	OrderID string `json:"order_id"`
	RiderID string `json:"rider_id"`
}

type PickUpOrderResponse struct {
	Status string `json:"status"`
}

type RiderService interface {
	PickUpOrder(PickUpOrderRequest) (*PickUpOrderResponse, error)
}
