package server

import (
	"log"
	"start/config"
	"start/datasources/myminio"
)

func Init() {
	log.Println("server init")

	myminio.Init()

	config := config.GetConfig()
	r := NewRouter()
	log.Println("Server is running ...", "listen", config.GetString("backend_port"))
	r.Run(config.GetString("backend_port"))

	// ============
	// MigrateDb()
}
