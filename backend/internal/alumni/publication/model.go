package publication

import "time"

type Publication struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	CreatedAt            time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ArticleTitle         string    `gorm:"type:varchar(255);not null" json:"article_title"`
	JournalTitle         string    `gorm:"type:varchar(255)" json:"journal_title"`
	PublicationType      string    `gorm:"type:publication_type;not null" json:"publication_type"`
	SequenceNo           int       `gorm:"type:int" json:"sequence_no"`
	AlumniID             uint      `gorm:"not null" json:"alumni_id"`
	Quartile             string    `gorm:"type:varchar(10)" json:"quartile"`
	Status               string    `gorm:"type:publication_status;not null" json:"status"`
	AcceptedDate         time.Time `gorm:"type:date" json:"accepted_date"`
	Authors              string    `gorm:"type:varchar(255);not null" json:"authors"`
	CorrespondingAuthors string    `gorm:"type:varchar(255);not null" json:"corresponding_authors"`
	AlumniName           string    `json:"alumni_name" gorm:"-"`
}

// Alumni represents an alumni record for this package use
type Alumni struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Pagination meta data
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// Meta data
type Meta struct {
	Pagination Pagination `json:"pagination"`
	Statistics Statistics `json:"statistics"`
}

// Statistics data
type Statistics struct {
	PublicationTypeCount map[string]int `json:"publicationTypeCount"`
	StatusCount          map[string]int `json:"statusCount"`
	QuartileCount        map[string]int `json:"quartileCount"`
}
