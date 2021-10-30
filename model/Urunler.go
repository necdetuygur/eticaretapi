package model

import (
    "database/sql"
    "eticaretapi/config"

    _ "github.com/mattn/go-sqlite3"
)

type Urunler struct {
    UrunlerID int    `json:"UrunlerID"`
    ad string `json:"ad"`
    fiyat string `json:"fiyat"`
    resim string `json:"resim"`
    aciklama string `json:"aciklama"`
    kategoriId string `json:"kategoriId"`
}

func UrunlerCreateTable() {
    db, _ := sql.Open("sqlite3", config.DB_NAME)
    defer db.Close()
    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Urunler
        (
            UrunlerID INTEGER PRIMARY KEY,
            ad TEXT,
            fiyat TEXT,
            resim TEXT,
            aciklama TEXT,
            kategoriId TEXT
        )
    `)
    statement.Exec()
    defer statement.Close()
}
