package server

import (
	"log"
	"start/config"
)

func Init() {
	log.Println("server init")

	config := config.GetConfig()
	r := NewRouter()
	log.Println("Server is running ...", "listen", config.GetString("backend_port"))
	r.Run(config.GetString("backend_port"))

	// ============
	// MigrateDb()
}
