package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UrunlerAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Urunler{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Urunler (Ad, Fiyat, Resim, Aciklama, KategoriID) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(mdl.Ad, mdl.Fiyat, mdl.Resim, mdl.Aciklama, mdl.KategoriID)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func UrunlerList(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT UrunlerID, Ad, Fiyat, Resim, Aciklama, KategoriID FROM Urunler")
	defer rows.Close()
	mdl := []model.Urunler{}
	for rows.Next() {
		item := model.Urunler{}
		rows.Scan(&item.UrunlerID, &item.Ad, &item.Fiyat, &item.Resim, &item.Aciklama, &item.KategoriID)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func UrunlerGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Urunler{}
	statement, _ := db.Prepare("SELECT UrunlerID, Ad, Fiyat, Resim, Aciklama, KategoriID FROM Urunler WHERE UrunlerID = ?")
	err = statement.QueryRow(id).Scan(&mdl.UrunlerID, &mdl.Ad, &mdl.Fiyat, &mdl.Resim, &mdl.Aciklama, &mdl.KategoriID)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func UrunlerDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Urunler WHERE UrunlerID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func UrunlerSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Urunler{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Urunler SET Ad = ?, Fiyat = ?, Resim = ?, Aciklama = ?, KategoriID = ? WHERE UrunlerID = ?")
	statement.Exec(mdl.Ad, mdl.Fiyat, mdl.Resim, mdl.Aciklama, mdl.KategoriID, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
