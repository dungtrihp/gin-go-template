package server

import (
	"log"
	"start/config"
	"start/datasources"
)

func Init() {
	log.Println("server init")

	// ============ Use mysql
	// datasources.MySqlConnect(*isTesting)
	// server.MigrateDb(*isTesting)
	datasources.MySqlConnect(false)
	MigrateDb(false)
	// ============ Use redis
	// datasources.MyRedisConnect()
	// myminio.Init()

	config := config.GetConfig()
	r := NewRouter()
	log.Println("Server is running ...", "listen", config.GetString("backend_port"))
	r.Run(config.GetString("backend_port"))

	// ============
	// MigrateDb()
}
