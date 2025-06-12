package models

type User struct {
	BaseModel
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
}
