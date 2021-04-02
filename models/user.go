package models

// User struct contains details of attributes of user details
type User struct {
	ID      int     `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int     `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}
