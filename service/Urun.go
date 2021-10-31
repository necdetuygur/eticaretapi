package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UrunAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Urun{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Urun (Ad, Fiyat, Resim, Aciklama, KategoriID) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(mdl.Ad, mdl.Fiyat, mdl.Resim, mdl.Aciklama, mdl.KategoriID)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func UrunList(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT UrunID, Ad, Fiyat, Resim, Aciklama, KategoriID FROM Urun")
	defer rows.Close()
	mdl := []model.Urun{}
	for rows.Next() {
		item := model.Urun{}
		rows.Scan(&item.UrunID, &item.Ad, &item.Fiyat, &item.Resim, &item.Aciklama, &item.KategoriID)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func UrunGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Urun{}
	statement, _ := db.Prepare("SELECT UrunID, Ad, Fiyat, Resim, Aciklama, KategoriID FROM Urun WHERE UrunID = ?")
	err = statement.QueryRow(id).Scan(&mdl.UrunID, &mdl.Ad, &mdl.Fiyat, &mdl.Resim, &mdl.Aciklama, &mdl.KategoriID)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func UrunDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Urun WHERE UrunID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func UrunSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Urun{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Urun SET Ad = ?, Fiyat = ?, Resim = ?, Aciklama = ?, KategoriID = ? WHERE UrunID = ?")
	statement.Exec(mdl.Ad, mdl.Fiyat, mdl.Resim, mdl.Aciklama, mdl.KategoriID, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
