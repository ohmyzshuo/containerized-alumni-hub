package email

import (
	"time"
)

type Email struct {
	ID          uint         `gorm:"primarykey"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	To          string       `json:"to"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	Attachments []Attachment `json:"attachments" gorm:"foreignKey:EmailID"`
}

type Attachment struct {
	ID             uint      `gorm:"primarykey"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	EmailID        uint      `json:"email_id"`
	OriginalName   string    `json:"original_name"`
	AttachmentPath string    `json:"attachment_path"`
}

func (Attachment) TableName() string {
	return "email_attachments"
}
