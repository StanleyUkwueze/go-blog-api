package main

import (
	loadinitializers "apicrud/LoadInitializers"
	"apicrud/models"
)

func init() {

	loadinitializers.LoadEnvVariables()
	loadinitializers.ConnectToDb()
}

func main() {
	loadinitializers.DB.AutoMigrate(&models.Post{})
}
