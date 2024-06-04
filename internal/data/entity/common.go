package entity

import "time"

// CommonInfo 通用类，每个实体类都应该加上
type CommonInfo struct {
	CreatedAt time.Time `gorm:"type:datetime(3);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(3);not null"`
	DeletedAt time.Time `gorm:"type:datetime(3);not null"`
}
