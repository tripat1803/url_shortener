package models

import "time"

type Url struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex:user_cypher"`
	PlainText uint64
	Cypher    string `gorm:"uniqueIndex:user_cypher"`
	LongUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
