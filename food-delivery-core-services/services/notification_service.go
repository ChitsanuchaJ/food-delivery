package services

import "fmt"

type notificationService struct{}

func NewNotificationService() NotificationService {
	return notificationService{}
}

func (s notificationService) SendNotification(notiReq *NotificationRequest) (notiResp *NotificationResponse, err error) {

	// Log logic
	fmt.Println("Log from send notification ")

	return nil, nil
}
