package pubExporter

import (
	"alumni_hub/internal/alumni"
	"alumni_hub/internal/alumni/publication"
	"alumni_hub/internal/alumni/study"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"log"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) ExportGraduatesByYear(convocationYear int) (*excelize.File, error) {
	studies, err := s.getStudiesByYear(convocationYear)
	if err != nil {
		return nil, err
	}

	rows := [][]string{}
	for _, studyRecord := range studies {
		alum, err := s.getAlumniByID(uint(studyRecord.AlumniID))
		if err != nil {
			return nil, err
		}

		publications, err := s.getPublicationsByAlumniID(uint(studyRecord.AlumniID))
		if err != nil {
			return nil, err
		}

		rows = append(rows, s.formatRows(studyRecord, alum, publications)...)
	}

	return s.generateExcel(rows, convocationYear)
}

func (s *Service) getStudiesByYear(year int) ([]study.Study, error) {
	var studies []study.Study
	err := s.db.Where("convocation_year = ?", year).Find(&studies).Error
	if err != nil {
		log.Printf("Error retrieving studies records: %v", err)
		return nil, err
	}
	return studies, nil
}

func (s *Service) getAlumniByID(alumniID uint) (alumni.Alumni, error) {
	var alum alumni.Alumni
	err := s.db.First(&alum, alumniID).Error
	if err != nil {
		log.Printf("Error retrieving alumni details for ID %d: %v", alumniID, err)
		return alumni.Alumni{}, err
	}
	return alum, nil
}

func (s *Service) getPublicationsByAlumniID(alumniID uint) ([]publication.Publication, error) {
	var publications []publication.Publication
	err := s.db.Where("alumni_id = ?", alumniID).Find(&publications).Error
	if err != nil {
		log.Printf("Error retrieving publications for alumni %d: %v", alumniID, err)
		return nil, err
	}
	return publications, nil
}

func (s *Service) formatRows(studyRecord study.Study, alum alumni.Alumni, publications []publication.Publication) [][]string {
	rows := [][]string{}
	for _, pub := range publications {
		rows = append(rows, []string{
			"", // Placeholder for Serial No, added later
			s.defaultIfEmpty(alum.Name, "N/A"),
			s.defaultIfEmpty(alum.MatricNo, "N/A"),
			s.defaultIfEmpty(studyRecord.LevelOfStudy, "N/A"),
			s.defaultIfEmpty(studyRecord.Programme, "N/A"),
			s.defaultIfEmpty(studyRecord.Supervisor, "N/A"),
			s.defaultIfEmpty(studyRecord.TitleOfThesis, "N/A"),
			s.defaultIfEmpty(pub.ArticleTitle, "N/A"),
			s.defaultIfEmpty(pub.JournalTitle, "N/A"),
			s.defaultIfEmpty(pub.PublicationType, "N/A"),
			s.defaultIfEmpty(pub.Quartile, "N/A"),
			s.defaultIfEmpty(pub.Status, "N/A"),
			s.defaultIfEmpty(pub.AcceptedDate.Format("2006-01-02"), "N/A"),
			s.defaultIfEmpty(pub.Authors, "N/A"),
			s.defaultIfEmpty(pub.CorrespondingAuthors, "N/A"),
			s.defaultIfEmpty(pub.CreatedAt.Format("2006-01-02 15:04:05"), "N/A"),
			s.defaultIfEmpty(pub.UpdatedAt.Format("2006-01-02 15:04:05"), "N/A"),
		})
	}
	if len(publications) == 0 {
		rows = append(rows, []string{
			"", // Placeholder for Serial No
			s.defaultIfEmpty(alum.Name, "N/A"),
			s.defaultIfEmpty(alum.MatricNo, "N/A"),
			s.defaultIfEmpty(studyRecord.LevelOfStudy, "N/A"),
			s.defaultIfEmpty(studyRecord.Programme, "N/A"),
			s.defaultIfEmpty(studyRecord.Supervisor, "N/A"),
			s.defaultIfEmpty(studyRecord.TitleOfThesis, "N/A"),
			"N/A", "N/A", "N/A", "N/A", "N/A", "N/A", "N/A", "N/A", "N/A", "N/A",
		})
	}
	return rows
}

func (s *Service) defaultIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to generate column labels (A, B, ..., Z, AA, AB, ...)
func getColumnLabel(index int) string {
	label := ""
	for index >= 0 {
		label = string('A'+(index%26)) + label
		index = index/26 - 1
	}
	return label
}

func (s *Service) generateExcel(rows [][]string, year int) (*excelize.File, error) {
	sheetName := fmt.Sprintf("Graduates_%d", year)
	f := excelize.NewFile()
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	// Headers for the table
	headers := []string{
		"Serial No", "Name", "Matric No", "Level of Study", "Programme",
		"Supervisor", "Thesis Title", "Article Title", "Journal Title",
		"Publication Type", "Quartile", "Status", "Accepted Date",
		"Authors", "Corresponding Authors", "Created At", "Updated At",
	}

	// Predefined column widths
	columnWidths := map[string]float64{
		"A": 10, "B": 20, "C": 15, "D": 15, "E": 30,
		"F": 20, "G": 40, "H": 60, "I": 30, "J": 20,
		"K": 10, "L": 15, "M": 15, "N": 40, "O": 40,
		"P": 20, "Q": 20,
	}

	// Apply styles
	commonStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 11},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// 为表头创建独立样式
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true, // 加粗
			Size: 12,   // 字体大小
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D9E1F2"}, // 背景颜色（淡蓝色）
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平居中
			Vertical:   "center", // 垂直居中
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// 设置表头样式
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", getColumnLabel(i))      // 表头所在行，例如 A1, B1...
		f.SetCellValue(sheetName, cell, header)            // 设置表头内容
		f.SetCellStyle(sheetName, cell, cell, headerStyle) // 应用表头样式
	}

	// Set column widths
	for col, width := range columnWidths {
		f.SetColWidth(sheetName, col, col, width)
	}

	// Add data rows
	for i, row := range rows {
		row[0] = fmt.Sprintf("%d", i+1) // Add serial number
		for j, cell := range row {
			col := getColumnLabel(j)
			cellAddress := fmt.Sprintf("%s%d", col, i+2)
			f.SetCellValue(sheetName, cellAddress, cell)
			f.SetCellStyle(sheetName, cellAddress, cellAddress, commonStyle)
		}
	}

	log.Println("Graduates report generated successfully.")
	return f, nil
}
