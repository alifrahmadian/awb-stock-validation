package model

type (
	Order struct {
		// TODO: add other fields if needed (ex: created_at, etc)
		ID          int64   `json:"id"`
		AWBNumber   string  `json:"awb_number"`
		Sender      string  `json:"sender"`
		Receiver    string  `json:"receiver"`
		TotalWeight float64 `json:"total_weight"`
		TotalPrice  int     `json:"total_price"`
		Status      string  `json:"status"`
	}
)

var (
	OrderList = []Order{
		// TODO: add more data here
	}
)
