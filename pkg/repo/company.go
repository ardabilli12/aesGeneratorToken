package repo

import (
	"log"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

type CompanyRepositoryContract interface {
	GetDetailCompanyById(companyId string) Companies
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (c *CompanyRepository) GetDetailCompanyById(companyId string) Companies {
	var company Companies
	err := c.db.Where("uuid = ?", companyId).First(&company).Error
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	return company
}

type Companies struct {
	Uuid                    string `json:"uuid"`
	CompanyName             string `json:"company_name"`
	CompanyEmail            string `json:"company_email"`
	CompanyPhone            string `json:"company_phone"`
	CompanyAddress1         string `json:"company_address1"`
	CompanyAddress2         string `json:"company_address2"`
	CompanyCity             string `json:"company_city"`
	CompanyState            string `json:"company_state"`
	CompanyZipCode          string `json:"company_zip_code"`
	CompanyCountry          string `json:"company_country"`
	CompanyContact          string `json:"company_contact"`
	CompanyVat              string `json:"company_vat"`
	CompanyWebsite          string `json:"company_website"`
	CompanyLogo             string `json:"company_logo"`
	CompanyDateFormat       string `json:"company_date_format"`
	CompanyBankAccount      string `json:"company_bank_account"`
	CompanyBankName         string `json:"company_bank_name"`
	CompanyFavicon          string `json:"company_favicon"`
	CompanySubscriptionType string `json:"company_subscription_type"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
	LanguageId              string `json:"language_id"`
}
