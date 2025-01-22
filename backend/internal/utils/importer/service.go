package importer

import (
	"alumni_hub/internal/alumni"
	"alumni_hub/internal/alumni/study"
	"alumni_hub/internal/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
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
func (s *Service) ImportAlumniStudiesFromExcel(filePath string) (int, int, int, int, int, int, int, int, string, []string, []string, []string, []string, []string, error) {
	var fileIssue string
	var skippedRows []string
	var headIssue []string
	var failedQueryRows []string
	var failedUpdateRows []string
	var failedCreateRows []string
	var ifNeedUpdate bool
	var ifupdated bool
	var readCount, validCount, exsitNo, sameNo, needUpdateNo, updatedNo, newNo, createNo int

	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fileIssue = fmt.Sprintf("Failed to open file")
		fileIssue = "Failed to open file"

		return 0, 0, 0, 0, 0, 0, 0, 0, fileIssue, failedQueryRows, failedUpdateRows, failedCreateRows, skippedRows, headIssue, nil
	}
	// Get the active sheet name
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fileIssue = "Failed to get rows"
		return 0, 0, 0, 0, 0, 0, 0, 0, fileIssue, failedQueryRows, failedUpdateRows, failedCreateRows, skippedRows, headIssue, nil
	}

	expectedHeader := []string{
		"Faculty",
		"Matric No.",
		"Name",
		"Gender",
		"Nationality",
		"Intake Year",
		"Programme",
		"Graduation Year",
		"Title of Thesis",
		"Email",
		"Phone",
	}
	if len(rows) <= 1 {
		fileIssue = "No data found in the file"
		// Return with error message
		return 0, 0, 0, 0, 0, 0, 0, 0, fileIssue, failedQueryRows, failedUpdateRows, failedCreateRows, skippedRows, headIssue, nil
	}

	// Validate header
	if err := isValidHeader(rows[0], &headIssue); err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fileIssue, failedQueryRows, failedUpdateRows, failedCreateRows, skippedRows, headIssue, nil
	}

	for i, row := range rows {
		readCount++
		ifNeedUpdate = false
		ifupdated = false
		if i == 0 {
			continue // Skip header row
		}
		// ensure same length for each row
		if len(row) < len(expectedHeader) {
			// 填充空列
			for len(row) < len(expectedHeader) {
				row = append(row, "")
			}
		}
		matricNo := cleanMatricNo(row[1])
		name := row[2]
		gender := row[3]
		nationality := row[4]
		programme := row[6]
		intakeYear := parseInt(row[5])
		convocationYear := parseInt(row[7])
		titleOfThesis := row[8]
		email := row[9]
		phone := row[10]

		if matricNo == "" || intakeYear == 0 || programme == "" {
			missingFields := []string{}
			if matricNo == "" {
				missingFields = append(missingFields, "Matric No.")
			}
			if intakeYear == 0 {
				missingFields = append(missingFields, "Intake Year")
			}
			if programme == "" {
				missingFields = append(missingFields, "Programme")
			}

			// Lack of necessary
			skippedRows = append(skippedRows, fmt.Sprintf("Row %d: missing required fields (%s)", i+1, strings.Join(missingFields, ", ")))
			continue
		} else {
			validCount++
		}

		// Check if alumni with the same matric_no already exists
		var existingAlumni alumni.Alumni
		err := s.db.Where("matric_no = ?", matricNo).First(&existingAlumni).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

			} else {
				failedQueryRows = append(failedQueryRows, fmt.Sprintf("Failed to query the matric no of row %d", err))
				continue
			}
		}

		if err == nil {
			exsitNo++
			// Alumni exists, check for existing study record
			var existingStudies []study.Study
			err = s.db.Where("alumni_id = ? AND intake_year = ? AND programme = ?", existingAlumni.ID, intakeYear, programme).Find(&existingStudies).Error
			if err != nil {
				failedQueryRows = append(failedQueryRows, fmt.Sprintf("Membership exitsts for row %d, but failed to query the intake year and program", err))
				continue
			}

			// Update existing alumni fields if they are not the same and not null
			if name != existingAlumni.Name {
				existingAlumni.Name = name
				ifNeedUpdate = true
			}
			if gender != existingAlumni.Gender {
				existingAlumni.Gender = gender
				ifNeedUpdate = true
			}
			if email != existingAlumni.Email {
				existingAlumni.Email = email
				ifNeedUpdate = true
			}
			if nationality != existingAlumni.Nationality {
				existingAlumni.Nationality = nationality
				ifNeedUpdate = true
			}
			if phone != existingAlumni.Phone {
				existingAlumni.Phone = phone
				ifNeedUpdate = true
			}

			if ifNeedUpdate {
				if err := s.db.Save(&existingAlumni).Error; err != nil {
					failedUpdateRows = append(failedUpdateRows, fmt.Sprintf("Membership exitsts for row %d, but failed to update the basic info", err))
				} else {
					ifupdated = true
				}
			}

			if len(existingStudies) > 0 {
				// Update existing study records with new data
				for _, studyRecord := range existingStudies {
					if convocationYear != studyRecord.ConvocationYear {
						studyRecord.ConvocationYear = convocationYear // Update convocation year
						ifNeedUpdate = true
					}
					if titleOfThesis != studyRecord.TitleOfThesis {
						studyRecord.TitleOfThesis = titleOfThesis // Update title of thesis
						ifNeedUpdate = true
					}
					if ifNeedUpdate {
						if err := s.db.Save(&studyRecord).Error; err != nil {
							failedUpdateRows = append(failedUpdateRows, fmt.Sprintf("Membership exitsts for row %d, but failed to update the graduation year and title of thesis", err))
						} else {
							ifupdated = true
						}
					}

				}
			} else {
				// Create new study record
				newStudy := study.Study{
					AlumniID:        existingAlumni.ID,
					FacultyID:       3,
					Programme:       programme,
					IntakeYear:      intakeYear,
					ConvocationYear: convocationYear,
					TitleOfThesis:   titleOfThesis,
				}
				ifNeedUpdate = true
				if err := s.db.Create(&newStudy).Error; err != nil {
					failedUpdateRows = append(failedUpdateRows, fmt.Sprintf("Membership exitsts for row %d, but failed to create the new graduation year and title of thesis", err))
				} else {
					ifupdated = true
				}
			}
			if ifNeedUpdate {
				needUpdateNo++
				if ifupdated {
					updatedNo++
				}

			} else {
				if existingAlumni.IsHidden == false {
					sameNo++
				} else {
					needUpdateNo++
					existingAlumni.IsHidden = false
					if err := s.db.Save(&existingAlumni).Error; err != nil {
						fmt.Printf("Failed to update IsHidden for alumni: %v\n", err)
					} else {
						fmt.Println("Successfully updated IsHidden to false.")
						updatedNo++

					}
				}
			}

		} else {
			newNo++
			// Create new alumni if not found
			newAlumni := alumni.Alumni{
				Name:        name,
				Nationality: nationality,
				Gender:      gender,
				Email:       email,
				MatricNo:    matricNo,
				Phone:       phone,
			}

			// Set default password
			loginName, err := utils.ExtractLoginName(newAlumni.MatricNo)
			if err != nil {
				log.Printf("failed to extract login name: %v", err)
			}
			hashedPassword, err := utils.HashPassword(loginName)
			if err != nil {
				log.Printf("failed to hash password: %v", err)
			}
			newAlumni.Password = hashedPassword

			if err := s.db.Create(&newAlumni).Error; err != nil {
				failedCreateRows = append(failedCreateRows, fmt.Sprintf("Failed to create membership for row %d", err))
			} else {
				createNo++
				// Create initial study
				newStudy := study.Study{
					AlumniID:        newAlumni.ID,
					FacultyID:       3,
					Programme:       programme,
					IntakeYear:      intakeYear,
					ConvocationYear: convocationYear,
					TitleOfThesis:   titleOfThesis,
				}
				if err = s.db.Create(&newStudy).Error; err != nil {
					failedCreateRows = append(failedCreateRows, fmt.Sprintf("Create membership for row %d successfully, but failed to create study reord", err))
				}
			}
		}
	}

	return readCount, validCount, exsitNo, sameNo, needUpdateNo, updatedNo, newNo, createNo, fileIssue, failedQueryRows, failedUpdateRows, failedCreateRows, skippedRows, headIssue, nil
}

