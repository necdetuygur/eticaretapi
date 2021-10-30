package model

import (
    "database/sql"
    "eticaretapi/config"

    _ "github.com/mattn/go-sqlite3"
)

type Todo struct {
    TodoID int    `json:"TodoID"`
    Baslik string `json:"Baslik"`
    Icerik string `json:"Icerik"`
}

func TodoCreateTable() {
    db, _ := sql.Open("sqlite3", config.DB_NAME)
    defer db.Close()
    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Todo
        (
            TodoID INTEGER PRIMARY KEY,
            Baslik TEXT,
            Icerik TEXT
        )
    `)
    statement.Exec()
    defer statement.Close()
}
