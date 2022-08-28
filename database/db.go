package database

import (
	"fmt"
	"log"
	"os"

	"github.com/auth-server-go/model"

	"github.com/jinzhu/gorm" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// ConnectDB function: Make database connection
func ConnectDB() {

	//Load environmental variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	//Define DB connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

	//connect to db URI
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&model.User{})

	fmt.Println("Successfully connected!", db)
}
