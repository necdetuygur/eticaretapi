package model

import (
	"database/sql"
	"eticaretapi/config"

	_ "github.com/mattn/go-sqlite3"
)

type Kullanici struct {
	KullaniciID int    `json:"KullaniciID"`
	Ad          string `json:"Ad"`
	Soyad       string `json:"Soyad"`
	Eposta      string `json:"Eposta"`
	Sifre       string `json:"Sifre"`
	Rol         string `json:"Rol"`
}

func KullaniciCreateTable() {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS Kullanici
        (
            KullaniciID INTEGER PRIMARY KEY,
            Ad TEXT,
            Soyad TEXT,
            Eposta TEXT,
            Sifre TEXT,
            Rol TEXT
        )
    `)
	statement.Exec()
	defer statement.Close()
	// statement2, _ := db.Prepare(`
	//     INSERT INTO
	//         Kullanici (Ad, Soyad, Eposta, Sifre, Rol)
	//     VALUES
	//         (
	//             'Sistem',
	//             'Yöneticisi',
	//             'sistem@admin.com',
	//             'e9fd588b5872543d86c44e763356a495',
	//             'admin'
	//         )
	// `)
	// statement2.Exec()
	// defer statement2.Close()
}
