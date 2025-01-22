package staff

import (
	"alumni_hub/internal/utils"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetStaffs(page, pageSize int, searchQuery string, filters map[string]string) ([]Staff, int64, error) {
	var staff []Staff

	offset := (page - 1) * pageSize
	query := s.db.Debug().Offset(offset).Limit(pageSize).Where("is_hidden = ?", false)

	if searchQuery != "" {
		query = query.Where("username ILIKE ? OR name ILIKE ? OR email ILIKE ? OR phone ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	if isSuperAdminStr, ok := filters["is_super_admin"]; ok && isSuperAdminStr != "" {
		isSuperAdmin, err := strconv.ParseBool(isSuperAdminStr)
		if err == nil {
			query = query.Where("is_super_admin = ?", isSuperAdmin)
		}
	}
	if facultyIDStr, ok := filters["faculty_id"]; ok && facultyIDStr != "" {
		facultyID, err := strconv.Atoi(facultyIDStr)
		if err == nil {
			query = query.Where("faculty_id = ?", facultyID)
		}
	}

	if err := query.Find(&staff).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	countQuery := s.db.Model(&Staff{}).Where("is_hidden = ?", false).Where("username ILIKE ? OR name ILIKE ? OR email ILIKE ? OR phone ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%")
	if isSuperAdminStr, ok := filters["is_super_admin"]; ok && isSuperAdminStr != "" {
		isSuperAdmin, err := strconv.ParseBool(isSuperAdminStr)
		if err == nil {
			countQuery = countQuery.Where("is_super_admin = ?", isSuperAdmin)
		}
	}
	if facultyIDStr, ok := filters["faculty_id"]; ok && facultyIDStr != "" {
		facultyID, err := strconv.Atoi(facultyIDStr)
		if err == nil {
			countQuery = countQuery.Where("faculty_id = ?", facultyID)
		}
	}
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return staff, total, nil
}

func (s *Service) CreateStaff(staff *Staff) (*Staff, error) {
	var existingStaff Staff
	if err := s.db.Where("username = ?", staff.Username).First(&existingStaff).Error; err == nil {
		return nil, errors.New("duplicated username")
	}

	if staff.Password == "" {
		staff.Password = staff.Username // Set default password if empty
	}

	hashedPassword, err := utils.HashPassword(staff.Password)
	if err != nil {
		return nil, err
	}
	staff.Password = hashedPassword

	if err := s.db.Create(staff).Error; err != nil {
		return nil, err
	}

	return staff, nil
}
func (s *Service) UpdateStaff(id uint, updatedStaff *Staff) (*Staff, error) {
	var staff Staff
	if err := s.db.First(&staff, id).Error; err != nil {
		return nil, err
	}

	if updatedStaff.Username != "" && updatedStaff.Username != staff.Username {
		var existingStaff Staff
		if err := s.db.Where("username = ?", updatedStaff.Username).First(&existingStaff).Error; err == nil {
			return nil, errors.New("duplicated username")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if err := s.db.Model(&staff).Updates(updatedStaff).Error; err != nil {
		return nil, err
	}

	return &staff, nil
}

func (s *Service) DeleteStaff(id uint) error {
	return s.db.Model(&Staff{}).Where("id = ?", id).Update("is_hidden", true).Error
}

func (s *Service) GetStaffByToken(token string) (*Staff, error) {
	var sta Staff
	if err := s.db.Where("token = ?", token).First(&sta).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sta, nil
}

func (s *Service) GetStaffByID(id uint) (*Staff, error) {
	var staff Staff
	if err := s.db.First(&staff, id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}