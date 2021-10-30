package model

import (
    "database/sql"
    "eticaretapi/config"

    _ "github.com/mattn/go-sqlite3"
)

type Kategoriler struct {
    KategorilerID int    `json:"KategorilerID"`
    ad string `json:"ad"`
    slug string `json:"slug"`
}

func KategorilerCreateTable() {
    db, _ := sql.Open("sqlite3", config.DB_NAME)
    defer db.Close()
    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Kategoriler
        (
            KategorilerID INTEGER PRIMARY KEY,
            ad TEXT,
            slug TEXT
        )
    `)
    statement.Exec()
    defer statement.Close()
}
