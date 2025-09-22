package dto

type CreateOrderRequest struct {
	AWBNumber   string  `json:"awb_number"`
	Sender      string  `json:"sender"`
	Receiver    string  `json:"receiver"`
	TotalWeight float64 `json:"total_weight"`
}

type CreateOrderResponse struct {
	ID          int64   `json:"id"`
	AWBNumber   string  `json:"awb_number"`
	Sender      string  `json:"sender"`
	Receiver    string  `json:"receiver"`
	TotalWeight float64 `json:"total_weight"`
	TotalPrice  int     `json:"total_price"`
	Status      string  `json:"status"`
}
