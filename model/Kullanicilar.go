package model

import (
    "database/sql"
    "eticaretapi/config"

    _ "github.com/mattn/go-sqlite3"
)

type Kullanicilar struct {
    KullanicilarID int    `json:"KullanicilarID"`
    Ad string `json:"Ad"`
    Soyad string `json:"Soyad"`
    Eposta string `json:"Eposta"`
    Sifre string `json:"Sifre"`
    Rol string `json:"Rol"`
}

func KullanicilarCreateTable() {
    db, _ := sql.Open("sqlite3", config.DB_NAME)
    defer db.Close()
    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Kullanicilar
        (
            KullanicilarID INTEGER PRIMARY KEY,
            Ad TEXT,
            Soyad TEXT,
            Eposta TEXT,
            Sifre TEXT,
            Rol TEXT
        )
    `)
    statement.Exec()
    defer statement.Close()
}
