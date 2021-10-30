package model

import (
    "database/sql"
    "eticaretapi/config"

    _ "github.com/mattn/go-sqlite3"
)

type Urunler struct {
    UrunlerID int    `json:"UrunlerID"`
    Ad string `json:"Ad"`
    Fiyat string `json:"Fiyat"`
    Resim string `json:"Resim"`
    Aciklama string `json:"Aciklama"`
    KategoriID string `json:"KategoriID"`
}

func UrunlerCreateTable() {
    db, _ := sql.Open("sqlite3", config.DB_NAME)
    defer db.Close()
    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Urunler
        (
            UrunlerID INTEGER PRIMARY KEY,
            Ad TEXT,
            Fiyat TEXT,
            Resim TEXT,
            Aciklama TEXT,
            KategoriID TEXT
        )
    `)
    statement.Exec()
    defer statement.Close()
}
