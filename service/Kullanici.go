package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func KullaniciAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Kullanici{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Kullanici (Ad, Soyad, Eposta, Sifre, Rol) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(mdl.Ad, mdl.Soyad, mdl.Eposta, mdl.Sifre, mdl.Rol)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func KullaniciList(c echo.Context) error {
	Eposta_sql := ""
	if c.QueryParam("Eposta") != "" {
		Eposta := c.QueryParam("Eposta")
		Sifre := c.QueryParam("Sifre")
		Eposta_sql += " AND Eposta = '" + Eposta + "' "
		Eposta_sql += " AND Sifre = '" + Sifre + "' "
	}
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT KullaniciID, Ad, Soyad, Eposta, Sifre, Rol FROM Kullanici WHERE 1 = 1" + Eposta_sql)
	defer rows.Close()
	mdl := []model.Kullanici{}
	for rows.Next() {
		item := model.Kullanici{}
		rows.Scan(&item.KullaniciID, &item.Ad, &item.Soyad, &item.Eposta, &item.Sifre, &item.Rol)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func KullaniciGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Kullanici{}
	statement, _ := db.Prepare("SELECT KullaniciID, Ad, Soyad, Eposta, Sifre, Rol FROM Kullanici WHERE KullaniciID = ?")
	err = statement.QueryRow(id).Scan(&mdl.KullaniciID, &mdl.Ad, &mdl.Soyad, &mdl.Eposta, &mdl.Sifre, &mdl.Rol)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func KullaniciDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Kullanici WHERE KullaniciID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func KullaniciSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Kullanici{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Kullanici SET Ad = ?, Soyad = ?, Eposta = ?, Sifre = ?, Rol = ? WHERE KullaniciID = ?")
	statement.Exec(mdl.Ad, mdl.Soyad, mdl.Eposta, mdl.Sifre, mdl.Rol, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
