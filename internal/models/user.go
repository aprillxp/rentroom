package models

type User struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Username   string     `gorm:"unique" json:"username"`
	Email      string     `json:"email"`
	Phone      string     `gorm:"uniqueIndex;size:20"`
	Password   string     `json:"-"`
	Bank       string     `json:"bank"`
	BankNumber string     `json:"bank_number"`
	IsTenant   bool       `json:"is_tenant"`
	Property   []Property `gorm:"many2many:user_properties;" json:"properties"`
}

type UserProperties struct {
	UserID     uint `gorm:"primaryKey;not null"`
	PropertyID uint `gorm:"primaryKey;not null"`

	User     User     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Property Property `gorm:"foreignKey:PropertyID;references:ID;constraint:OnDelete:CASCADE"`
}

type UserRegisterRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Bank       string `json:"bank" validate:"required,min=3,max=50,alphanum"`
	BankNumber string `json:"bank_number" validate:"required"`
	IsTenant   bool   `json:"is_tenant"`
}
type UserEditRequest struct {
	Username   *string `json:"username" validate:"omitempty,min=3,max=50,alphanum"`
	Email      *string `json:"email" validate:"omitempty,email"`
	Phone      *string `json:"phone" validate:"omitempty"`
	Password   *string `json:"password" validate:"omitempty"`
	Bank       *string `json:"bank" validate:"omitempty,min=3,max=50,alphanum"`
	BankNumber *string `json:"bank_number" validate:"omitempty"`
}
type UserResponse struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Bank       string `json:"bank"`
	BankNumber string `json:"bank_number"`
	IsTenant   bool   `json:"is_tenant"`
}
type UserLoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
type UserLoginResponse struct {
	Token string `json:"token"`
}
type UserRegisterResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
