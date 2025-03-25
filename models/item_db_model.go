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

// Item instance
type Item struct {
	ID             string `gorm:"primaryKey"`
	MarketHashName string `gorm:"unique;not null"`
	Type           string `gorm:"type:item_type"`

	Props      *ItemProperties `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`
	Attributes *ItemAttributes `gorm:"foreignKey:ItemID;references:ID;constraint:OnDelete:CASCADE"`
}

type ItemProperties struct {
	ID        string `gorm:"primaryKey"`
	ItemID    string `gorm:"not null"`
	PatchId   *string
	AgentId   *string
	CharmId   *string
	StickerId *string
	SkinId    *string
	CaseId    *string

	Item    Item      `gorm:"foreignKey:ItemID;references:ID;constraint:OnDelete:CASCADE"`
	Case    *Case     `gorm:"foreignKey:CaseId;references:ID;constraint:OnDelete:CASCADE"`
	Skin    *ItemSkin `gorm:"foreignKey:SkinId;references:ItemSkinId;constraint:OnDelete:CASCADE"`
	Patch   *Patch    `gorm:"foreignKey:PatchId;references:ID;constraint:OnDelete:CASCADE"`
	Agent   *Agent    `gorm:"foreignKey:AgentId;references:ID;constraint:OnDelete:CASCADE"`
	Charm   *Charm    `gorm:"foreignKey:CharmId;references:ID;constraint:OnDelete:CASCADE"`
	Sticker *Sticker  `gorm:"foreignKey:StickerId;references:ID;constraint:OnDelete:CASCADE"`
}

type ItemAttributes struct {
	ID     string `gorm:"primaryKey"`
	ItemID string `gorm:"not null"`
	Item   Item   `gorm:"foreignKey:ItemID;references:ID;constraint:OnDelete:CASCADE"`

	// Specific Item Attributes
	Float    *float64
	Stickers []StickerAttributes `gorm:"foreignKey:AttributesID;references:ID;constraint:OnDelete:CASCADE"`
	Patches  []PatchAttributes   `gorm:"foreignKey:AttributesID;references:ID;constraint:OnDelete:CASCADE"`
	Charms   []CharmAttributes   `gorm:"foreignKey:AttributesID;references:ID;constraint:OnDelete:CASCADE"`
}

// Attributes for Patch on Item
type PatchAttributes struct {
	ID           string `gorm:"primaryKey"`
	AttributesID string `gorm:"primaryKey"`
	PatchID      string

	Attributes ItemAttributes `gorm:"foreignKey:AttributesID;references:ID;constraint:OnDelete:CASCADE"`
	Patch      Patch          `gorm:"foreignKey:PatchID;references:ID;constraint:OnDelete:CASCADE"`
}

// Attributes for Agent on Item
type CharmAttributes struct {
	ID           string `gorm:"primaryKey"`
	AttributesID string `gorm:"primaryKey"`
	CharmID      string
	PatternId    uint16

	Attributes ItemAttributes `gorm:"foreignKey:AttributesID;references:ID;constraint:OnDelete:CASCADE"`
	Charm      Charm          `gorm:"foreignKey:CharmID;references:ID;constraint:OnDelete:CASCADE"`
}

// Attributes for Sticker on Item
type StickerAttributes struct {
	ID           string `gorm:"primaryKey"` //individual ID incase slot is not unique
	AttributesID string `gorm:"primaryKey"`
	Slot         uint8  //1-5
	StickerID    string

	Attributes ItemAttributes `gorm:"foreignKey:AttributesID;references:ID;constraint:OnDelete:CASCADE"`
	Sticker    Sticker        `gorm:"foreignKey:StickerID;references:ID;constraint:OnDelete:CASCADE"`

	Perc float64 //percentage of sticker wear on item
}

// Specific ItemSkin (meaning the actual item)
type ItemSkin struct {
	ItemSkinId     string `gorm:"primaryKey"` // Composite primary key
	MarketHashName string
	WearId         string `gorm:"not null"` // Foreign key reference
	SkinId         string `gorm:"not null"` // Foreign key reference

	Wear Wear `gorm:"foreignKey:WearId;references:ID;constraint:OnDelete:CASCADE"` // Ensures correct mapping to Wear.ID
	Skin Skin `gorm:"foreignKey:SkinId;references:ID;constraint:OnDelete:CASCADE"` // Ensures correct mapping to Skin.ID

	Stattrak bool // Defines if a skin is stattrak
	Souvenir bool // Defines if a skin is souvenir
}

// define base skin without specific wears
type Skin struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `gorm:"unique;not null"`
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
	ID           string     `gorm:"primaryKey"` //id from game files
	Name         string     `gorm:"unique;not null"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`

	//Tournament stickers
	TournamentId string         //optional
	TeamId       string         //optional
	Tournament   Tournament     `gorm:"foreignKey:TournamentId"`
	Team         TournamentTeam `gorm:"foreignKey:TeamId"`
}

type Patch struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `gorm:"unique;not null"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
}

type Agent struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `gorm:"unique;not null"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
}

type Charm struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `gorm:"unique;not null"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
	RarityId     string     `gorm:"not null"`
	Rarity       Rarity     `gorm:"foreignKey:RarityId"`
}

type Case struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `gorm:"unique;not null"`
	CollectionId string     `gorm:"not null"`
	Collection   Collection `gorm:"foreignKey:CollectionId"`
}
