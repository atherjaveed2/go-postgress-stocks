package models

// User schema of the user table
type Stock struct {
	StockID int64   `json:"stockid"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Company string  `json:"company"`
}