func isValidHeader(header []string, headIssue *[]string) error {
	if len(header) < 11 {
		*headIssue = append(*headIssue, "Header is incomplete, expected at least 11 columns.")
		return fmt.Errorf("header validation failed")
	}
	if header[0] != "Faculty" {
		*headIssue = append(*headIssue, fmt.Sprintf("A column should be 'Faculty', got '%s'", header[0]))
	}
	if header[1] != "Matric No." {
		*headIssue = append(*headIssue, fmt.Sprintf("B column should be 'Matric No.', got '%s'", header[1]))
	}
	if header[2] != "Name" {
		*headIssue = append(*headIssue, fmt.Sprintf("C column should be 'Name', got '%s'", header[2]))
	}
	if header[3] != "Gender" {
		*headIssue = append(*headIssue, fmt.Sprintf("D column should be 'Gender', got '%s'", header[3]))
	}
	if header[4] != "Nationality" {
		*headIssue = append(*headIssue, fmt.Sprintf("E column should be 'Nationality', got '%s'", header[4]))
	}
	if header[5] != "Intake Year" {
		*headIssue = append(*headIssue, fmt.Sprintf("F column should be 'Intake Year', got '%s'", header[5]))
	}
	if header[6] != "Programme" {
		*headIssue = append(*headIssue, fmt.Sprintf("G column should be 'Programme', got '%s'", header[6]))
	}
	if header[7] != "Graduation Year" {
		*headIssue = append(*headIssue, fmt.Sprintf("H column should be 'Graduation Year', got '%s'", header[7]))
	}
	if header[8] != "Title of Thesis" {
		*headIssue = append(*headIssue, fmt.Sprintf("I column should be 'Title of Thesis', got '%s'", header[8]))
	}
	if header[9] != "Email" {
		*headIssue = append(*headIssue, fmt.Sprintf("J column should be 'Email', got '%s'", header[9]))
	}
	if header[10] != "Phone" {
		*headIssue = append(*headIssue, fmt.Sprintf("K column should be 'Phone', got '%s'", header[10]))
	}
	if len(*headIssue) > 0 {
		return fmt.Errorf("header validation failed")
	}

	return nil // Header is valid
}

func parseInt(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func cleanMatricNo(matricNo string) string {
	parts := strings.Split(matricNo, "/")
	return parts[0]
}
