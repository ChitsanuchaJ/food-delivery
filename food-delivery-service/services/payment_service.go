package services

import "fmt"

// type PaymentService interface {
// 	Payment(float64) bool
// }

func Payment(orderID string, amount float64) error {
	// return errors.New("")

	fmt.Println("Order Id:", orderID, "has been paid.", "Total amount:", amount, "$")
	return nil
}
