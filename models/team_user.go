package models

type TeamUser struct {
	Model

	Role   TeamRole `json:"-" gorm:"type:varchar(50);not null"`
	TeamID uint     `json:"-" gorm:"type:int references teams(id);not null"`
	UserID uint     `json:"-" gorm:"type:int references users(id);not null"`
}

type TeamRole string

const (
	TeamLeader = TeamRole("leader")
	TeamMember = TeamRole("member")
)
