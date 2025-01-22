package study

import (
	"alumni_hub/internal/alumni"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateStudy(study *Study) error {
	return s.db.Create(study).Error
}

func (s *Service) GetStudiesByAlumniID(alumniID uint) ([]Study, error) {
	var studies []Study
	err := s.db.Where("alumni_id = ?", alumniID).Order("created_at DESC").Find(&studies).Error
	return studies, err
}

func (s *Service) UpdateStudy(id uint, updatedStudy *Study) (*Study, error) {
	var study Study
	if err := s.db.First(&study, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&study).Updates(updatedStudy).Error; err != nil {
		return nil, err
	}

	return &study, nil
}

func (s *Service) DeleteStudy(id uint) error {
	if err := s.db.Delete(&Study{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) GetStudiesByToken(token string) ([]Study, error) {
	var studies []Study
	var alu alumni.Alumni
	if err := s.db.Where("token = ?", token).First(&alu).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("alumni_id = ?", alu.ID).Find(&studies).Error; err != nil {
		return nil, err
	}
	return studies, nil
}
