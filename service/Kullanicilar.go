package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func KullanicilarAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Kullanicilar{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Kullanicilar (ad, soyad, eposta, sifre, rol) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(mdl.ad, mdl.soyad, mdl.eposta, mdl.sifre, mdl.rol)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func KullanicilarList(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT KullanicilarID, ad, soyad, eposta, sifre, rol FROM Kullanicilar")
	defer rows.Close()
	mdl := []model.Kullanicilar{}
	for rows.Next() {
		item := model.Kullanicilar{}
		rows.Scan(&item.KullanicilarID, &item.ad, &item.soyad, &item.eposta, &item.sifre, &item.rol)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func KullanicilarGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Kullanicilar{}
	statement, _ := db.Prepare("SELECT KullanicilarID, ad, soyad, eposta, sifre, rol FROM Kullanicilar WHERE KullanicilarID = ?")
	err = statement.QueryRow(id).Scan(&mdl.KullanicilarID, &mdl.ad, &mdl.soyad, &mdl.eposta, &mdl.sifre, &mdl.rol)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func KullanicilarDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Kullanicilar WHERE KullanicilarID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func KullanicilarSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Kullanicilar{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Kullanicilar SET ad = ?, soyad = ?, eposta = ?, sifre = ?, rol = ? WHERE KullanicilarID = ?")
	statement.Exec(mdl.ad, mdl.soyad, mdl.eposta, mdl.sifre, mdl.rol, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
