package model

type (
	/*
		initial data:
		001-10000001 | not_in_use
		001-10008272 | not_in_use
		001-10726855 | not_in_use
		001-36748596 | not_in_use
		001-63748392 | not_in_use
		001-12345675 | not_in_use
		001-12345686 | not_in_use
	*/
	AWBStock struct {
		// TODO: add other fields if needed (ex: created_at, etc)
		AWBNumber string `json:"awb_number"`
		Status    string `json:"status"`
	}
)

var (
	AWBStockList = []AWBStock{
		{
			AWBNumber: "001-10000001",
			Status:    "not_in_use",
		},
		// TODO: add more data here
	}
)
