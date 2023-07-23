package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Video struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Title          string    `gorm:"type:text;not null" json:"title"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	FileName       string    `gorm:"type:text;not null" json:"file_name"`
	FileSize       string    `gorm:"type:text;not null" json:"file_size"`
	LastModified   string    `gorm:"type:text;not null" json:"last_modified"`
	ChecksumSHA256 string    `gorm:"type:text;not null" json:"checksum_sha256"`
	// CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	// UpdatedAt      time.Time `gorm:"not null" json:"updated_at"`
}

// See: https://gorm.io/docs/create.html#Create-Hooks
func (video *Video) BeforeCreate(*gorm.DB) error {
	video.ID = uuid.NewV4()
	return nil
}
