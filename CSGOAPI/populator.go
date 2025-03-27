package CSGOAPI

import (
	"fmt"
	"sync"
	"time"

	"github.com/massimomarsiglia/cs-skins-market-models/CSGOAPI/client"
	"github.com/massimomarsiglia/cs-skins-market-models/CSGOAPI/repository"
	"github.com/massimomarsiglia/cs-skins-market-models/database"
	"gorm.io/gorm"
)

type Populator struct {
	c *client.CSGOAPIClient
	r *repository.Repository
}

func NewPopulator(c *client.CSGOAPIClient) *Populator {
	return &Populator{c: c}
}

func (p *Populator) PopulateDB() {
	t := time.Now()

	data, err := p.fetchData()
	if err != nil {
		panic(err)
	}

	err = p.processStickers(data.Stickers)
	if err != nil {
		panic(err)
	}

	err = p.processSkins(data.Skins)
	if err != nil {
		panic(err)
	}

	fmt.Println("Time since start: ", time.Since(t))
	time.Sleep(10 * time.Second)
}

func (p *Populator) processStickers(s client.StickerResponse) error {
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		for _, sticker := range s {
			crates, err := p.r.CreateCrate(sticker.Crate, tx)
			if err != nil {
				return err
			}

			rarity, err := p.r.CreateRarity(&sticker.Rarity, tx)
			if err != nil {
				return err
			}

			tournament, err := p.r.CreateTournament(&sticker.TournamentEvent, tx)
			if err != nil {
				return err
			}

			team, err := p.r.CreateTournamentTeam(&sticker.TournamentTeam, tx)
			if err != nil {
				return err
			}

			if _, err := p.r.CreateTournamentTeamRelation(&tournament, tx); err != nil {
				return err
			}

			if _, err := p.r.CreateSticker(&sticker, &tournament, &team, &rarity, crates, tx); err != nil {
				return err
			}

		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (p *Populator) processSkins(s client.SkinResponse) error {
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		for _, skin := range s {

			rarity, err := p.r.CreateRarity(&skin.Rarity, tx)
			if err != nil {
				return err
			}

			collections, err := p.r.CreateCollection(skin.Collections, tx)
			if err != nil {
				return err
			}

			weapon, err := p.r.CreateWeapon(&skin.Weapon, tx)
			if err != nil {
				return err
			}

			category, err := p.r.CreateCategory(&skin.Category, tx)
			if err != nil {
				return err
			}

			team, err := p.r.CreateTeam(&skin.Team, tx)
			if err != nil {
				return err
			}

			pattern, err := p.r.CreatePattern(&skin.Pattern, tx)
			if err != nil {
				return err
			}

			wears, err := p.r.CreateWears(skin.Wears, tx)
			if err != nil {
				return err
			}

			var collectionID *string
			if len(collections) > 0 {
				collectionID = &collections[0].ID
			}

			_, err = p.r.CreateSkin(&skin, &rarity.ID, &weapon.ID, collectionID, &category.ID, &team.ID, &pattern.ID, wears, tx)
			if err != nil {
				return err
			}

		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

type FetchedData struct {
	Agents      client.AgentResponse
	Patches     client.PatchResponse
	Charms      client.CharmResponse
	Cases       client.CaseResponse
	Collections client.CollectionResponse
	Skins       client.SkinResponse
	Stickers    client.StickerResponse
	SkinItems   client.SkinItemResponse
	Errors      map[string]error
}

func (p *Populator) fetchData() (*FetchedData, error) {

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	result := &FetchedData{
		Errors: make(map[string]error),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		agents, err := p.c.FetchAgents()
		mu.Lock()
		if err != nil {
			result.Errors["agents"] = err
		} else {
			result.Agents = agents
			fmt.Printf("Fetched %d agents\n", len(agents))
		}
		mu.Unlock()
	}()

	// Fetch patches
	wg.Add(1)
	go func() {
		defer wg.Done()
		patches, err := p.c.FetchPatches()
		mu.Lock()
		if err != nil {
			result.Errors["patches"] = err
		} else {
			result.Patches = patches
			fmt.Printf("Fetched %d patches\n", len(patches))
		}
		mu.Unlock()
	}()

	// Fetch charms
	wg.Add(1)
	go func() {
		defer wg.Done()
		charms, err := p.c.FetchCharms()
		mu.Lock()
		if err != nil {
			result.Errors["charms"] = err
		} else {
			result.Charms = charms
			fmt.Printf("Fetched %d charms\n", len(charms))
		}
		mu.Unlock()
	}()

	// Fetch cases
	wg.Add(1)
	go func() {
		defer wg.Done()
		cases, err := p.c.FetchCases()
		mu.Lock()
		if err != nil {
			result.Errors["cases"] = err
		} else {
			result.Cases = cases
			fmt.Printf("Fetched %d cases\n", len(cases))
		}
		mu.Unlock()
	}()

	// Fetch collections
	wg.Add(1)
	go func() {
		defer wg.Done()
		collections, err := p.c.FetchCollections()
		mu.Lock()
		if err != nil {
			result.Errors["collections"] = err
		} else {
			result.Collections = collections
			fmt.Printf("Fetched %d collections\n", len(collections))
		}
		mu.Unlock()
	}()

	// Fetch skins
	wg.Add(1)
	go func() {
		defer wg.Done()
		skins, err := p.c.FetchSkins()
		mu.Lock()
		if err != nil {
			result.Errors["skins"] = err
		} else {
			result.Skins = skins
			fmt.Printf("Fetched %d skins\n", len(skins))
		}
		mu.Unlock()
	}()

	// Fetch stickers
	wg.Add(1)
	go func() {
		defer wg.Done()
		stickers, err := p.c.FetchStickers()
		mu.Lock()
		if err != nil {
			result.Errors["stickers"] = err
		} else {
			result.Stickers = stickers
			fmt.Printf("Fetched %d stickers\n", len(stickers))
		}
		mu.Unlock()
	}()

	// Fetch skin items
	wg.Add(1)
	go func() {
		defer wg.Done()
		skinItems, err := p.c.FetchSkinItems()
		mu.Lock()
		if err != nil {
			result.Errors["skinItems"] = err
		} else {
			result.SkinItems = skinItems
			fmt.Printf("Fetched %d skin items\n", len(skinItems))
		}
		mu.Unlock()
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if there were any errors
	if len(result.Errors) > 0 {
		return result, fmt.Errorf("some fetch operations failed")
	}

	return result, nil
}
