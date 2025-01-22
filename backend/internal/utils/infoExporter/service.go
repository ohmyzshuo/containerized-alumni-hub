package infoExporter

import (
	"alumni_hub/internal/alumni"
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

func (s *Service) ExportInfoByYear(convocationYear int) (*excelize.File, error) {
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

		studies, err := s.getStudiesByAlumniID(uint(studyRecord.AlumniID))
		if err != nil {
			return nil, err
		}

		rows = append(rows, s.formatRows(studyRecord, alum, studies)...)
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

func (s *Service) getStudiesByAlumniID(alumniID uint) ([]study.Study, error) {
	var studies []study.Study
	err := s.db.Where("alumni_id = ?", alumniID).Find(&studies).Error
	if err != nil {
		log.Printf("Error retrieving studies for alumni %d: %v", alumniID, err)
		return nil, err
	}
	return studies, nil
}

func (s *Service) formatRows(studyRecord study.Study, alum alumni.Alumni, studies []study.Study) [][]string {
	rows := [][]string{}
	for _, stu := range studies {
		rows = append(rows, []string{
			"", // Placeholder for Serial No, added later
			s.defaultIfEmpty(alum.Name, "N/A"),
			s.defaultIfEmpty(alum.MatricNo, "N/A"),
			s.defaultIfEmpty(studyRecord.LevelOfStudy, "N/A"),
			s.defaultIfEmpty(studyRecord.Programme, "N/A"),
			s.defaultIfEmpty(studyRecord.Supervisor, "N/A"),
			s.defaultIfEmpty(studyRecord.TitleOfThesis, "N/A"),
			s.defaultIfEmpty(fmt.Sprintf("%d", stu.IntakeYear), "N/A"), // 转换为字符串
			s.defaultIfEmpty(studyRecord.IntakeSession, "N/A"),
			s.defaultIfEmpty(fmt.Sprintf("%d", stu.ConvocationYear), "N/A"), // 转换为字符串
			s.defaultIfEmpty(alum.Nationality, "N/A"),
			s.defaultIfEmpty(alum.Ethnicity, "N/A"),
			s.defaultIfEmpty(alum.DOB.Format("2006-01-02"), "N/A"),
			s.defaultIfEmpty(alum.Gender, "N/A"),
			s.defaultIfEmpty(alum.Marital, "N/A"),
			s.defaultIfEmpty(alum.Address, "N/A"),
			s.defaultIfEmpty(alum.Location, "N/A"),
			s.defaultIfEmpty(alum.Email, "N/A"),
			s.defaultIfEmpty(alum.Phone, "N/A"),
			s.defaultIfEmpty(alum.LinkedIn, "N/A"),
			s.defaultIfEmpty(alum.CreatedAt.Format("2006-01-02 15:04:05"), "N/A"),
			s.defaultIfEmpty(alum.UpdatedAt.Format("2006-01-02 15:04:05"), "N/A"),
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
		"Serial No", "Name", "Matric No", "Level of Study", "Programme", "Supervisor", "Title Of Thesis",
		"Intake Year", "Intake Session", "Convocation Year", "Nationality", "Ethnicity", "DOB", "Gender", "Marital", "Address", "Location", "Email", "Phone", "LinkedIn",
		"Created At", "Updated At",
	}

	// Predefined column widths
	columnWidths := map[string]float64{
		"A": 10, "B": 20, "C": 15, "D": 15, "E": 30, "F": 20, "G": 40,
		"H": 15, "I": 15, "J": 15, "K": 20, "L": 15, "M": 20, "N": 10, "O": 15, "P": 40, "Q": 20, "R": 30, "S": 20, "T": 30, "U": 30, "V": 30,
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
