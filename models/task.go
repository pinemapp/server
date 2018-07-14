package models

import "time"

type Task struct {
	Model

	List    List       `json:"-"`
	Board   Board      `json:"-"`
	ListID  uint       `json:"list_id" gorm:"type:int references lists(id);not null"`
	BoardID uint       `json:"board_id" gorm:"type:int references boards(id);not null"`
	Name    string     `json:"name" gorm:"type:varchar(255);not null"`
	Desc    string     `json:"desc" gorm:"type:text"`
	Order   int        `json:"order" gorm:"type:int;not null"`
	StartAt *time.Time `json:"start_at" gorm:"type:timestamptz"`
	EndAt   *time.Time `json:"end_at" gorm:"type:timestamptz"`
}
