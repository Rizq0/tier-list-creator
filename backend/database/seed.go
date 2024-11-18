package database

import (
	"fmt"
	"tierlist/models"
)

func Seed() {
	SeedUsers()
	SeedTierlist()
}

func SeedTierlist() error {
	if DB == nil {
		return fmt.Errorf("database connection unavailable")
	}

	var tierlistResult []models.Tierlist
	result := DB.Find(&tierlistResult)

	if result.RowsAffected > 0 {
		return nil
	}

	tierlists := []models.Tierlist{
		{Name: "Tierlist 1",
		Description:  "This is the first Tierlist",
		CreatorID: 1,
		Tiers: []models.Tier{
			{Text: "S", Colour: "f4ee1d"},
			{Text: "A", Colour: "d81621"},
			{Text: "B", Colour: "00bef1"},
			{Text: "C", Colour: "37f100"},
			{Text: "D", Colour: "ffffff"},
			{Text: "F", Colour: "000000"},
		},
		Items: []models.Item{
		{Text: "Nacht der Untoten", Image: "https://static.wikia.nocookie.net/callofduty/images/2/2b/Nacht_Der_Untoten_Menu_Selection_WaW.png/revision/latest?cb=20161009103531"},
		{Text: "Verrückt", Image: "https://static.wikia.nocookie.net/callofduty/images/a/a1/Verruckt_Menu_Selection_WaW.png/revision/latest?cb=20161009103542"},
		{Text: "Shi No Numa", Image: "https://static.wikia.nocookie.net/callofduty/images/2/2f/Shi_No_Numa_Menu_Selection_WaW.png/revision/latest?cb=20161009103553"},
		{Text: "Der Riese", Image: "https://static.wikia.nocookie.net/callofduty/images/8/86/Der_Riese_Menu_Selection_WaW.png/revision/latest?cb=20161009103603"},
		{Text: "Moon", Image: "https://static.wikia.nocookie.net/callofduty/images/c/cc/Moon_Menu_Selection_BO.png/revision/latest?cb=20240710075602"},
		},
		Version: 1,},
		{Name: "Tierlist 2",
		Description:  "This is the second Tierlist",
		CreatorID: 2,
		Tiers: []models.Tier{
			{Text: "S", Colour: "f4ee1d"},
			{Text: "A", Colour: "d81621"},
			{Text: "B", Colour: "00bef1"},
			{Text: "C", Colour: "37f100"},
			{Text: "D", Colour: "ffffff"},
			{Text: "F", Colour: "000000"},
		},
		Items: []models.Item{
			{Text: "Green Hill Zone", Image: "https://example.com/green_hill_zone.png"},
			{Text: "Chemical Plant Zone", Image: "https://example.com/chemical_plant_zone.png"},
			{Text: "Casino Night Zone", Image: "https://example.com/casino_night_zone.png"},
			{Text: "Ice Cap Zone", Image: "https://example.com/ice_cap_zone.png"},
			{Text: "Sky Sanctuary Zone", Image: "https://example.com/sky_sanctuary_zone.png"},
		},
		Version: 1,},
		{Name: "Tierlist 3",
		Description:  "This is the third Tierlist",
		CreatorID: 3,
		Tiers: []models.Tier{
			{Text: "S", Colour: "f4ee1d"},
			{Text: "A", Colour: "d81621"},
			{Text: "B", Colour: "00bef1"},
			{Text: "C", Colour: "37f100"},
			{Text: "D", Colour: "ffffff"},
			{Text: "F", Colour: "000000"},
		},
		Items: []models.Item{
			{Text: "Mario", Image: "https://example.com/mario.png"},
			{Text: "Luigi", Image: "https://example.com/luigi.png"},
			{Text: "Peach", Image: "https://example.com/peach.png"},
			{Text: "Bowser", Image: "https://example.com/bowser.png"},
			{Text: "Yoshi", Image: "https://example.com/yoshi.png"},
		},
		},
	}
	for _, tierlist := range tierlists {
	if err := DB.Create(&tierlist).Error; err != nil {
		return fmt.Errorf("failed to seed tierlist")
	}
}

	fmt.Println("Tierlists Seeded Successfully")
	return nil
}

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