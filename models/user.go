package models

type User struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Username   string     `gorm:"unique" json:"username"`
	Password   string     `json:"password"`
	BankNumber int        `json:"bank_number"`
	BankID     uint       `json:"bank_id"`
	IsTenant   int        `json:"is_tenant"`
	Phone      string     `gorm:"uniqueIndex;size:20"`
	Email      string     `json:"email"`
	Property   []Property `gorm:"many2many:user_properties;" json:"amenities"`
}

type UserLoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
