package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:char(36);not null;index:idx_uuid" json:"uuid"`
	UserName string    `gorm:"column:username;type:varchar(50);not null;index:idx_username" json:"username"`
	IsActive bool      `gorm:"column:is_active;type:tinyint(1);not null;default:1" json:"is_active"`
	Roles    []Role    `gorm:"many2many:go_user_roles;" json:"users"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
