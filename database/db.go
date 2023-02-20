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


func Get_user_security_information(user_email string) model.User {
	db := ConnectDB()
	// Get first matched record where email is user_email
	var user model.User
	db.Where("Username= ?", user_email).First(&user)
	return user
}

// ConnectDB function: Make database connection
func ConnectDB() *gorm.DB {

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
	Db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	// Migrate the schema
	Db.AutoMigrate(
		&model.User{})

	fmt.Println("Successfully connected!", Db)
	return Db
}
