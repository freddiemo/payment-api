package seeders

import (
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"

	"github.com/freddiemo/payment-api/internal/customers/model"
)

type CustomerSeeder interface {
	Populate() error
}

type customerSeeder struct {
	db   *gorm.DB
	size uint8
}

func NewCustomersSeeder(db *gorm.DB) CustomerSeeder {
	return &customerSeeder{
		db: db,
	}
}

func (customerSeeder *customerSeeder) Populate() error {
	const defaultSize uint8 = 20

	customers := [defaultSize]model.Customer{}
	for i := 1; i <= int(defaultSize); i++ {
		customer := model.Customer{}
		customer.FirstName = faker.FirstName()
		customer.LastName = faker.LastName()
		customer.Email = faker.Email()
		customer.DocumentID = faker.CCNumber()
		customer.Phone = faker.E164PhoneNumber()
		var customerId uint = 1
		if i > 2 {
			customerId = uint(i) / 2
		}
		customer.MerchantID = customerId

		customers[i-1] = customer
	}

	if result := customerSeeder.db.Model(&model.Customer{}).Create(&customers); result.Error != nil {
		return result.Error
	}

	return nil
}
