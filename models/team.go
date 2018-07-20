package models

type Team struct {
	Model

	Name    string  `json:"name" gorm:"type:varchar(255);not null"`
	Slug    string  `json:"slug" gorm:"type:varchar(255);unique_index;not null"`
	Desc    *string `json:"desc"`
	Website *string `json:"desc" gorm:"type:varchar(255)"`
}
