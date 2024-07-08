package utils

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "yourpassword"
    dbname   = "bookdb"
)

var db *sql.DB

func InitDB() *sql.DB {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    conn, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    db = conn

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database!")

    return db
}

func GetDB() *sql.DB {
    return db
}
