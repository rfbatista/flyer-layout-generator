package repositories

func NewCompanyRepository() (CompanyRepository, error) {
	return CompanyRepository{}, nil
}

type CompanyRepository struct{}

func (c CompanyRepository) GetByID() {}
