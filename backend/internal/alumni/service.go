package alumni

type ServiceInterface interface {
	GetAlumnus(id int) (*Alumni, error)
	GetAlumni(page, pageSize int, searchQuery string) ([]Alumni, int64, error)
	CreateAlumni(alumni *Alumni) (*Alumni, error)
	UpdateAlumni(id uint, updatedAlumni *Alumni) (*Alumni, error)
	DeleteAlumni(id int) error
	CheckAlumniExistence(matricNo string) (Alumni, bool, error)
	GetAlumniByEmail(email string) (*Alumni, error)
	GetAlumniByMatricNo(matricNo string) (*Alumni, error)
	GetAlumniByToken(token string) (*Alumni, error)
	GetAlumnusByID(id uint) (*Alumni, error)
	SendUpdateReminders() error
}
