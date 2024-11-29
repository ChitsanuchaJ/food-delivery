package services

import "fmt"

type notificationService struct{}

func NewNotificationService() NotificationService {
	return notificationService{}
}

func (s notificationService) SendNotification(notiReq NotificationRequest) (notiResp *NotificationResponse, err error) {

	notiResp = &NotificationResponse{
		Status: "sent",
	}
	// Log logic
	fmt.Println("Log from notification ")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("To:", notiReq.Recipient, ":", notiReq.Message)
	fmt.Println("--------------------------------------------------------")

	return notiResp, nil
}
