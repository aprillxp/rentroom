package models

type User struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Username   string     `gorm:"unique" json:"username"`
	Email      string     `json:"email"`
	Phone      string     `gorm:"uniqueIndex;size:20"`
	Password   string     `json:"-"`
	BankID     uint       `json:"bank_id"`
	BankNumber string     `json:"bank_number"`
	IsTenant   bool       `json:"is_tenant"`
	Property   []Property `gorm:"many2many:user_properties;" json:"properties"`
}
type UserProperties struct {
	UserID     uint `gorm:"primaryKey" json:"user_id"`
	PropertyID uint `gorm:"primaryKey" json:"property_id"`
}
type UserRegisterRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required"`
	Password   string `json:"password" validate:"required"`
	BankID     uint   `json:"bank_id" validate:"required,gt=0"`
	BankNumber string `json:"bank_number" validate:"required"`
	IsTenant   bool   `json:"is_tenant"`
}
type UserLoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
type UserLoginResponse struct {
	Token string `json:"token"`
}
