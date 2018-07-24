package models

type ProjectUser struct {
	Model

	User      User        `json:"-"`
	Project   Project     `json:"-"`
	UserID    uint        `json:"user_id" gorm:"type:int references users(id);not null"`
	ProjectID uint        `json:"project_id" gorm:"type:int references projects(id);not null"`
	Role      ProjectRole `json:"role" gorm:"type:varchar(50);not null"`
}

type ProjectRole string

const (
	ProjectOwner   = ProjectRole("owner")
	ProjectMember  = ProjectRole("member")
	ProjectVisitor = ProjectRole("visitor")
)
