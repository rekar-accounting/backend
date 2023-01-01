package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name       string
	Family     string
	Username   string
	Email      string
	Password   string
	Mobile     string
	VerifiedAt time.Time
}
