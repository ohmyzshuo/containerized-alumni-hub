package study

import "time"

type Study struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	AlumniID        uint      `gorm:"not null" json:"alumni_id"`
	LevelOfStudy    string    `gorm:"type:varchar(50);not null" json:"level_of_study"`
	FacultyID       int       `gorm:"" json:"faculty_id"`
	Programme       string    `gorm:"type:varchar(100)" json:"programme"`
	IntakeYear      int       `gorm:"type:int" json:"intake_year"`
	IntakeSession   string    `gorm:"type:varchar(10)" json:"intake_session"`
	ConvocationYear int       `gorm:"type:int" json:"convocation_year"`
	Status          string    `gorm:"type:varchar(50)" json:"status"`
	TitleOfThesis   string    `gorm:"type:varchar(255)" json:"title_of_thesis"`
	Supervisor      string    `gorm:"type:varchar(100)" json:"supervisor"`
}
