package email

import (
	"alumni_hub/internal/utils"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}
func (s *Service) CreateEmail(email *Email) error {
	var attachmentPaths []string
	for _, attachment := range email.Attachments {
		attachmentPaths = append(attachmentPaths, attachment.AttachmentPath)
	}

	err := utils.SendEmail(email.To, email.Subject, email.Body, attachmentPaths)
	if err != nil {
		return err
	}

	if err := s.db.Create(email).Error; err != nil {
		return err
	}

	return nil
}
func (s *Service) GetEmails(page, pageSize int) ([]Email, int64, error) {
	var emails []Email
	var total int64

	offset := (page - 1) * pageSize
	if err := s.db.Model(&Email{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := s.db.Preload("Attachments").Limit(pageSize).Offset(offset).Find(&emails).Error; err != nil {
		return nil, 0, err
	}

	return emails, total, nil
}
