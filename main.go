package main

import (
	"github.com/vikash-parashar/goe_com/config"
	"github.com/vikash-parashar/goe_com/db"
	"github.com/vikash-parashar/goe_com/routes"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = db.InitDB()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes()
	r.Run(":8080")
}
