package staff

import (
	"time"
)

type Staff struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name         string    `json:"name" validate:"required"`
	Email        string    `json:"email" validate:"email"`
	IsSuperAdmin bool      `json:"is_super_admin" gorm:"default:false"`
	FacultyID    uint      `json:"faculty_id" gorm:"default:1"`
	Phone        string    `json:"phone"`
	Username     string    `json:"username" validate:"required"`
	Password     string    `json:"password" validate:"required"`
	IsHidden     bool      `json:"-" gorm:"default:false"`
	Gender       string    `json:"gender"`
	Position     string    `json:"position"`
}

//func (s *Staff) AfterFind(db *gorm.db) (err error) {
//	var f faculty.Faculty
//	if err := db.First(&f, s.FacultyID).Error; err != nil {
//		return err
//	}
//	s.FacultyName = f.Name
//	return nil
//}

type StaffToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	StaffID   uint      `gorm:"not null" json:"staff_id"`
	Token     string    `gorm:"type:varchar(255);not null" json:"token"`
	ExpiresAt int64     `gorm:"type:bigint;not null" json:"expires_at"`
}
