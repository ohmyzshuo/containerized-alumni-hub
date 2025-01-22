package content

import (
	"time"
)

type Content struct {
	ID               uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	Title            string       `json:"title" gorm:"size:255"`
	Description      string       `json:"description" gorm:"type:text"`
	Contact          string       `json:"contact" gorm:"type:text"`
	CreatedBy        uint         `json:"created_by" gorm:"foreignKey:StaffID"`
	UpdatedBy        uint         `json:"updated_by" gorm:"foreignKey:StaffID"`
	IsHidden         bool         `json:"is_hidden"`
	IsPinned         bool         `json:"is_pinned"`
	CreatedAt        time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	Venue            string       `json:"venue" gorm:"type:varchar(100)"`
	FacultyID        uint         `json:"faculty_id" gorm:"foreignKey:FacultyID"`
	StartTime        time.Time    `json:"start_time"`
	EndTime          time.Time    `json:"end_time"`
	ContentType      uint         `json:"content_type"` // {1: "Event", 2: "Announcement", 3: "Advertisement", 0: "Other"}
	NumOfLikes       uint         `json:"num_of_likes"`
	Tags             []Tag        `json:"-" gorm:"many2many:content_tags;"`
	ParticipantQuota uint         `json:"participant_quota"` // 0 means unlimited quota
	Attachments      []Attachment `json:"attachments" gorm:"-"`
	Status           uint         `json:"status" gorm:"-"`
	CreatedByName    string       `json:"created_by_name" gorm:"-"`
}

type Attachment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	ContentId      uint      `json:"content_id" gorm:"foreignKey:ContentID"`
	OriginalName   string    `json:"original_name" gorm:"type:varchar(255)"`
	AttachmentPath string    `json:"attachment_path" gorm:"type:varchar(500)"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;unique"`
}

type ContentTag struct {
	ContentID uint
	TagID     uint
}

type EventParticipant struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	AlumniID  uint   `json:"alumni_id" gorm:"not null"`
	ContentID uint   `json:"content_id" gorm:"not null"`
	Status    uint   `json:"status" gorm:"not null"`
	Comment   string `json:"comment"`
}
