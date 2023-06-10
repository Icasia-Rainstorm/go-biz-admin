package main

import (
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()
	_ = r.Run(":8080")
}
