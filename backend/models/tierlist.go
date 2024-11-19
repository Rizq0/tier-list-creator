package models

type Tierlist struct {
	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreatorID int `json:"creator_id"`
	Creator User `gorm:"foreignKey:CreatorID"`
	Tiers []Tier `gorm:"foreignKey:TierlistID"`
	Items []Item `gorm:"foreignKey:TierlistID"`
	Version int `gorm:"default:1" json:"version"`
}

type Tier struct {
	ID int `gorm:"primaryKey" json:"id"`
	TierlistID int `json:"tierlist_id"`
	Text string `json:"text"`
	Colour string `json:"colour"`
	Tierlist Tierlist `gorm:"foreignKey:TierlistID"`
	Items []Item `gorm:"foreignKey:TierID"`
}

type Item struct {
	ID int `gorm:"primaryKey" json:"id"`
	TierlistID int `json:"tierlist_id"`
	Text string `json:"text"`
	Image string `json:"image"`
	TierID int `json:"tier_id"`
	Tier Tier `gorm:"foreignKey:TierID"`
	Tierlist Tierlist `gorm:"foreignKey:TierlistID"`
}