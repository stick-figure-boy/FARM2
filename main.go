package main

import (
	"time"

	"github.com/hiroki-Fukumoto/farm2/database"
	"github.com/hiroki-Fukumoto/farm2/route"
	"github.com/joho/godotenv"
)

const location = "Asia/Tokyo"

// @title farm2
// @version 1.0
// @description  farm2
// @host localhost:8080
func main() {
	loadEnv()

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	db := database.NewDB()
	r := route.SetupRouter(db)

	r.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
