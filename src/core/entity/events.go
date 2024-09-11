package entity

type OrderEvent struct {
	Id          string `json:"id"`
	EventType   string `json:"eventType"`
	OrderStatus string `json:"orderStatus"`
	Order       *Order `json:"order"`
}

type PaymentEvent struct {
	Id        string `json:"id"`
	EventType string `json:"eventType"`
	Order     struct {
		Id     string  `json:"id"`
		Amount float64 `json:"amount"`
	}
}
