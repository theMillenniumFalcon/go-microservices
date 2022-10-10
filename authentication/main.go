package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/theMillenniumFalcon/microservices/db"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.BoolVar(&local, "local", true, "run authentication service local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
}
