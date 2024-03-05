package initializers

import (
	"ecom/models"
	"log"
)
//======================================= creating tables on DB 
func TableCreate() {
	err := DB.AutoMigrate(&models.User{}, &models.Admin{}, &models.Product{}, &models.Otp{},&models.Category{},
		&models.Rating{},&models.Review{},&models.Address{},&models.Cart{})
	if err != nil {
		log.Fatal("Failed to Automigrate", err)
	}
}
