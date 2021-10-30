package model

import (
    "database/sql"
    "eticaretapi/config"

    _ "github.com/mattn/go-sqlite3"
)

type Kullanicilar struct {
    KullanicilarID int    `json:"KullanicilarID"`
    ad string `json:"ad"`
    soyad string `json:"soyad"`
    eposta string `json:"eposta"`
    sifre string `json:"sifre"`
    rol string `json:"rol"`
}

func KullanicilarCreateTable() {
    db, _ := sql.Open("sqlite3", config.DB_NAME)
    defer db.Close()
    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Kullanicilar
        (
            KullanicilarID INTEGER PRIMARY KEY,
            ad TEXT,
            soyad TEXT,
            eposta TEXT,
            sifre TEXT,
            rol TEXT
        )
    `)
    statement.Exec()
    defer statement.Close()
}
