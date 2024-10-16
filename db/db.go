package db

import (
	"github.com/charmbracelet/log"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const dbFile = "amonhen.db"

var DB *gorm.DB

func Initialize() {
	log.Info("🔌 'Connecting' to the database...")

	db, err := gorm.Open(sqlite.Open(dbFile+"?_foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	DB = db
}
