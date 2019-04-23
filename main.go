package main

import (
	"flag"
	"log"
	"start/config"
	"start/server"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var environmentVar string

func main() {
	environment := flag.String("ENV", "development", "descriptions") // -ENV is option for command line
	isTesting := flag.Bool("TEST", false, "run with db test")        // -TEST
	flag.Parse()
	config.ReadConfigFile(*environment)
	// envConfig := config.GetConfig()
	environmentVar := environment
	log.Println("Env", *environmentVar)
	log.Println("Is Testing", *isTesting)

	// ============ Server init
	server.Init()
}
