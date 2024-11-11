package database

import (
	"fmt"
	"tierlist/models"
)

func SeedUsers() error {
	if DB == nil {
		return fmt.Errorf("database connection unavailable")
	}

	var userResult []models.User
	result := DB.Find(&userResult)

	if result.RowsAffected > 0 {
		return nil
	}

	users := []models.User{
		{Name: "Test 1", Email: "test1@test.com"},
		{Name: "Test 2", Email: "test2@test.com"},
		{Name: "Test 3", Email: "test3@test.com"},
	}

	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to seed user")
		}
	}

	
	fmt.Println("Users Seeded Successfully")
	return nil
}

func SeedTierlists() error {
	if DB == nil {
		return fmt.Errorf("database connection unavailable")
	}

	var tierlistResult []models.Tierlist
	result := DB.Find(&tierlistResult)

	if result.RowsAffected > 0 {
		return nil
	}

	tierlists := []models.Tierlist{
		{UserID: 1},
		{UserID: 2},
		{UserID: 3},
	}

	for _, tierlist := range tierlists {
		if err := DB.Create(&tierlist).Error; err != nil {
			return fmt.Errorf("failed to seed tierlists")
		}
	}

	fmt.Println("Tierlists Seeded Successfully")
	return nil
}