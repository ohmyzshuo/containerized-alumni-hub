package content

import (
	alu "alumni_hub/internal/alumni"
	"alumni_hub/internal/staff"
	"bytes"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
	"gorm.io/gorm"
	"html/template"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateContent(content *Content, attachments []Attachment, tags []string) error {
	if attachments == nil {
		attachments = []Attachment{}
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(content).Error; err != nil {
			return err
		}

		if len(attachments) > 0 {
			for i := range attachments {
				attachments[i].ContentId = content.ID
			}
			if err := tx.Create(&attachments).Error; err != nil {
				return err
			}
		}

		for _, tagName := range tags {
			var tag Tag
			if err := tx.Where("name = ?", tagName).FirstOrCreate(&tag, Tag{Name: tagName}).Error; err != nil {
				return err
			}

			contentTag := ContentTag{
				ContentID: content.ID,
				TagID:     tag.ID,
			}
			if err := tx.Create(&contentTag).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *Service) SendContentEmail(contentID uint, alumniIDs []uint) error {
	var content Content
	if err := s.db.First(&content, contentID).Error; err != nil {
		return err
	}

	var attachments []Attachment
	if err := s.db.Where("content_id = ?", contentID).Find(&attachments).Error; err != nil {
		return err
	}

	var alumni []alu.Alumni
	if err := s.db.Where("id IN ?", alumniIDs).Find(&alumni).Error; err != nil {
		return err
	}

	var emails []string
	for _, a := range alumni {
		emails = append(emails, a.Email)
	}

	title := "[AlumniHub] " + content.Title

	for _, em := range emails {
		if err := SendEmail(em, title, content.Description, attachments, false); err != nil {
			return err
		}
	}

	return nil
}
func SendEmail(to string, subject string, body string, attachments []Attachment, isHTML bool) error {
	err := godotenv.Load()
	e := email.NewEmail()
	e.From = os.Getenv("EMAIL_FROM")
	e.To = []string{to}
	e.Subject = subject

	if isHTML {
		e.HTML = []byte(body)
	} else {
		e.Text = []byte(body)
	}

	if attachments != nil {
		for _, attachment := range attachments {
			data, err := os.ReadFile(attachment.AttachmentPath)
			if err != nil {
				return err
			}

			filename := filepath.Base(attachment.AttachmentPath)
			contentType := mime.TypeByExtension(filepath.Ext(filename))

			_, err = e.Attach(bytes.NewReader(data), filename, contentType)
			if err != nil {
				return err
			}
		}
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	err = e.Send(smtpServer+":"+smtpPort, smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer))
	return err
}

func (s *Service) GetContentByID(id uint) (*Content, []Attachment, error) {
	var content Content
	if err := s.db.First(&content, id).Error; err != nil {
		return nil, nil, err
	}

	var attachments []Attachment
	if err := s.db.Where("content_id = ?", id).Find(&attachments).Error; err != nil {
		return nil, nil, err
	}

	return &content, attachments, nil
}

func (s *Service) ListDefault(page, pageSize int, search string) ([]Content, int64, error) {
	var contents []Content
	var total int64

	offset := (page - 1) * pageSize
	query := s.db.Order("created_at DESC").Offset(offset).Limit(pageSize)

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ? OR venue ILIKE ?", searchTerm, searchTerm, searchTerm)
	}

	if err := query.Find(&contents).Error; err != nil {
		return nil, 0, err
	}

	for i, content := range contents {
		var attachments []Attachment
		if err := s.db.Where("content_id = ?", content.ID).Find(&attachments).Error; err != nil {
			return nil, 0, err
		}
		contents[i].Attachments = attachments

		var st staff.Staff
		if err := s.db.Select("name").Where("id = ?", content.CreatedBy).First(&st).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				contents[i].CreatedByName = "-"
			} else {
				return nil, 0, err
			}
		} else {
			contents[i].CreatedByName = st.Name
		}
	}

	if err := s.db.Model(&Content{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return contents, total, nil
}

func (s *Service) ListPinned() ([]Content, error) {
	var contents []Content
	if err := s.db.Where("is_pinned = ?", true).Order("created_at DESC").Find(&contents).Error; err != nil {
		return nil, err
	}

	for i, content := range contents {
		var attachments []Attachment
		if err := s.db.Where("content_id = ?", content.ID).Find(&attachments).Error; err != nil {
			return nil, err
		}
		contents[i].Attachments = attachments
	}

	return contents, nil
}

func (s *Service) UpdateContent(id uint, updatedContent *Content, attachments []Attachment) (*Content, error) {
	var content Content
	if err := s.db.First(&content, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&content).Updates(updatedContent).Error; err != nil {
			return err
		}

		if len(attachments) > 0 {
			if err := tx.Where("content_id = ?", id).Delete(&Attachment{}).Error; err != nil {
				return err
			}

			for i := range attachments {
				attachments[i].ContentId = id
			}

			if err := tx.Create(&attachments).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &content, nil
}

func (s *Service) DeleteContent(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("content_id = ?", id).Delete(&Attachment{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&Content{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *Service) getAlumniTagsVector(alumniID uint) (map[uint]float64, error) {
	var alumniTags []alu.AlumniTag

	if err := s.db.Where("alumni_id = ?", alumniID).Find(&alumniTags).Error; err != nil {
		return nil, err
	}

	tagVector := make(map[uint]float64)
	for _, alumniTag := range alumniTags {
		tagVector[alumniTag.TagID] = alumniTag.Weight
	}

	return tagVector, nil
}

func (s *Service) getContentTagsVector(contentID uint) (map[uint]float64, error) {
	var content Content

	if err := s.db.Preload("Tags").First(&content, contentID).Error; err != nil {
		return nil, err
	}

	tagVector := make(map[uint]float64)
	for _, tag := range content.Tags {
		tagVector[tag.ID] = 1.0
	}

	return tagVector, nil
}

func cosineSimilarity(vec1, vec2 map[uint]float64) float64 {
	var dotProduct, mag1, mag2 float64

	for key, value := range vec1 {
		dotProduct += value * vec2[key]
		mag1 += value * value
	}

	for _, value := range vec2 {
		mag2 += value * value
	}

	if mag1 == 0 || mag2 == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(mag1) * math.Sqrt(mag2))
}

func (s *Service) updateAlumniTagWeights(alumniID uint, contentID uint, like bool) error {
	var content Content

	if err := s.db.Preload("Tags").First(&content, contentID).Error; err != nil {
		return err
	}

	for _, tag := range content.Tags {
		var alumniTag alu.AlumniTag
		if err := s.db.Where("alumni_id = ? AND tag_id = ?", alumniID, tag.ID).First(&alumniTag).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				alumniTag = alu.AlumniTag{
					AlumniID: alumniID,
					TagID:    tag.ID,
					Weight:   1.0,
				}
				if !like {
					alumniTag.Weight = -1.0
				}
				if err := s.db.Create(&alumniTag).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			if like {
				alumniTag.Weight += 1.0
			} else {
				alumniTag.Weight -= 1.0
			}
			if err := s.db.Save(&alumniTag).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) LikeContent(alumniID uint, contentID uint) error {
	if err := s.updateAlumniTagWeights(alumniID, contentID, true); err != nil {
		return err
	}
	return nil
}

func (s *Service) DislikeContent(alumniID uint, contentID uint) error {
	if err := s.updateAlumniTagWeights(alumniID, contentID, false); err != nil {
		return err
	}
	return nil
}

func (s *Service) ListAllByPreferences(alumniID uint, page int, pageSize int) ([]Content, int64, error) {
	var contents []Content

	alumniTagsVector, err := s.getAlumniTagsVector(alumniID)
	if err != nil {
		return nil, 0, err
	}

	if err := s.db.Preload("Tags").Find(&contents).Error; err != nil {
		return nil, 0, err
	}

	contentScores := make(map[uint]float64)
	for _, content := range contents {
		contentTagsVector, err := s.getContentTagsVector(content.ID)
		if err != nil {
			return nil, 0, err
		}
		contentScores[content.ID] = cosineSimilarity(alumniTagsVector, contentTagsVector)
	}

	sort.Slice(contents, func(i, j int) bool {
		return contentScores[contents[i].ID] > contentScores[contents[j].ID]
	})

	total := int64(len(contents))
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > int(total) {
		start = int(total)
	}
	if end > int(total) {
		end = int(total)
	}

	pagedContents := contents[start:end]

	return pagedContents, total, nil
}

func (s *Service) ChangeAlumniStatusForEvent(alumniID uint, contentID uint, status uint, comment string) error {
	if status < 0 || status > 5 {
		return errors.New("status must be between 0 and 5")
	}

	var eventParticipant EventParticipant
	err := s.db.Where("alumni_id = ? AND content_id = ?", alumniID, contentID).First(&eventParticipant).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		eventParticipant = EventParticipant{
			AlumniID:  alumniID,
			ContentID: contentID,
			Status:    status,
			Comment:   comment,
		}
		if err := s.db.Create(&eventParticipant).Error; err != nil {
			return err
		}
	} else if err == nil {
		eventParticipant.Status = status
		eventParticipant.Comment = comment
		if err := s.db.Save(&eventParticipant).Error; err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (s *Service) GetEventParticipants(contentID int, status *int) ([]alu.Alumni, error) {
	var participants []EventParticipant
	query := s.db.Where("content_id = ?", contentID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if err := query.Find(&participants).Error; err != nil {
		return nil, err
	}

	var alumniIDs []uint
	participantInfoMap := make(map[uint]struct {
		Status  uint
		Comment string
	})

	for _, participant := range participants {
		alumniIDs = append(alumniIDs, participant.AlumniID)
		participantInfoMap[participant.AlumniID] = struct {
			Status  uint
			Comment string
		}{
			Status:  participant.Status,
			Comment: participant.Comment,
		}
	}

	var alumni []alu.Alumni
	if err := s.db.Where("id IN (?)", alumniIDs).Find(&alumni).Error; err != nil {
		return nil, err
	}

	for i, alumnus := range alumni {
		if info, ok := participantInfoMap[alumnus.ID]; ok {
			alumni[i].ParticipantStatus = info.Status
			alumni[i].ParticipantComment = info.Comment
		}
	}

	return alumni, nil
}
func (s *Service) GetAlumnusParticipation(alumniID int, status *int, page, pageSize int) ([]Content, int64, error) {
	var participants []EventParticipant
	query := s.db.Where("alumni_id = ?", alumniID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if err := query.Find(&participants).Error; err != nil {
		return nil, 0, err
	}

	var contentIDs []uint
	contentStatusMap := make(map[uint]uint)
	for _, participant := range participants {
		contentIDs = append(contentIDs, participant.ContentID)
		contentStatusMap[participant.ContentID] = participant.Status
	}

	var contents []Content
	offset := (page - 1) * pageSize
	if err := s.db.Where("id IN (?)", contentIDs).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&contents).Error; err != nil {
		return nil, 0, err
	}

	for i, content := range contents {
		var attachments []Attachment
		if err := s.db.Where("content_id = ?", content.ID).Find(&attachments).Error; err != nil {
			return nil, 0, err
		}
		contents[i].Attachments = attachments

		if status, ok := contentStatusMap[content.ID]; ok {
			contents[i].Status = status
		}
	}

	var total int64
	if err := s.db.Model(&Content{}).Where("id IN (?)", contentIDs).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return contents, total, nil
}

type imageInfo struct {
	tempFilePath string
	cid          string
}

func (s *Service) SendNewsletter(contentIDs []uint, alumniIDs []uint) error {
	var contents []Content
	if err := s.db.Where("id IN ?", contentIDs).Find(&contents).Error; err != nil {
		return err
	}

	contentAttachments := make(map[uint][]Attachment)
	for _, content := range contents {
		var attachments []Attachment
		if err := s.db.Where("content_id = ?", content.ID).Find(&attachments).Error; err != nil {
			return err
		}
		contentAttachments[content.ID] = attachments
	}

	var alumni []alu.Alumni
	if err := s.db.Where("id IN ?", alumniIDs).Find(&alumni).Error; err != nil {
		return err
	}

	imageMap := make(map[uint]imageInfo)
	defer func() {
		for _, info := range imageMap {
			os.Remove(info.tempFilePath)
		}
	}()

	for _, content := range contents {
		if attachments := contentAttachments[content.ID]; len(attachments) > 0 {
			cid := fmt.Sprintf("image_%d", content.ID)

			baseURL := os.Getenv("BASE_URL")
			if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
				baseURL = "http://" + baseURL
			}
			imageURL := baseURL + attachments[0].AttachmentPath

			resp, err := http.Get(imageURL)
			if err != nil {
				continue
			}

			imageData, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				continue
			}

			tempFile, err := os.CreateTemp("", "newsletter_image_*.jpg")
			if err != nil {
				continue
			}

			if _, err := tempFile.Write(imageData); err != nil {
				tempFile.Close()
				continue
			}
			tempFile.Close()

			imageMap[content.ID] = imageInfo{
				tempFilePath: tempFile.Name(),
				cid:          cid,
			}
		}
	}

	for _, alum := range alumni {
		htmlContent, err := s.generateNewsletterHTML(contents, contentAttachments, imageMap, alum.ID)
		if err != nil {
			return err
		}

		now := time.Now()
		dateStr := now.Format("Jan 2, 2006")
		e := email.NewEmail()
		e.From = os.Getenv("EMAIL_FROM")
		e.To = []string{alum.Email}
		e.Subject = "[AlumniHub] Newsletter - " + dateStr
		e.HTML = []byte(htmlContent)

		for _, info := range imageMap {
			f, err := os.Open(info.tempFilePath)
			if err != nil {
				return err
			}

			content, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				return err
			}

			e.Attach(bytes.NewReader(content), filepath.Base(info.tempFilePath), "image/jpeg")

			attachment := e.Attachments[len(e.Attachments)-1]
			attachment.Header.Set("Content-ID", "<"+info.cid+">")
			attachment.Header.Set("Content-Disposition", "inline")
		}

		smtpServer := os.Getenv("SMTP_SERVER")
		smtpPort := os.Getenv("SMTP_PORT")
		smtpUser := os.Getenv("SMTP_USER")
		smtpPassword := os.Getenv("SMTP_PASSWORD")

		err = e.Send(smtpServer+":"+smtpPort, smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer))
		if err != nil {
			return err
		}
	}

	return nil
}

type TemplateContent struct {
	ID                uint
	ContentType       uint
	Title             string
	Description       string
	Venue             string
	ParticipantQuota  uint
	CreatedAt         time.Time
	FirstImage        string
	NotInterestedLink string
	InterestedLink    string
	RegisterLink      string
}

func (c TemplateContent) ProcessFirstImage() string {
	if c.FirstImage == "" {
		return ""
	}
	return c.FirstImage
}

func (c TemplateContent) TypeString() string {
	types := map[uint]string{
		1: "Event",
		2: "Announcement",
		3: "Advertisement",
		0: "Other",
	}
	return types[c.ContentType]
}

func (c TemplateContent) TruncatedDescription() string {
	if len(c.Description) > 1000 {
		return c.Description[:1000] + "..."
	}
	return c.Description
}

func (c TemplateContent) FormattedDate() string {
	return c.CreatedAt.Format("06-01-2006")
}

func (c TemplateContent) ShowVenue() bool {
	return c.ContentType == 1 && c.Venue != ""
}

func (c TemplateContent) ShowQuota() bool {
	return c.ContentType == 1
}

func (c TemplateContent) QuotaDisplay() string {
	if c.ParticipantQuota == 0 {
		return "Unlimited"
	}
	return fmt.Sprintf("%d", c.ParticipantQuota)
}
func (s *Service) generateNewsletterHTML(contents []Content, contentAttachments map[uint][]Attachment, imageMap map[uint]imageInfo, alumniID uint) (string, error) {
	templateContents := make([]TemplateContent, 0, len(contents))
	for _, content := range contents {
		var firstImage string
		if attachments := contentAttachments[content.ID]; len(attachments) > 0 {
			if info, exists := imageMap[content.ID]; exists {
				firstImage = info.cid
			}
		}

		notInterestedLink, interestedLink, registerLink, err := generateEmailLinks(alumniID, content.ID)
		if err != nil {
			return "", err
		}

		templateContents = append(templateContents, TemplateContent{
			ID:                content.ID,
			ContentType:       content.ContentType,
			Title:             content.Title,
			Description:       content.Description,
			Venue:             content.Venue,
			ParticipantQuota:  content.ParticipantQuota,
			CreatedAt:         content.CreatedAt,
			FirstImage:        firstImage,
			NotInterestedLink: notInterestedLink,
			InterestedLink:    interestedLink,
			RegisterLink:      registerLink,
		})
	}

	funcMap := template.FuncMap{
		"typeString":           func(c TemplateContent) string { return c.TypeString() },
		"truncatedDescription": func(c TemplateContent) string { return c.TruncatedDescription() },
		"formattedDate":        func(c TemplateContent) string { return c.FormattedDate() },
		"showVenue":            func(c TemplateContent) bool { return c.ShowVenue() },
		"showQuota":            func(c TemplateContent) bool { return c.ShowQuota() },
		"quotaDisplay":         func(c TemplateContent) string { return c.QuotaDisplay() },
		"processFirstImage": func(c TemplateContent) string {
			if c.FirstImage == "" {
				return ""
			}
			return "cid:" + c.FirstImage
		},
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"safeURL": func(s string) template.URL {
			return template.URL(s)
		},
	}

	tmpl, err := template.New("newsletter").Funcs(funcMap).Parse(emailTemplate)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, struct{ Contents []TemplateContent }{templateContents}); err != nil {
		return "", err
	}
	log.Print(body.String())
	return body.String(), nil
}

const emailTemplate = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <style>
        .content-card {
            margin: 20px auto;
            padding: 15px;
            max-width: 600px;
            border: 1px solid #ddd;
            border-radius: 8px;
            background-color: #fff;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
		.content-type {
		        display: inline-block;
		        padding: 4px 8px;
		        border-radius: 4px;
		        font-size: 12px;
		        font-weight: bold;
		        margin-bottom: 10px;
		        background-color: #e8f0fe;
		        color: #1a73e8;
		}
		.content-title {
		        color: #333;
		        font-size: 20px;
		        margin: 10px 0;
		}
	    .content-description {
	        color: #666;
	        font-size: 16px;
	        line-height: 1.5;
	        margin-bottom: 10px;
	    }
	    .content-details {
	        color: #888;
	        font-size: 14px;
	        margin-top: 10px;
	    }
	    .content-image {
	        width: 100%;
	        max-height: 300px;
	        object-fit: cover;
	        border-radius: 4px;
	        margin: 10px 0;
	    }
	    .detail-item {
	        margin: 5px 0;
	    }
	    @media screen and (max-width: 480px) {
	        .content-card {
	            margin: 10px;
	            padding: 10px;
	        }
	        .content-title {
	            font-size: 18px;
	        }
	        .content-description {
	            font-size: 14px;
	        }
	    }
        .btn {
            display: inline-block;
            padding: 10px 10px;
            text-decoration: none;
            border-radius: 8px;
            text-align: center;
            font-weight: bold;
            color: white;
        }
        .btn-register {
            background-color: #7EACB5;
        }
        .btn-interested {
            background-color: #FADFA1;
        }
        .btn-not-interested {
            background-color: #C96868;
        }

    </style>
</head>
<body style="background-color: #f5f5f5; padding: 20px;">
    <div style="max-width: 600px; margin: 0 auto;">
        {{range .Contents}}
        <div class="content-card">
            	<div class="content-type">{{typeString .}}</div>
			    <h2 class="content-title">{{.Title}}</h2>
				{{with processFirstImage .}}
				    <img class="content-image" src="{{. | safeURL}}" alt="Content Image" style="max-width: 100%; height: auto;">
				{{end}}
			    <div class="content-description">{{truncatedDescription .}}</div>
            <div class="content-details">
                {{if showVenue .}}
                <div class="detail-item">Venue: {{.Venue}}</div>
                {{end}}
                {{if showQuota .}}
                <div class="detail-item">Participant Quota: {{quotaDisplay .}}</div>
                {{end}}
                <div class="detail-item">Created: {{formattedDate .}}</div>
            </div>
            <div class="btn-container">
				{{if eq .ContentType 1}}
				    <a href="{{.RegisterLink | safeURL}}" class="btn btn-register">Register</a>
				    <a href="{{.InterestedLink | safeURL}}" class="btn btn-interested">Interested But Not Available</a>
				    <a href="{{.NotInterestedLink | safeURL}}" class="btn btn-not-interested">Not Interested</a>
				{{end}}
            </div>
        </div>
        {{end}}
    </div>
</body>
</html>
`
