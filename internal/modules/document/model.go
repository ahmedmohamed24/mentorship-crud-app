package document

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
