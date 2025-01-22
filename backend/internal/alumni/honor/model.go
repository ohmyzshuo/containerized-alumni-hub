package honor

import "time"

type Honor struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	AlumniID    uint      `gorm:"not null" json:"alumni_id"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Date        time.Time `gorm:"type:date;not null" json:"date"`
}
