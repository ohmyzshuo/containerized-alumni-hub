package alumni

import (
	"alumni_hub/internal/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"reflect"
	"time"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetAlumnus(id int) (*Alumni, error) {
	var alumnus Alumni
	result := s.db.Where("id = ? AND is_hidden = false", id).First(&alumnus)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &alumnus, nil
}

func (s *Service) GetAlumni(page, pageSize int, searchQuery string) ([]Alumni, int64, error) {
	var alumni []Alumni

	offset := (page - 1) * pageSize
	query := s.db.Debug().Offset(offset).Limit(pageSize).Where("is_hidden = ?", false) // 使用 Debug() 输出 SQL 语句

	if searchQuery != "" {
		var publicationAlumniIDs, studyAlumniIDs, workExperienceAlumniIDs []uint

		s.db.Table("publications").Select("alumni_id").Where("journal_title ILIKE ? OR article_title ILIKE ? OR authors ILIKE ? OR corresponding_authors ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").Pluck("alumni_id", &publicationAlumniIDs)
		s.db.Table("studies").Select("alumni_id").Where("title_of_thesis ILIKE ? OR supervisor ILIKE ? OR programme ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").Pluck("alumni_id", &studyAlumniIDs)
		s.db.Table("work_experiences").Select("alumni_id").Where("workplace ILIKE ? OR position ILIKE ? OR country ILIKE ? OR city ILIKE ? OR occupation_field ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").Pluck("alumni_id", &workExperienceAlumniIDs)

		alumniIDMap := make(map[uint]struct{})
		for _, id := range publicationAlumniIDs {
			alumniIDMap[id] = struct{}{}
		}
		for _, id := range studyAlumniIDs {
			alumniIDMap[id] = struct{}{}
		}
		for _, id := range workExperienceAlumniIDs {
			alumniIDMap[id] = struct{}{}
		}

		var uniqueAlumniIDs []uint
		for id := range alumniIDMap {
			uniqueAlumniIDs = append(uniqueAlumniIDs, id)
		}

		query = query.Where(
			"id IN (?) OR name ILIKE ? OR email ILIKE ? OR phone ILIKE ? OR matric_no ILIKE ? OR location ILIKE ? OR address ILIKE ? OR nationality ILIKE ? OR linked_in ILIKE ?",
			uniqueAlumniIDs,
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
		)
	}

	if err := query.Find(&alumni).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	countQuery := s.db.Model(&Alumni{}).Debug().Where("is_hidden = ?", false)
	if searchQuery != "" {
		var publicationAlumniIDs, studyAlumniIDs, workExperienceAlumniIDs []uint

		s.db.Table("publications").Select("alumni_id").Where("journal_title ILIKE ? OR article_title ILIKE ? OR authors ILIKE ? OR corresponding_authors ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").Pluck("alumni_id", &publicationAlumniIDs)
		s.db.Table("studies").Select("alumni_id").Where("title_of_thesis ILIKE ? OR supervisor ILIKE ? OR programme ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").Pluck("alumni_id", &studyAlumniIDs)
		s.db.Table("work_experiences").Select("alumni_id").Where("workplace ILIKE ? OR position ILIKE ? OR country ILIKE ? OR city ILIKE ? OR occupation_field ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").Pluck("alumni_id", &workExperienceAlumniIDs)

		alumniIDMap := make(map[uint]struct{})
		for _, id := range publicationAlumniIDs {
			alumniIDMap[id] = struct{}{}
		}
		for _, id := range studyAlumniIDs {
			alumniIDMap[id] = struct{}{}
		}
		for _, id := range workExperienceAlumniIDs {
			alumniIDMap[id] = struct{}{}
		}

		var uniqueAlumniIDs []uint
		for id := range alumniIDMap {
			uniqueAlumniIDs = append(uniqueAlumniIDs, id)
		}

		countQuery = countQuery.Where(
			"id IN (?) OR name ILIKE ? OR email ILIKE ? OR phone ILIKE ? OR matric_no ILIKE ? OR location ILIKE ? OR address ILIKE ? OR nationality ILIKE ? OR linked_in ILIKE ?",
			uniqueAlumniIDs,
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
		)

	}

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return alumni, total, nil
}

func (s *Service) CreateAlumni(alumni *Alumni) (*Alumni, error) {
	var existingAlumni Alumni
	if err := s.db.Where("matric_no = ?", alumni.MatricNo).First(&existingAlumni).Error; err == nil {
		return nil, errors.New("duplicated matric_no")
	}

	// If password is empty, set it to the extracted login name
	if alumni.Password == "" {
		loginName, err := utils.ExtractLoginName(alumni.MatricNo)
		if err != nil {
			return nil, err
		}
		alumni.Password = loginName
	}

	hashedPassword, err := utils.HashPassword(alumni.Password)
	if err != nil {
		return nil, err
	}
	alumni.Password = hashedPassword

	if err := s.db.Create(alumni).Error; err != nil {
		return nil, err
	}
	return alumni, nil
}

func (s *Service) UpdateAlumni(id uint, updatedAlumni *Alumni) (*Alumni, error) {
	var alumni Alumni
	if err := s.db.First(&alumni, id).Error; err != nil {
		return nil, err
	}

	if updatedAlumni.MatricNo != "" && updatedAlumni.MatricNo != alumni.MatricNo {
		var existingAlumni Alumni
		if err := s.db.Where("matric_no = ?", updatedAlumni.MatricNo).First(&existingAlumni).Error; err == nil {
			return nil, errors.New("duplicated matric_no")
		}
		alumni.MatricNo = updatedAlumni.MatricNo
	}

	reflectValue := reflect.ValueOf(updatedAlumni).Elem()
	reflectAlumni := reflect.ValueOf(&alumni).Elem()

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Field(i)
		if field.IsValid() && field.CanSet() && !field.IsZero() {
			reflectAlumni.Field(i).Set(field)
		}
	}

	if err := s.db.Save(&alumni).Error; err != nil {
		return nil, err
	}
	return &alumni, nil
}


