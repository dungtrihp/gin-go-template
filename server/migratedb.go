package server

import (
	"log"
	"start/common/models"
	"start/datasources"
)

func MigrateDb(isTesting bool) {
	db := datasources.GetDatabase()
	if isTesting == true {
		db.DropTableIfExists(
		// &models.Admins{},
		)
	}

	db.AutoMigrate(
		&models.Customer{},
	// &models.UserWallet{},
	)
	log.Println("Migrate db")
}
