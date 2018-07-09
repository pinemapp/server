package models

type Board struct {
	Model

	User   User   `json:"-"`
	Name   string `json:"name" gorm:"type:varchar(255);not null"`
	Desc   string `json:"desc" gorm:"type:text"`
	Public bool   `json:"public" gorm:"type:boolean default false"`
	UserID uint   `json:"user_id" gorm:"type:int references users(id)"`
}
