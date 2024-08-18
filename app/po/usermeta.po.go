package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMeta struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid" gorm:"type:char(36);index; not null"`
	MetaKey   string    `json:"meta_key" gorm:"type:varchar(255)"`
	MetaValue string    `json:"meta_value" gorm:"type:text"`
	MetaJson  string    `json:"meta_json" gorm:"type:json"`
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"foreignKey:UsersRefer"`
}

func (r *UserMeta) TableName() string {
	return "usermeta"
}
