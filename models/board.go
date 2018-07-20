package models

type Board struct {
	Model

	Team    *Team       `json:"-"`
	Lists   []List      `json:"lists"`
	Members []BoardUser `json:"members"`
	Name    string      `json:"name" gorm:"type:varchar(255);not null"`
	Desc    *string     `json:"desc" gorm:"type:text"`
	Public  bool        `json:"public" gorm:"type:boolean default false"`
	TeamID  *uint       `json:"team_id" gorm:"type:int references teams(id)"`
}

type SimpleBoard struct {
	*Board
	Team    omit `json:"team,omitempty"`
	Lists   omit `json:"lists,omitempty"`
	Members omit `json:"members,omitempty"`
}
