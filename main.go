package main

import (
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/config"
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/routes"
	"log"
	"os"

	_ "FP-Sanbercode-Go-48-Tubagus_Saifulloh/docs"
)

// @contact.name API Support
// @contact.url http://github.com/bagus2x
// @contact.email tubagus.sflh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db := config.ConnectDatabase(dbUsername, dbPassword, dbHost, dbPort, dbName)

	r := routes.SetupRoutes(db)

	log.Fatal(r.Run())
}
