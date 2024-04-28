package main

import (
	_ "github.com/microsoft/go-mssqldb"
	"database/sql"
	"context"
	"log"
	"fmt"
	"time"
    "os"
    "github.com/joho/godotenv"
)

var db *sql.DB
var server = ""
var port = 1433
var user = ""
var password = ""
var database = "test"

func main() {

	var err error
    
    err = godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading environment variables file")
    }
    
    user = os.Getenv("DATABASE")
    password = os.Getenv("KEY")
    server = os.Getenv("DB_ENDPOINT")

    // Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close() // Close the connection when main exits

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected!")

	count, err := ReadBuildVersion()
	if err != nil {
		log.Fatal("Error reading build version: ", err.Error())
	}

	fmt.Printf("Read %d row(s) successfully.\n", count)
}

func ReadBuildVersion() (int, error) {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := "SELECT TOP 1 VersionDate FROM dbo.BuildVersion;"

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		var versionDate time.Time
		err := rows.Scan(&versionDate)
		if err != nil {
			return -1, err
		}

        fmt.Printf("Version Date: %s\n", versionDate.Format("2006-01-02 15:04:05"))
		count++
	}

	return count, nil
}
