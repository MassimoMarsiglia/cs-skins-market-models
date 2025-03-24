package models

type WearType string

const (
	FactoryNew    WearType = "Factory_New"
	MinimalWear   WearType = "Minimal_Wear"
	FieldTested   WearType = "Field_Tested"
	WellWorn      WearType = "Well_Worn"
	BattleScarred WearType = "Battle_Scarred"
)

type Tournament struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type TournamentTeam struct {
	ID         string     `gorm:"primaryKey"`
	Tournament Tournament `gorm:"foreignKey:ID"`
	Team       string     `gorm:"not null"`
}

type Wear struct {
	ID       string   `gorm:"primaryKey"`
	Name     WearType `gorm:"type:wear_type;not null"`
	MinFloat float64  `gorm:"not null"` //minimum Float for Wear
	MaxFloat float64  `gorm:"not null"` //maximum Float for Wear
}

type Rarity struct {
	ID         string `gorm:"primaryKey"`
	Name       string `gorm:"unique;not null"`
	NameWeapon string `gorm:"unique;not null"`
}

type Weapon struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type Collection struct {
	ID       string    `gorm:"primaryKey"`
	Name     string    `gorm:"unique;not null"`
	Skins    []Skin    `gorm:"foreignKey:CollectionId"`
	Stickers []Sticker `gorm:"foreignKey:CollectionId"`
	Agents   []Agent   `gorm:"foreignKey:CollectionId"`
	Patches  []Patch   `gorm:"foreignKey:CollectionId"`
	Charms   []Charm   `gorm:"foreignKey:CollectionId"`
}

//Item instance
type Item struct {
	ID             string `gorm:"primaryKey"`
	MarketHashName string `gorm:"unique;not null"`
	Type           string `gorm:"type:item_type"`
}

// define concrete skin with specific wears and markethashname
type ItemSkin struct {
	MarketHashName string `gorm:"primaryKey"` //MarketHashName of the skin
	Item           Item   `gorm:"primaryKey;foreignKey:ID,constraint:OnDelete:CASCADE"`
	Wear           Wear   `gorm:"foreignKey:ID,constraint:OnDelete:CASCADE"`
	Stattrak       bool   //defines if a skin is stattrak
	Souvenir       bool   //defines if a skin is souvenir
}

// define base skin without specific wears
type Skin struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `gorm:"not null"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	WeaponId     string     `gorm:"not null"`
	Weapon       Weapon     `gorm:"foreignKey:WeaponId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
	MinFloat     float64    `gorm:"not null"`
	MaxFloat     float64    `gorm:"not null"`
	Stattrak     bool       //defines if a skin can be stattrak
	Souvenir     bool       //defines if a skin can be souvenir
	PaintSeed    uint
}

type Sticker struct {
	ID string `gorm:"primaryKey"`

	Item         Item       `gorm:"primaryKey;foreignKey:ID,constraint:OnDelete:CASCADE"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`

	//Tournament stickers
	TournamentId string
	TeamId       string
	Tournament   Tournament     `gorm:"foreignKey:TournamentId"`
	Team         TournamentTeam `gorm:"foreignKey:TeamId"`
}

type Patch struct {
	Item         Item       `gorm:"primaryKey;foreignKey:ID,constraint:OnDelete:CASCADE"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
}

type Agent struct {
	Item         Item       `gorm:"primaryKey;foreignKey:ID,constraint:OnDelete:CASCADE"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
}

type Charm struct {
	Item         Item       `gorm:"primaryKey;foreignKey:ID,constraint:OnDelete:CASCADE"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
}

type Case struct {
	Item         Item       `gorm:"primaryKey;foreignKey:ID,constraint:OnDelete:CASCADE"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
}
