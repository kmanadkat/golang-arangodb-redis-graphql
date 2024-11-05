package db

import (
	"context"
	"fmt"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var DBClient driver.Client
var DBCollection driver.Collection

func InitializeDb() {
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASS")

	var dbUrl = "http://" + host + ":" + port
	conn, err := http.NewConnection(http.ConnectionConfig{Endpoints: []string{dbUrl}})
	if err != nil {
		fmt.Println("Error Creating Connection:", err)
		panic(err)
	}

	// Initialize the DB client
	DBClient, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(user, pass),
	})
	if err != nil {
		fmt.Println("Error creating DB client:", err)
		panic(err)
	}

	// Check Database
	db_exists, err := DBClient.DatabaseExists(context.Background(), os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Println("Error Checking Database:", err)
		panic(err)
	}
	if db_exists {
		fmt.Println("Existing Database Found")
		return
	}

	// Create Database
	_, err = DBClient.CreateDatabase(context.Background(), os.Getenv("DB_NAME"), nil)
	if err != nil {
		fmt.Println("Error Creating Database:", err)
		panic(err)
	}
}

func InitializeCollection() {
	if DBClient == nil {
		InitializeDb()
	}
	// Connect To DB
	db_conn, err := DBClient.Database(context.Background(), os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Println("Error Connecting to Database:", err)
		panic(err)
	}

	// Check Collection Exists
	col_exists, err := db_conn.CollectionExists(context.Background(), os.Getenv("DB_COLLECTION"))
	if err != nil {
		fmt.Println("Error Checking Collection:", err)
		panic(err)
	}
	if !col_exists {
		// Create Collection
		_, err = db_conn.CreateCollection(context.Background(), os.Getenv("DB_COLLECTION"), nil)
		if err != nil {
			fmt.Println("Error Creating Collection:", err)
			panic(err)
		}
	}
	DBCollection, err = db_conn.Collection(context.Background(), os.Getenv("DB_COLLECTION"))
	if err != nil {
		fmt.Println("Error Connecting to Collection:", err)
		panic(err)
	}
}

// Get Collection Instance
func GetCollection() driver.Collection {
	if DBCollection == nil {
		InitializeCollection()
	}
	return DBCollection
}
