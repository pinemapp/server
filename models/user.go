package models

type User struct {
	Model

	Name     string `json:"name" gorm:"type:varchar(150);unique_index;not null"`
	Email    string `json:"email" gorm:"type:varchar(150);unique_index;not null"`
	Password string `json:"-" gorm:"type:varchar(150);not null"`
}
