package core

import (
	"github.com/weiloon1234/go_core2/db"
	"github.com/weiloon1234/go_core2/localization"
	"github.com/weiloon1234/go_core2/redis"
	"github.com/weiloon1234/go_core2/timezone"
	"log"
)

// Config holds configuration for all components
type Config struct {
	DBConfig           db.Config
	RedisConfig        redis.Config
	LocalizationConfig localization.Config
	Timezone           string
}

// Init initializes all core components
func Init(config Config) {
	// Initialize Database
	if err := db.Init(config.DBConfig); err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}

	// Initialize Redis
	if err := redis.Init(config.RedisConfig); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	// Initialize Localization
	localization.Init(config.LocalizationConfig)

	// Set Timezone
	timezone.Init(config.Timezone)

	log.Println("Core initialized successfully")
}
