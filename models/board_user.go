package models

type BoardUser struct {
	Model

	User    User      `json:"-"`
	Board   Board     `json:"-"`
	UserID  uint      `json:"user_id" gorm:"type:int references users(id)"`
	BoardID uint      `json:"board_id" gorm:"type:int references boards(id)"`
	Role    BoardRole `json:"role" gorm:"type:varchar(50);not null"`
}

type BoardRole string

const (
	BoardOwner   = BoardRole("owner")
	BoardMember  = BoardRole("member")
	BoardVisitor = BoardRole("visitor")
)
