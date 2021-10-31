package model

import (
	"database/sql"
	"eticaretapi/config"

	_ "github.com/mattn/go-sqlite3"
)

type Urun struct {
	UrunID     int     `json:"UrunID"`
	Ad         string  `json:"Ad"`
	Fiyat      float32 `json:"Fiyat"`
	Resim      string  `json:"Resim"`
	Aciklama   string  `json:"Aciklama"`
	KategoriID int     `json:"KategoriID"`
}

func UrunCreateTable() {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Urun
        (
            UrunID INTEGER PRIMARY KEY,
            Ad TEXT,
            Fiyat TEXT,
            Resim TEXT,
            Aciklama TEXT,
            KategoriID TEXT
        )
    `)
	statement.Exec()
	defer statement.Close()
	// statement2, _ := db.Prepare(`
	// 	INSERT INTO
	// 		Urun (Ad, Fiyat, Resim, Aciklama, KategoriID)
	// 	VALUES
	// 	(
	// 		'iPhone 13',
	// 		'19999.99',
	// 		'https://picsum.photos/300/200?random=1635676170217',
	// 		'iPhone 13',
	// 		'1'
	// 	),
	// 	(
	// 		'Parfüm',
	// 		'19.99',
	// 		'https://picsum.photos/300/200?random=1635676189133',
	// 		'Parfüm',
	// 		'2'
	// 	)
	// `)
	// statement2.Exec()
	// defer statement2.Close()
}
