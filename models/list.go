package models

type List struct {
	Model

	Board   Board  `json:"-"`
	Tasks   []Task `json:"tasks"`
	Order   int    `json:"order" gorm:"int;not null"`
	BoardID uint   `json:"board_id" gorm:"type:int references boards(id);not null"`
	Name    string `json:"name" gorm:"type:varchar(255);not null"`
	Color   string `json:"color" gorm:"type:varchar(7) default '#cccccc'"`
}

type SimpleList struct {
	*List
	Tasks omit `json:"tasks,omitempty"`
}

var DefaultListColor = "#cccccc"
