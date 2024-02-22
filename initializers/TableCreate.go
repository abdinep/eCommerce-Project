package initializers

import (
	"ecom/models"
	"log"
)
//======================================= creating tables on DB 
func TableCreate() {
	err := DB.AutoMigrate(&models.User{}, &models.Admin{}, &models.Product{}, &models.Otp{},&models.Category{})
	if err != nil {
		log.Fatal("Failed to Automigrate", err)
	}
}
