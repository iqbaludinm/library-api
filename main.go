package main

import (
	"github.com/iqbaludinm/library-api/app"
	"github.com/iqbaludinm/library-api/config"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init () {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	// init postgre db
	// err := config.InitDB()
	// if err != nil {
	// 	panic(err)
	// }

	// init gorm
	err := config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApp()
}