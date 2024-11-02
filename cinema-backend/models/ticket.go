package models

type Ticket struct {
	ID            int     `json:"id"`
	ScheduleID    int     `json:"schedule_id"`
	UserID        int     `json:"user_id"`
	Quantity      int     `json:"quantity"`
	TotalPrice    float64 `json:"total_price"`
	PaymentStatus string  `json:"payment_status"`
}
