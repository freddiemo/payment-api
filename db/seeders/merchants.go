package seeders

import (
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"

	"github.com/freddiemo/payment-api/internal/merchants/model"

	paymentPlatformModel "github.com/freddiemo/payment-api/internal/payment-platforms/model"
)

const defaultSize uint8 = 10

type MerchantSeeder interface {
	Populate() error
}

type merchantSeeder struct {
	db   *gorm.DB
	size uint8
}

func NewMerchantsSeeder(db *gorm.DB) MerchantSeeder {
	return &merchantSeeder{
		db: db,
	}
}

func (merchantSeeder *merchantSeeder) Populate() error {
	merchants := [defaultSize]model.Merchant{}
	for i := 0; i < int(defaultSize); i++ {
		merchant := model.Merchant{}
		merchant.Name = faker.Word()
		merchant.Email = faker.Email()
		merchant.DocumentID = faker.CCNumber()
		merchant.Phone = faker.E164PhoneNumber()

		merchants[i] = merchant
	}

	if result := merchantSeeder.db.Model(&model.Merchant{}).Create(&merchants); result.Error != nil {
		return result.Error
	}

	j := 0
	paymentPlatforms := make([]paymentPlatformModel.PaymentPlatform, defaultSize*3)
	for i := 0; i < int(defaultSize); i++ {
		paymentPlatform := paymentPlatformModel.PaymentPlatform{}

		paymentPlatform.PlatformKey = faker.CCNumber()
		paymentPlatform.Type = "paypal"
		paymentPlatform.MerchantID = uint(i + 1)
		paymentPlatforms[j] = paymentPlatform
		j++

		paymentPlatform.PlatformKey = faker.CCNumber()
		paymentPlatform.Type = "payoneer"
		paymentPlatform.MerchantID = uint(i + 1)
		paymentPlatforms[j] = paymentPlatform
		j++

		paymentPlatform.PlatformKey = faker.CCNumber()
		paymentPlatform.Type = "stripe"
		paymentPlatform.MerchantID = uint(i + 1)
		paymentPlatforms[j] = paymentPlatform
		j++
	}

	if result := merchantSeeder.db.Model(&paymentPlatformModel.PaymentPlatform{}).Create(&paymentPlatforms); result.Error != nil {
		return result.Error
	}

	return nil
}
