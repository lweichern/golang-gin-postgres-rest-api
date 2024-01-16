package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var Db *sql.DB // create db var

func ConnectDatabase(){
	err := godotenv.Load() // access .env file
	if err != nil {
		fmt.Println("Error occured on .env file...")
	}

	// read postgres details in .env
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // strconv.atoi to convert port num from string to int 
	user := os.Getenv("USERNAME")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	// set up postgresql to open
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host , port, user, dbname, password)
	db, errSql := sql.Open("postgres",psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to pg database!")
	}
}