package services

type PickUpOrderRequest struct {
	OrderID string `json:"order_id"`
	RiderID string `json:"rider_id"`
}

type PickUpOrderResponse struct {
	Status string `json:"status"`
}

//////////////////////////////////////////////////////

type DeliverOrderRequest struct {
	OrderID string `json:"order_id"`
	RiderID string `json:"rider_id"`
}

type DeliverOrderResponse struct {
	Status string `json:"status"`
}

type RiderService interface {
	PickUpOrder(PickUpOrderRequest) (*PickUpOrderResponse, error)
	DeliverOrder(DeliverOrderRequest) (*DeliverOrderResponse, error)
}
