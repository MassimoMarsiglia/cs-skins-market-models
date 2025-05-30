package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CSGOAPIClient struct{}

func NewCSGOAPIClient() *CSGOAPIClient {
	return &CSGOAPIClient{}
}

// getRequest Fetches the requested URL and unmarshal the response to the passed interface
// T is the type of the response
func getRequest[T any](url string) (T, error) {
	//Fetch the requested URL
	res, err := http.Get(url)

	//set zeroValue to return in case of error
	//this is needed because we are returning a generic type
	var zeroValue T
	if err != nil {
		return zeroValue, err
	}
	defer res.Body.Close()

	//return error if the status code is not 200 with zero value
	if res.StatusCode != http.StatusOK {
		return zeroValue, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	//unmarshal to the passed interface
	var response T
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return zeroValue, err
	}

	return response, nil
}

type NameID struct {
	ID   string     `json:"id"`
	Name json.Token `json:"name"`
}

type NameIDImage struct {
	NameID
	Image string `json:"image"`
}

type Collections NameIDImage

type Rarity struct {
	NameID
	Color string `json:"color"`
}

type Crate NameIDImage

type StickerResponse []Sticker

type BaseItem struct {
	NameIDImage
	Rarity Rarity `json:"rarity"`
}
type BaseItemInstance struct {
	BaseItem
	MarketHashName string `json:"market_hash_name"`
}

type Sticker struct {
	BaseItemInstance
	Crate           []Crate `json:"crates"`
	TournamentEvent string  `json:"tournament_event"`
	TournamentTeam  string  `json:"tournament_team"`
}

type CollectionResp NameIDImage

// Fetches the stickers
func (c *CSGOAPIClient) FetchStickers() (StickerResponse, error) {
	stickers, err := getRequest[StickerResponse]("https://bymykel.github.io/CSGO-API/api/en/stickers.json")
	if err != nil {
		return StickerResponse{}, err
	}

	return stickers, nil
}

type AgentResponse []Agent

type Agent struct {
	BaseItemInstance
	Collections []CollectionResp `json:"collections"`
	Team        Team             `json:"team"`
}

func (c *CSGOAPIClient) FetchAgents() (AgentResponse, error) {
	agents, err := getRequest[AgentResponse]("https://bymykel.github.io/CSGO-API/api/en/agents.json")
	if err != nil {
		return AgentResponse{}, err
	}

	return agents, nil
}

type PatchResponse []Patch

type Patch struct {
	BaseItemInstance
}

func (c *CSGOAPIClient) FetchPatches() (PatchResponse, error) {
	patches, err := getRequest[PatchResponse]("https://bymykel.github.io/CSGO-API/api/en/patches.json")
	if err != nil {
		return PatchResponse{}, err
	}

	return patches, nil
}

type CharmResponse []Charm

type Charm struct {
	BaseItemInstance
	Collections []CollectionResp `json:"collections"`
}

func (c *CSGOAPIClient) FetchCharms() (CharmResponse, error) {
	charms, err := getRequest[CharmResponse]("https://bymykel.github.io/CSGO-API/api/en/keychains.json")
	if err != nil {
		return CharmResponse{}, err
	}

	return charms, nil
}

type CaseResponse []Cases

type Cases struct {
	BaseItem
	Type         string             `json:"type"`
	Contains     []BaseItemInstance `json:"contains"`
	ContainsRare []BaseItemInstance `json:"contains_rare"`
}

func (c *CSGOAPIClient) FetchCases() (CaseResponse, error) {
	cases, err := getRequest[CaseResponse]("https://bymykel.github.io/CSGO-API/api/en/crates.json")
	if err != nil {
		return CaseResponse{}, err
	}

	return cases, nil
}

type SkinResponse []Skin

type Weapon struct {
	NameID
	WeaponId uint16 `json:"weapon_id"`
}

type Pattern NameID

type Category NameID

type Team NameID

type Wear NameID

type Skin struct {
	BaseItem
	MinFloat    float64          `json:"min_float"`
	MaxFloat    float64          `json:"max_float"`
	Stattrak    bool             `json:"stattrak"`
	Souvenir    bool             `json:"souvenir"`
	PaintIndex  json.Number      `json:"paint_index"`
	Collections []CollectionResp `json:"collections"`
	Crates      []Crate          `json:"crates"`
	Weapon      Weapon           `json:"weapon"`
	Category    Category         `json:"category"`
	Team        Team             `json:"team"`
	Wears       []Wear           `json:"wears"`
	Pattern     Pattern          `json:"pattern"`
}

func (c *CSGOAPIClient) FetchSkins() (SkinResponse, error) {
	skins, err := getRequest[SkinResponse]("https://bymykel.github.io/CSGO-API/api/en/skins.json")
	if err != nil {
		return SkinResponse{}, err
	}

	return skins, nil
}

type SkinItemResponse []SkinItem

type SkinItem struct {
	BaseItemInstance
	SkinId     string      `json:"skin_id"`
	Weapon     Weapon      `json:"weapon"`
	Pattern    Pattern     `json:"pattern"`
	Category   Category    `json:"category"`
	MinFloat   float64     `json:"min_float"`
	MaxFloat   float64     `json:"max_float"`
	Wear       Wear        `json:"wear"`
	Stattrak   bool        `json:"stattrak"`
	Souvenir   bool        `json:"souvenir"`
	PaintIndex json.Number `json:"paint_index"`
}

func (c *CSGOAPIClient) FetchSkinItems() (SkinItemResponse, error) {
	skins, err := getRequest[SkinItemResponse]("https://bymykel.github.io/CSGO-API/api/en/skins_not_grouped.json")
	if err != nil {
		return SkinItemResponse{}, err
	}

	return skins, nil
}

type CollectionResponse []Collection

type CollectionSkins struct {
	BaseItemInstance
	PaintIndex json.Number `json:"paint_index"`
}

type Collection struct {
	NameIDImage
	Crates []Crate           `json:"crates"`
	Skins  []CollectionSkins `json:"contains"`
}

func (c *CSGOAPIClient) FetchCollections() (CollectionResponse, error) {
	collections, err := getRequest[CollectionResponse]("https://bymykel.github.io/CSGO-API/api/en/collections.json")
	if err != nil {
		return CollectionResponse{}, err
	}

	return collections, nil
}
