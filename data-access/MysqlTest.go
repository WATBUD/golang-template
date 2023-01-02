package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "github.com/go-sql-driver/mysql"
    //"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
    // err2 := godotenv.Load()
    // if (err2 != nil) {
    //     log.Fatal("Error loading .env file");
    // } else {
    //      fmt.Println("有env檔");
    // }
    os.Setenv("DBUSER", "root")
    os.Setenv("DBPASS", "")
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "test",
        AllowNativePasswords: true,
    }

    fmt.Println("os.DBUSER=" + os.Getenv("DBUSER"));
    fmt.Println("os.DBPASS=" + os.Getenv("DBPASS"));



    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}