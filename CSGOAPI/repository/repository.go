package repository

import (
	"fmt"
	"log"
	"reflect"

	"github.com/massimomarsiglia/cs-skins-market-models/CSGOAPI/client"
	"github.com/massimomarsiglia/cs-skins-market-models/models"
	"gorm.io/gorm"
)

type Repository struct{}

func (r *Repository) CreateCrate(c []client.Crate, tx *gorm.DB) ([]models.Case, error) {
	var crates []models.Case
	for _, crate := range c {
		var crateModel models.Case

		// Check if the crate already exists in the database, if not create
		if err := tx.FirstOrCreate(&crateModel, models.Case{
			ID:    crate.ID,
			Name:  crate.Name.(string),
			Image: crate.Image,
		}).Error; err != nil {
			return []models.Case{}, err
		}
		crates = append(crates, crateModel)
	}
	return crates, nil
}

func (r *Repository) CreateTournament(t *string, tx *gorm.DB) (models.Tournament, error) {
	var tournament *models.Tournament

	// Check if the tournament already exists in the database, if not create
	if err := tx.FirstOrCreate(&tournament, models.Tournament{
		Name: *t,
	}).Error; err != nil {
		return models.Tournament{}, err
	}
	return *tournament, nil
}

func (r *Repository) CreateTournamentTeam(t *string, tx *gorm.DB) (models.TournamentTeam, error) {
	var team *models.TournamentTeam

	// Check if the tournament team already exists in the database, if not create
	if err := tx.FirstOrCreate(&team, models.TournamentTeam{
		Team: *t,
	}).Error; err != nil {
		return models.TournamentTeam{}, err
	}
	return *team, nil
}

func (r *Repository) CreateTournamentTeamRelation(t *models.Tournament, tx *gorm.DB) (models.TournamentTeamRelation, error) {
	var tournamentTeamRelation *models.TournamentTeamRelation

	// Check if the tournament team relation already exists in the database, if not create
	if err := tx.FirstOrCreate(&tournamentTeamRelation, models.TournamentTeamRelation{
		TournamentID:     t.ID,
		TournamentTeamID: t.ID,
	}).Error; err != nil {
		return models.TournamentTeamRelation{}, err
	}
	return *tournamentTeamRelation, nil
}

func (r *Repository) CreateRarity(rar *client.Rarity, tx *gorm.DB) (models.Rarity, error) {
	var rarity *models.Rarity

	// Check if the rarity already exists in the database, if not create
	if err := tx.FirstOrCreate(&rarity, models.Rarity{
		ID:    rar.ID,
		Name:  rar.Name.(string),
		Color: rar.Color,
	}).Error; err != nil {
		return models.Rarity{}, err
	}
	return *rarity, nil
}

