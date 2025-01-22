package publication

import (
	"alumni_hub/internal/alumni"
	"gorm.io/gorm"
	"math"
	"time"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreatePublication(publication *Publication) error {
	return s.db.Create(publication).Error
}

func (s *Service) GetPublicationsByAlumniID(alumniID uint) ([]Publication, error) {
	var publications []Publication
	err := s.db.Where("alumni_id = ?", alumniID).Order("created_at DESC").Find(&publications).Error
	return publications, err
}

func (s *Service) UpdatePublication(id uint, updatedPublication *Publication) (*Publication, error) {
	var publication Publication
	if err := s.db.First(&publication, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&publication).Updates(updatedPublication).Error; err != nil {
		return nil, err
	}

	return &publication, nil
}

func (s *Service) DeletePublication(id uint) error {
	if err := s.db.Delete(&Publication{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) GetPublicationsByToken(token string) ([]Publication, error) {
	var publications []Publication
	var alu alumni.Alumni
	if err := s.db.Where("token = ?", token).First(&alu).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("alumni_id = ?", alu.ID).Find(&publications).Error; err != nil {
		return nil, err
	}
	return publications, nil
}

// GetStatistics fetches publication statistics based on the given filters
func (s *Service) GetStatistics(acceptedDateStart, acceptedDateEnd string, publicationTypes, statuses, quartiles []string, search string, pageSize, currentPage int) ([]Publication, Meta, Statistics, error) {
	var publications []Publication

	query := s.db.Model(&Publication{})

	// Filter by accepted_date_start if provided
	if acceptedDateStart != "" {
		startDate, err := time.Parse("2006-01-02", acceptedDateStart)
		if err != nil {
			return nil, Meta{}, Statistics{}, err
		}
		query = query.Where("accepted_date >= ?", startDate)
	}

	// Filter by accepted_date_end if provided
	if acceptedDateEnd != "" {
		endDate, err := time.Parse("2006-01-02", acceptedDateEnd)
		if err != nil {
			return nil, Meta{}, Statistics{}, err
		}
		query = query.Where("accepted_date <= ?", endDate)
	}

	// Filter by publication_type if provided
	if len(publicationTypes) > 0 {
		query = query.Where("publication_type IN ?", publicationTypes)
	}

	// Filter by status if provided
	if len(statuses) > 0 {
		query = query.Where("status IN ?", statuses)
	}

	// Filter by quartile if provided
	if len(quartiles) > 0 {
		query = query.Where("quartile IN ?", quartiles)
	}

	// Filter by search if provided
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("journal_title ILIKE ? OR article_title ILIKE ? OR authors ILIKE ? OR corresponding_authors ILIKE ?", searchPattern, searchPattern, searchPattern, searchPattern)
	}

	// Calculate total count for pagination
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, Meta{}, Statistics{}, err
	}

	// Apply pagination
	offset := (currentPage - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// Fetch the filtered publications
	if err := query.Find(&publications).Error; err != nil {
		return nil, Meta{}, Statistics{}, err
	}

	// Fetch alumni names and add to publications
	alumniMap := make(map[uint]string)
	for i := range publications {
		if name, exists := alumniMap[publications[i].AlumniID]; exists {
			publications[i].AlumniName = name
		} else {
			var alu Alumni
			if err := s.db.First(&alu, publications[i].AlumniID).Error; err == nil {
				publications[i].AlumniName = alu.Name
				alumniMap[publications[i].AlumniID] = alu.Name
			}
		}
	}

	// Calculate statistics
	var statistics Statistics
	statistics.PublicationTypeCount = make(map[string]int)
	statistics.StatusCount = make(map[string]int)
	statistics.QuartileCount = make(map[string]int)

	for _, pub := range publications {
		statistics.PublicationTypeCount[pub.PublicationType]++
		statistics.StatusCount[pub.Status]++
		statistics.QuartileCount[pub.Quartile]++
	}

	// Calculate pagination meta
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	meta := Meta{
		Pagination: Pagination{
			Page:       currentPage,
			PageSize:   pageSize,
			Total:      int(total),
			TotalPages: totalPages,
		},
		Statistics: statistics,
	}

	return publications, meta, statistics, nil
}
