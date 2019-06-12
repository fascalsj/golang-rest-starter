package model

type User struct {
	ID        int    `json:"id" gorm:"AUTO_INCREMENT"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}
