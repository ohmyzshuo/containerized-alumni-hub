package honor

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

func (s *Service) CreateHonor(h *Honor) error {
	return s.db.Create(h).Error
}

func (s *Service) GetHonorByAlumniID(alumniID uint) ([]Honor, error) {
	var honors []Honor
	if err := s.db.Debug().Where("alumni_id = ?", alumniID).Find(&honors).Error; err != nil {
		return nil, err
	}
	return honors, nil
}

func (s *Service) UpdateHonor(id uint, updatedHonor *Honor) (*Honor, error) {
	var h Honor
	if err := s.db.First(&h, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&h).Updates(updatedHonor).Error; err != nil {
		return nil, err
	}

	return &h, nil
}

func (s *Service) DeleteHonor(id uint) error {
	return s.db.Delete(&Honor{}, id).Error
}

func (s *Service) GetHonorsByToken(token string) ([]Honor, error) {
	var honors []Honor
	var alu alumni.Alumni
	if err := s.db.Where("token = ?", token).First(&alu).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("alumni_id = ?", alu.ID).Find(&honors).Error; err != nil {
		return nil, err
	}
	return honors, nil
}