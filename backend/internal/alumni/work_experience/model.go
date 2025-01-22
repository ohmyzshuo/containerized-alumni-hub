package work_experience

import "time"

type WorkExperience struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	AlumniID        uint      `gorm:"not null" json:"alumni_id"`
	Workplace       string    `gorm:"type:varchar(100);not null" json:"workplace"`
	Position        string    `gorm:"type:varchar(100)" json:"position"`
	Country         string    `gorm:"type:varchar(100);not null" json:"country"`
	City            string    `gorm:"type:varchar(100);not null" json:"city"`
	StartDate       time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate         time.Time `gorm:"type:date" json:"end_date"`
	Status          string    `gorm:"type:varchar(50);not null" json:"status"`
	OccupationField string    `gorm:"type:varchar(100);not null" json:"occupation_field"`
}
