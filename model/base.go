package model

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	ID        uint         `json:"id" gorm:"primaryKey;autoIncrement:true;type:bigint unsigned;notNull"`
	CreatedAt time.Time    `json:"createdAt" gorm:"type:datetime(6);notNull;default:CURRENT_TIMESTAMP(6)"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"type:datetime(6);notNull;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)"`
	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index;type:datetime(6);default:null;softDelete:soft"`
}
