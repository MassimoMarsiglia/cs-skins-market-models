package database

import (
	"log"
	"os"

	"github.com/massimomarsiglia/cs-skins-market-models/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connects to the Database
func Connect() {

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}

func InitDB() {
	Connect()
	DB.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	createEnums()

	if err := DB.AutoMigrate(
		&models.Category{},
		&models.Tournament{},
		&models.TournamentTeam{},
		&models.TournamentTeamRelation{},
		&models.Rarity{},
		&models.Weapon{},
		&models.Collection{},
		&models.Wear{},
		&models.Skin{},
		&models.Sticker{},
		&models.Patch{},
		&models.Agent{},
		&models.Charm{},
		&models.Case{},
		&models.Item{},
		&models.ItemProperties{},
		&models.StickerAttributes{},
		&models.PatchAttributes{},
		&models.CharmAttributes{},
		&models.ItemAttributes{},
		&models.ItemSkin{},
	); err != nil {
		log.Fatalf("Failed to migrate item tables: %v", err)
	}

	log.Println("Database migration completed successfully")
}

func createEnums() {
	// Create wear_type enum if it doesn't exist
	DB.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'wear_type') THEN
                CREATE TYPE wear_type AS ENUM ('Factory New', 'Minimal Wear', 'Field-Tested', 'Well-Worn', 'Battle-Scarred');
            END IF;
        END
        $$;
    `)

	// Create item_type enum if it doesn't exist
	DB.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'item_type') THEN
                CREATE TYPE item_type AS ENUM ('Charm', 'Skin', 'Sticker', 'Patch', 'Agent', 'Case');
            END IF;
        END
        $$;
    `)

	DB.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'team_type') THEN
                CREATE TYPE team_type AS ENUM ('Charm', 'Skin', 'Sticker', 'Patch', 'Agent', 'Case');
            END IF;
        END
        $$;
    `)
}
