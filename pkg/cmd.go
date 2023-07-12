package pkg

import (
	"aesGeneratorToken/pkg/repo"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var dataUser map[string]string
var dataCompany map[string]string

func GenerateToken(c *cli.Context) error {
	var email string

	if c.String("email") == "" {
		return errors.New("email is required")
	}

	email = c.String("email")

	// get data user from db
	fmt.Println("get data user from db...")
	err := getDataUserFromDb(email)
	if err != nil {
		panic(err)
	}

	fmt.Println("get data company from db...")
	err = getDataCompanyFromDb(dataUser["company_id"])
	if err != nil {
		return err
	}

	// generate token
	fmt.Println("generate token AES...")
	time.Sleep(1 * time.Second)
	fmt.Println()
	generateAes()

	return nil
}

func generateAes() {
	paperKey := os.Getenv("PAPER_KEY")
	paperIv := os.Getenv("PAPER_IV")

	iv, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(paperIv)
	if err != nil {
		panic(err)
	}

	key, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(paperKey)
	if err != nil {
		panic(err)
	}

	userByte, err := json.Marshal(dataUser)
	if err != nil {
		log.Fatalf("error marshaling user: %v", err)
		return
	}

	UserEncrypt := base64.StdEncoding.EncodeToString(AESEncrypt(userByte, key, iv))
	// fmt.Printf("companyEncrypt: %s\n", UserEncrypt)

	companyByte, err := json.Marshal(dataCompany)
	if err != nil {
		log.Fatalf("error marshaling company: %v", err)
		return
	}

	CompanyEncrypt := base64.StdEncoding.EncodeToString(AESEncrypt(companyByte, key, iv))
	// fmt.Printf("companyEncrypt: %s\n", CompanyEncrypt)

	fmt.Println("Success generate token, here we go: ")
	fmt.Println("=================================================")
	fmt.Println()
	// fmt.Println(string(result))
	fmt.Printf("companyEncrypt: %s\n", CompanyEncrypt)
	fmt.Println()
	fmt.Printf("UserEncrypt: %s\n", UserEncrypt)
	fmt.Println()
	fmt.Println("=================================================")

}

func getDataUserFromDb(email string) (err error) {
	db := OpenConnectionGorm()

	userRepo := repo.NewUserRepository(db)
	user := userRepo.GetDetailUserByEmail(email)

	if user.Uuid == "" {
		return errors.New("user not found")
	}

	dataUser = map[string]string{
		"uuid":       user.Uuid,
		"name":       user.Name,
		"email":      user.Email,
		"phone":      user.Phone,
		"photo":      user.Photo,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"company_id": user.CompanyId,
	}

	return nil
}

func getDataCompanyFromDb(companyId string) (err error) {
	db := OpenConnectionGorm()

	companyRepo := repo.NewCompanyRepository(db)
	company := companyRepo.GetDetailCompanyById(companyId)

	if company.Uuid == "" {
		return errors.New("company not found")
	}

	dataCompany = map[string]string{
		"uuid":                      company.Uuid,
		"company_name":              company.CompanyName,
		"company_email":             company.CompanyEmail,
		"company_phone":             company.CompanyPhone,
		"company_address1":          company.CompanyAddress1,
		"company_address2":          company.CompanyAddress2,
		"company_city":              company.CompanyCity,
		"company_state":             company.CompanyState,
		"company_zip_code":          company.CompanyZipCode,
		"company_country":           company.CompanyCountry,
		"company_contact":           company.CompanyContact,
		"company_vat":               company.CompanyVat,
		"company_website":           company.CompanyWebsite,
		"company_logo":              company.CompanyLogo,
		"company_date_format":       company.CompanyDateFormat,
		"company_bank_account":      company.CompanyBankAccount,
		"company_bank_name":         company.CompanyBankName,
		"company_favicon":           company.CompanyFavicon,
		"company_subscription_type": company.CompanySubscriptionType,
		"created_at":                company.CreatedAt,
	}

	return nil
}

func AESEncrypt(src []byte, key []byte, IV []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(src) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, IV)
	content := (src)
	content = PKCS5Padding(content, block.BlockSize())

	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
