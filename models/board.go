package models

type Board struct {
	Model

	Lists   []List      `json:"lists"`
	Members []BoardUser `json:"members"`
	Name    string      `json:"name" gorm:"type:varchar(255);not null"`
	Desc    *string     `json:"desc" gorm:"type:text"`
	Public  bool        `json:"public" gorm:"type:boolean default false"`
}

type SimpleBoard struct {
	*Board
	Lists   omit `json:"lists,omitempty"`
	Members omit `json:"members,omitempty"`
}
