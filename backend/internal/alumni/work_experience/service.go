package work_experience

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

func (s *Service) CreateWorkExperience(workExperience *WorkExperience) error {
	return s.db.Create(workExperience).Error
}

func (s *Service) GetWorkExperiencesByAlumniID(alumniID uint) ([]WorkExperience, error) {
	var workExperiences []WorkExperience
	err := s.db.Where("alumni_id = ?", alumniID).Order("start_date DESC").Find(&workExperiences).Error
	return workExperiences, err
}

func (s *Service) UpdateWorkExperience(id uint, updatedWorkExperience *WorkExperience) (*WorkExperience, error) {
	var workExperience WorkExperience
	if err := s.db.First(&workExperience, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&workExperience).Updates(updatedWorkExperience).Error; err != nil {
		return nil, err
	}

	return &workExperience, nil
}

func (s *Service) DeleteWorkExperience(id uint) error {
	if err := s.db.Delete(&WorkExperience{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) GetWorksByToken(token string) ([]WorkExperience, error) {
	var works []WorkExperience
	var alu alumni.Alumni
	if err := s.db.Where("token = ?", token).First(&alu).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("alumni_id = ?", alu.ID).Find(&works).Error; err != nil {
		return nil, err
	}
	return works, nil
}
