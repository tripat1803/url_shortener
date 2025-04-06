package models

import "time"

type Url struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:false"`
	UserID    uint   `gorm:"uniqueIndex:user_cypher"`
	Cipher    string `gorm:"uniqueIndex:user_cypher"`
	LongUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
