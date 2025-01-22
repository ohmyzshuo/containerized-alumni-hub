package pubImporter

import (
	"alumni_hub/internal/alumni"
	"alumni_hub/internal/alumni/publication"
	"alumni_hub/internal/alumni/study"
	"alumni_hub/internal/utils"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) ImportAlumuiPublicationsFromExcel(filePath string) (int, int, []string, error) {
	var skippedRows []string

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		skippedRows = append(skippedRows, fmt.Sprintf("Failed to open file: %v", err))
		return 0, 0, skippedRows, nil
	}

	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
	rows, err := f.GetRows(sheetName)
	if err != nil {
		skippedRows = append(skippedRows, fmt.Sprintf("Failed to get rows: %v", err))
		return 0, 0, skippedRows, nil
	}

	if len(rows) < 1 {
		skippedRows = append(skippedRows, "No rows found in the file")
		return 0, 0, skippedRows, nil
	}

	if err := isValidHeader(rows[0], &skippedRows); err != nil {
		return 0, 0, skippedRows, nil
	}

	var importedCount, invalidCount int

	for i, row := range rows[1:] {
		if len(row) < 3 {
			skippedRows = append(skippedRows, fmt.Sprintf("Row %d: insufficient columns", i+2))
			invalidCount++
			continue
		}

		name := row[0]
		matricNo := cleanMatricNo(row[1])
		levelOfStudy := row[2] // Changed from levelOfStudy to programme to match the example

		if matricNo == "" {
			skippedRows = append(skippedRows, fmt.Sprintf("Row %d: missing required field (Matric No.)", i+2))
			invalidCount++
			continue
		}

		// Find or create alumni
		var existingAlumni alumni.Alumni
		err := s.db.Where("matric_no = ?", matricNo).First(&existingAlumni).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create new alumni if not found
				loginName, err := utils.ExtractLoginName(matricNo)
				if err != nil {
					skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to extract login name from matric no %s", i+2, matricNo))
					invalidCount++
					log.Printf("failed to extract login name: %v", err)
					continue
				}
				hashedPassword, err := utils.HashPassword(loginName)
				if err != nil {
					skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to hash password for matric no %s", i+2, matricNo))
					invalidCount++
					log.Printf("failed to hash password: %v", err)
					continue
				}

				existingAlumni = alumni.Alumni{
					Name:     name,
					MatricNo: matricNo,
					Password: hashedPassword,
				}

				if err := s.db.Create(&existingAlumni).Error; err != nil {
					skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to create new alumni for matric no %s: %v", i+2, matricNo, err))
					invalidCount++
					continue
				}
			} else {
				skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to query alumni with matric no %s: %v", i+2, matricNo, err))
				invalidCount++
				continue
			}
		} else {
			// Update existing alumni
			if name != "" {
				existingAlumni.Name = name
			}

			if err := s.db.Save(&existingAlumni).Error; err != nil {
				skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to update alumni info for matric no %s: %v", i+2, matricNo, err))
				invalidCount++
				continue
			}
		}

		// Process study record (following the example's structure)
		if levelOfStudy != "" {
			var existingStudy study.Study
			err = s.db.Where("alumni_id = ? AND level_of_study = ?", existingAlumni.ID, levelOfStudy).First(&existingStudy).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// Create new study record
					newStudy := study.Study{
						AlumniID:     existingAlumni.ID,
						LevelOfStudy: levelOfStudy,
					}
					if err := s.db.Create(&newStudy).Error; err != nil {
						log.Printf("failed to create study record: %v", err)
						invalidCount++
					}
				}
			} else {
				// Update existing study record if needed
				if err := s.db.Save(&existingStudy).Error; err != nil {
					log.Printf("failed to update study record: %v", err)
					invalidCount++
				}
			}
		}

		// Process publications
		for j := 3; j < len(row); j += 3 {
			if j+2 >= len(row) {
				break
			}

			pubTitle := row[j]
			pubType := row[j+1]
			quartile := row[j+2]

			if pubTitle == "" {
				continue // Skip empty publication entries
			}

			// Find or create publication
			var existingPub publication.Publication
			err := s.db.Where("alumni_id = ? AND article_title = ?", existingAlumni.ID, pubTitle).First(&existingPub).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// Create new publication
					newPub := publication.Publication{
						AlumniID:        existingAlumni.ID,
						ArticleTitle:    pubTitle,
						PublicationType: pubType,
						Quartile:        quartile,
					}
					if err := s.db.Create(&newPub).Error; err != nil {
						skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to create new publication '%s': %v", i+2, pubTitle, err))
						invalidCount++
					}
				} else {
					skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to query publication '%s': %v", i+2, pubTitle, err))
					invalidCount++
				}
			} else {
				// Update existing publication
				if pubType != "" {
					existingPub.PublicationType = pubType
				}
				if quartile != "" {
					existingPub.Quartile = quartile
				}
				if err := s.db.Save(&existingPub).Error; err != nil {
					skippedRows = append(skippedRows, fmt.Sprintf("Row %d: failed to update publication '%s': %v", i+2, pubTitle, err))
					invalidCount++
				}
			}
		}
		importedCount++
	}

	return importedCount, invalidCount, skippedRows, nil
}
func cleanMatricNo(matricNo string) string {
	parts := strings.Split(matricNo, "/")
	return parts[0]
}

func isValidHeader(header []string, skippedRows *[]string) error {
	// Check the first three columns
	if len(header) < 3 {
		return fmt.Errorf("header is too short, expected at least 3 columns")
	}
	if header[0] != "Name" {
		*skippedRows = append(*skippedRows, fmt.Sprintf("A column should be 'Name', got '%s'", header[0]))
	}
	if header[1] != "Matric No." {
		*skippedRows = append(*skippedRows, fmt.Sprintf("B column should be 'Matric No.', got '%s'", header[1]))
	}
	if header[2] != "Level of Study" {
		*skippedRows = append(*skippedRows, fmt.Sprintf("C column should be 'Level of Study', got '%s'", header[2]))
	}

	// Check the subsequent columns for the correct pattern
	for i := 3; i < len(header); i++ {
		columnLetter := string('A' + i) // Convert index to corresponding uppercase letter
		switch (i - 3) % 3 {
		case 0:
			if header[i] != "Article Title" {
				*skippedRows = append(*skippedRows, fmt.Sprintf("%s column should be 'Article Title', got '%s'", columnLetter, header[i]))
			}
		case 1:
			if header[i] != "Publication Type" {
				*skippedRows = append(*skippedRows, fmt.Sprintf("%s column should be 'Publication Type', got '%s'", columnLetter, header[i]))
			}
		case 2:
			if header[i] != "Quartile" {
				*skippedRows = append(*skippedRows, fmt.Sprintf("%s column should be 'Quartile', got '%s'", columnLetter, header[i]))
			}
		}
	}

	if len(*skippedRows) > 0 {
		return fmt.Errorf("header validation failed")
	}

	return nil // Header is valid
}
