package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:bigint(20);not null;primary_key;auto_increment" json:"id"`
	RoleName string `gorm:"column:role_name;type:varchar(50);not null;index:idx_role_name" json:"role_name"`
	RoleNote string `gorm:"column:role_note;type:varchar(255);not null;index:idx_role_note" json:"role_note"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
