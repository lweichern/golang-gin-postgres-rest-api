package database

import (
	"example/http-server/models"
	"fmt"

	// "os"
	// "strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var Db *gorm.DB // create db var

func ConnectDatabase() {
	err := godotenv.Load() // access .env file
	if err != nil {
		panic("Error occured on .env file...")
	}

	// read postgres details in .env
	// host := os.Getenv("HOST")
	// port, _ := strconv.Atoi(os.Getenv("PORT")) // strconv.atoi to convert port num from string to int
	// user := os.Getenv("USERNAME")
	// dbname := os.Getenv("DB_NAME")
	// password := os.Getenv("PASSWORD")

	// set up postgresql to open
	// psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	psqlSetup := "postgresql://postgres:cD-eCfe*AGA4GCabcg*Bc*b5BC61D*Ga@viaduct.proxy.rlwy.net:48223/railway"

	// connect to postgres db
	db, errSql := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&models.Author{}, &models.Book{}, &models.User{})

	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to pg database!")
	}
}
