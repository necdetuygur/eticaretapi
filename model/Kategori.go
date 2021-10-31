package model

import (
	"database/sql"
	"eticaretapi/config"

	_ "github.com/mattn/go-sqlite3"
)

type Kategori struct {
	KategoriID int    `json:"KategoriID"`
	Ad         string `json:"Ad"`
	Slug       string `json:"Slug"`
}

func KategoriCreateTable() {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Kategori
        (
            KategoriID INTEGER PRIMARY KEY,
            Ad TEXT,
            Slug TEXT
        )
    `)
	statement.Exec()
	defer statement.Close()
	// statement2, _ := db.Prepare(`
	// 	INSERT INTO
	// 		Kategori (Ad, Slug)
	// 	VALUES
	// 	(
	// 		'Elektronik',
	// 		'/elektronik'
	// 	),
	// 	(
	// 		'Kozmetik',
	// 		'/kozmetik'
	// 	)
	// `)
	// statement2.Exec()
	// defer statement2.Close()
}