func (s *Service) DeleteAlumni(id int) error {
	return s.db.Model(&Alumni{}).Where("id = ?", id).Update("is_hidden", true).Error
}

func (s *Service) CheckAlumniExistence(matricNo string) (Alumni, bool, error) {
	var alumni Alumni
	search := matricNo + "%"
	result := s.db.Where("LOWER(matric_no) ILIKE LOWER(?)", search).First(&alumni)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return alumni, false, nil
	}
	return alumni, true, nil
}

func (s *Service) GetAlumniByEmail(email string) (*Alumni, error) {
	var alu Alumni
	if err := s.db.Where("email = ?", email).First(&alu).Error; err != nil {
		return nil, err
	}
	return &alu, nil
}

func (s *Service) GetAlumniByMatricNo(matric_no string) (*Alumni, error) {
	var alu Alumni
	log.Printf("matric_no: %s", matric_no)

	searchPattern := matric_no + "%"

	if err := s.db.Where("matric_no ILIKE ?", searchPattern).First(&alu).Error; err != nil {
		return nil, err
	}
	return &alu, nil
}

func (s *Service) GetAlumniByToken(token string) (*Alumni, error) {
	var alu Alumni
	if err := s.db.Where("token = ?", token).First(&alu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &alu, nil
}

func (s *Service) GetAlumnusByID(id uint) (*Alumni, error) {
	var alumni Alumni
	if err := s.db.First(&alumni, id).Error; err != nil {
		return nil, err
	}
	return &alumni, nil
}

func (s *Service) SendUpdateReminders() error {
    sixMonthsAgo := time.Now().AddDate(0, -6, 0)

    var alumni []Alumni
    if err := s.db.Where("updated_at < ?", sixMonthsAgo).Find(&alumni).Error; err != nil {
        return fmt.Errorf("failed to fetch alumni: %w", err)
    }

    for _, alum := range alumni {
        subject := "[AlumniHub] Inforamtion Update Reminder"
        body := fmt.Sprintf(`Dear %s:

We noticed that your alumni information has not been updated for more than 6 months. In order to maintain the accuracy of the information, please log in to the alumni system to update your personal information.

If your information has not changed, please log in to the system to confirm.

Best wishes!

AlumniHub`, alum.Name)

        if err := utils.SendEmail(alum.Email, subject, body, nil); err != nil {
            log.Printf("Failed to send email to %s: %v", alum.Email, err)
            continue
        }
    }

    return nil
}
