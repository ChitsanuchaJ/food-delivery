package services

type RiderWrapper struct {
	Riders []Rider `json:"rider"`
}

type Rider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RiderService interface {
	GetRiders() (*RiderWrapper, error)
}
