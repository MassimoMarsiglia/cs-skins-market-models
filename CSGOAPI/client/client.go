package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CSGOAPIClient struct{}

// getRequest fetches the requested URL and unmarshal the response to the passed interface
// T is the type of the response
func getRequest[T any](url string) (T, error) {
	//fetch the requested URL
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

type StickerResponse []Stickers

type Stickers struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	Crates []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Image string `json:"image"`
	}
	TournamentEvent string `json:"tournament_event"`
	TournamentTeam  string `json:"tournament_team"`
	Image           string `json:"image"`
}

// fetches the stickers
func (c *CSGOAPIClient) fetchStickers() (StickerResponse, error) {
	stickers, err := getRequest[StickerResponse]("https://bymykel.github.io/CSGO-API/api/en/stickers.json")
	if err != nil {
		return StickerResponse{}, err
	}

	return stickers, nil
}

type AgentResponse []Agents

type Agents struct {}

func (c *CSGOAPIClient) fetchAgents() (AgentResponse, error) {
	agents, err := getRequest[AgentResponse]("https://bymykel.github.io/CSGO-API/api/en/agents.json")
	if err != nil {
		return AgentResponse{}, err
	}

	return agents, nil
}

type PatchResponse []Patches

type Patches struct {}

func (c *CSGOAPIClient) fetchPatches() (PatchResponse, error) {
	patches, err := getRequest[PatchResponse]("https://bymykel.github.io/CSGO-API/api/en/patches.json")
	if err != nil {
		return PatchResponse{}, err
	}

	return patches, nil
}

type CharmResponse []Charms

type Charms struct {}

func (c *CSGOAPIClient) fetchCharms() (CharmResponse, error) {
	charms, err := getRequest[CharmResponse]("https://bymykel.github.io/CSGO-API/api/en/keychains.json")
	if err != nil {
		return CharmResponse{}, err
	}

	return charms, nil
}

type CaseResponse []Cases

type Cases struct {}


func (c *CSGOAPIClient) fetchCases() (CaseResponse, error) {
	cases, err := getRequest[CaseResponse]("https://bymykel.github.io/CSGO-API/api/en/crates.json")
	if err != nil {
		return CaseResponse{}, err
	}

	return cases, nil
}

type SkinResponse []Skins

type Skins struct {}

func (c *CSGOAPIClient) fetchSkins() (SkinResponse, error) {
	skins, err := getRequest[SkinResponse]("https://bymykel.github.io/CSGO-API/api/en/skins.json")
	if err != nil {
		return SkinResponse{}, err
	}

	return skins, nil
}

type SkinItemResponse []SkinItems

type SkinItems struct {}

func (c *CSGOAPIClient) fetchSkinItems() (SkinResponse, error) {
	skins, err := getRequest[SkinResponse]("https://bymykel.github.io/CSGO-API/api/en/skins_not_grouped.json")
	if err != nil {
		return SkinResponse{}, err
	}

	return skins, nil
}

type CollectionResponse []Collections

type Collections struct {}

func (c *CSGOAPIClient) fetchCollections() (CollectionResponse, error) {
	collections, err := getRequest[CollectionResponse]("https://bymykel.github.io/CSGO-API/api/en/collections.json")
	if err != nil {
		return CollectionResponse{}, err
	}

	return collections, nil
}