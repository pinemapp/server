package models

import "time"

type Task struct {
	Model

	List      List       `json:"-"`
	Project   Project    `json:"-"`
	ListID    uint       `json:"list_id" gorm:"type:int references lists(id);not null"`
	ProjectID uint       `json:"project_id" gorm:"type:int references projects(id);not null"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Desc      string     `json:"desc" gorm:"type:text"`
	Order     int        `json:"order" gorm:"type:int;not null"`
	StartAt   *time.Time `json:"start_at" gorm:"type:timestamptz"`
	EndAt     *time.Time `json:"end_at" gorm:"type:timestamptz"`
}
