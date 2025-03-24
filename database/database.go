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
	createEnums()
	err := DB.AutoMigrate(
		&models.TournamentTeam{}, 
		&models.Tournament{}, 
		&models.Wear{}, 
		&models.Rarity{}, 
		&models.Weapon{}, 
		&models.Collection{},
		&models.Item{},
		&models.Sticker{},  
		&models.ItemSkin{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}

func createEnums() {
	// Create wear_type enum if it doesn't exist
	DB.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'wear_type') THEN
                CREATE TYPE wear_type AS ENUM ('Factory_New', 'Minimal_Wear', 'Field_Tested', 'Well_Worn', 'Battle_Scarred');
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
}
