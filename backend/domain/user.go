package domain

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	Username string    `gorm:"size:100;not null" json:"username"`
	Password string    `gorm:"size:100;not null" json:"password"`
	Name     string    `gorm:"size:100;not null" json:"name"`
}