func (r *Repository) CreateCollection(c []client.CollectionResp, tx *gorm.DB) ([]models.Collection, error) {
	var collections []models.Collection

	// Check if the collection already exists in the database, if not create
	if len(c) == 0 {
		return []models.Collection{}, nil
	}
	for _, c := range c {
		var collection models.Collection
		if err := tx.FirstOrCreate(&collection, models.Collection{
			ID:    c.ID,
			Name:  c.Name.(string),
			Image: c.Image,
		}).Error; err != nil {
			return []models.Collection{}, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (r *Repository) CreateWears(w []client.Wear, tx *gorm.DB) ([]models.Wear, error) {
	var wears []models.Wear
	for _, wear := range w {
		var wearModel models.Wear

		// Transform wear name to enum
		wearName, err := func() (models.WearType, error) {
			if name, ok := wear.Name.(string); ok {
				return models.WearType(name), nil
			}
			return models.WearType(""), fmt.Errorf("invalid wear name")
		}()
		if err != nil {
			// Log the error and continue with the next item
			log.Printf("Error processing wear ID %v: %v", wear.ID, err)
			continue // Skip this wear and proceed with the next one
		}

		// Create wear model
		if err := tx.FirstOrCreate(&wearModel, models.Wear{
			ID:   wear.ID,
			Name: wearName,
		}).Error; err != nil {
			return nil, err // Return nil to indicate failure
		}
		wears = append(wears, wearModel)
	}

	if len(wears) == 0 {
		return nil, fmt.Errorf("no valid wears were created") // Optionally, return an error if no valid wears were processed
	}

	return wears, nil
}
func safeGetString(token interface{}, objectID string, fieldName string) string {
	if token == nil {
		log.Printf("WARNING: %s has nil %s field", objectID, fieldName)
		return fmt.Sprintf("Unknown_%s", objectID)
	}

	// Try type assertion to string
	if str, ok := token.(string); ok {
		return str
	}

	// If type assertion fails, log the actual type and value
	tokenType := reflect.TypeOf(token)
	tokenValue := fmt.Sprintf("%v", token)
	log.Printf("WARNING: Type assertion failed on %s.%s - Expected string but got %v (%s: %s)",
		objectID, fieldName, tokenType, fieldName, tokenValue)

	// Convert other types to string
	switch v := token.(type) {
	case float64:
		return fmt.Sprintf("%.0f", v) // Convert float64 to string without decimal places
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return tokenValue
	}
}

func (r *Repository) CreatePattern(p *client.Pattern, tx *gorm.DB) (models.Pattern, error) {
	var pattern *models.Pattern

	// Check if the pattern already exists in the database, if not create
	if err := tx.FirstOrCreate(&pattern, models.Pattern{
		ID:   p.ID,
		Name: safeGetString(p.Name, p.ID, "Name"),
	}).Error; err != nil {
		fmt.Println(p)
		return models.Pattern{}, err
	}
	return *pattern, nil
}

func (r *Repository) CreateTeam(t *client.Team, tx *gorm.DB) (models.Team, error) {
	var team *models.Team

	// Check if the team already exists in the database, if not create
	if err := tx.FirstOrCreate(&team, models.Team{
		ID:   t.ID,
		Name: t.Name.(string),
	}).Error; err != nil {
		return models.Team{}, err
	}
	return *team, nil
}

func (r *Repository) CreateCategory(c *client.Category, tx *gorm.DB) (models.Category, error) {
	var category *models.Category

	// Check if the category already exists in the database, if not create
	if err := tx.FirstOrCreate(&category, models.Category{
		ID:   c.ID,
		Name: c.Name.(string),
	}).Error; err != nil {
		return models.Category{}, err
	}
	return *category, nil
}

func (r *Repository) CreateWeapon(w *client.Weapon, tx *gorm.DB) (models.Weapon, error) {
	var weapon *models.Weapon

	// Check if the weapon already exists in the database, if not create
	if err := tx.FirstOrCreate(&weapon, models.Weapon{
		ID:   w.ID,
		Name: w.Name.(string),
	}).Error; err != nil {
		return models.Weapon{}, err
	}
	return *weapon, nil
}

func (r *Repository) CreateSticker(s *client.Sticker, t *models.Tournament, tot *models.TournamentTeam, rar *models.Rarity, c []models.Case, tx *gorm.DB) (models.Sticker, error) {
	var sticker models.Sticker
	var crateId *string
	if len(c) > 0 {
		crateId = &c[0].ID
	}
	// Check if the sticker already exists in the database, if not create
	if err := tx.FirstOrCreate(&sticker, models.Sticker{
		ID:           s.ID,
		Name:         s.Name.(string),
		Image:        s.Image,
		RarityId:     rar.ID,
		TeamId:       &tot.ID,
		TournamentId: &t.ID,
		CaseID:       crateId,
	}).Error; err != nil {
		return models.Sticker{}, err
	}
	return sticker, nil
}

func (r *Repository) CreateSkin(s *client.Skin, rarID *string, wID *string, colID *string, catID *string, teamID *string, patID *string, w []models.Wear, tx *gorm.DB) (models.Skin, error) {
	paintIndex := func() uint16 {
		if value, err := s.PaintIndex.Int64(); err == nil {
			return uint16(value)
		}
		return 0 // Default value in case of error
	}()

	skinModel := models.Skin{
		ID:         s.ID,
		Name:       s.Name.(string),
		Image:      s.Image,
		RarityId:   *rarID,
		WeaponId:   *wID,
		PaintIndex: paintIndex,
		MinFloat:   s.MinFloat,
		MaxFloat:   s.MaxFloat,
		Stattrak:   s.Stattrak,
		Souvenir:   s.Souvenir,
		CategoryId: *catID,
		TeamId:     *teamID,
		PatternId:  *patID,
		Wears:      w,
	}
	if colID != nil {
		skinModel.CollectionId = colID
	}

	if err := tx.FirstOrCreate(&models.Skin{}, skinModel).Error; err != nil {
		return models.Skin{}, err
	}
	return models.Skin{}, nil
}

func (r *Repository) CreateSkinItem(s *client.SkinItem, w []models.Wear, tx *gorm.DB) (models.ItemSkin, error) {

	var skinItem models.ItemSkin

	var skin = models.ItemSkin{
		ID:             s.ID,
		MarketHashName: s.MarketHashName,
		SkinId:         s.SkinId,
		Image:          s.Image,
		Stattrak:       s.Stattrak,
		Souvenir:       s.Souvenir,
	}
	if len(w) > 0 {
		skinItem.WearId = &w[0].ID
	}

	if err := tx.FirstOrCreate(&skinItem, &skin).Error; err != nil {
		return models.ItemSkin{}, err
	}
	return models.ItemSkin{}, nil
}

func (r *Repository) CreateAgent(a *client.Agent, tx *gorm.DB) (models.Agent, error) {
	var agent *models.Agent

	// Check if the agent already exists in the database, if not create
	if err := tx.FirstOrCreate(&agent, models.Agent{
		ID:           a.ID,
		Name:         a.Name.(string),
		CollectionId: a.Collections[0].ID,
		Image:        a.Image,
		TeamId:       a.Team.ID,
		RarityId:     a.Rarity.ID,
	}).Error; err != nil {
		return models.Agent{}, err
	}
	return *agent, nil
}
