package models

type Board struct {
	Model

	Members []BoardUser `json:"members"`
	Name    string      `json:"name" gorm:"type:varchar(255);not null"`
	Desc    string      `json:"desc" gorm:"type:text"`
	Public  bool        `json:"public" gorm:"type:boolean default false"`
}
