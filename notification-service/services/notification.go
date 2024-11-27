package services

type NotificationRequest struct {
	Recipient string `json:"recipient"`
	OrderID   string `json:"order_id"`
	Message   string `json:"message"`
}

type NotificationResponse struct {
	Status string `json:"status"`
}

type NotificationService interface {
	SendNotification(*NotificationRequest) (*NotificationResponse, error)
}
