package alumni

import (
	"time"
)

type Alumni struct {
	ID                 uint        `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	Name               string      `gorm:"type:varchar(100);not null" json:"name"`
	Nationality        string      `gorm:"type:varchar(50)" json:"nationality"`
	Ethnicity          string      `gorm:"type:varchar(50)" json:"ethnicity"`
	DOB                time.Time   `gorm:"type:date" json:"dob"`
	Gender             string      `gorm:"type:varchar(10)" json:"gender"`
	Marital            string      `gorm:"type:varchar(20)" json:"marital"`
	Address            string      `gorm:"type:varchar(255)" json:"address"`
	Email              string      `gorm:"type:varchar(100);" json:"email"`
	MatricNo           string      `gorm:"type:varchar(50);uniqueIndex;not null" json:"matric_no" validate:"required"`
	Password           string      `gorm:"type:varchar(255);not null" json:"password"`
	Phone              string      `gorm:"type:varchar(20)" json:"phone"`
	IsHidden           bool        `gorm:"default:false" json:"is_hidden"`
	HasVerified        bool        `gorm:"default:false" json:"has_verified"`
	Location           string      `gorm:"type:varchar(255)" json:"location"`
	Tags               []AlumniTag `gorm:"foreignKey:AlumniID"`
	LinkedIn           string      `json:"linkedin" gorm:"type:varchar(100);"`
	ParticipantStatus  uint        `json:"participant_status" gorm:"-"`
	ParticipantComment string      `json:"participant_comment" gorm:"-"`
}

type AlumniToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	AlumniID  uint      `gorm:"not null" json:"alumni_id"`
	Token     string    `gorm:"type:varchar(255);not null" json:"token"`
	ExpiresAt time.Time `gorm:"type:bigint;not null" json:"expires_at"`
}

type AlumniTag struct {
	AlumniID uint
	TagID    uint
	Weight   float64 `gorm:"default:1.0"`
}
