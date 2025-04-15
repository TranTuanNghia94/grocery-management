package main

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    seedUsers(db)
    seedGroceries(db)
}

func seedUsers(db *sql.DB) {
    users := []struct {
        Username string
        Password string
    }{
        {"user1", "password1"},
        {"user2", "password2"},
    }

    for _, user := range users {
        _, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
        if err != nil {
            log.Printf("Error seeding user %s: %v", user.Username, err)
        }
    }
}

func seedGroceries(db *sql.DB) {
    groceries := []struct {
        Name     string
        Quantity int
        Price    float64
    }{
        {"Apples", 10, 1.50},
        {"Bananas", 20, 0.75},
    }

    for _, grocery := range groceries {
        _, err := db.Exec("INSERT INTO groceries (name, quantity, price) VALUES ($1, $2, $3)", grocery.Name, grocery.Quantity, grocery.Price)
        if err != nil {
            log.Printf("Error seeding grocery %s: %v", grocery.Name, err)
        }
    }
}