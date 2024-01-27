package main

import (
	"github.com/mishra-rishabh/json_crud_api/initializers"
	"github.com/mishra-rishabh/json_crud_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
