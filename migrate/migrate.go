package main

import (
	"github.com/alpish/go-crud/intializers"
	"github.com/alpish/go-crud/models"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()
}

func main() {
	intializers.DB.AutoMigrate(&models.Post{})
}
