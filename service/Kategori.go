package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func KategoriAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Kategori{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Kategori (Ad, Slug) VALUES (?, ?)")
	statement.Exec(mdl.Ad, mdl.Slug)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func KategoriList(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT KategoriID, Ad, Slug FROM Kategori")
	defer rows.Close()
	mdl := []model.Kategori{}
	for rows.Next() {
		item := model.Kategori{}
		rows.Scan(&item.KategoriID, &item.Ad, &item.Slug)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func KategoriGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Kategori{}
	statement, _ := db.Prepare("SELECT KategoriID, Ad, Slug FROM Kategori WHERE KategoriID = ?")
	err = statement.QueryRow(id).Scan(&mdl.KategoriID, &mdl.Ad, &mdl.Slug)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func KategoriDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Kategori WHERE KategoriID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func KategoriSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Kategori{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Kategori SET Ad = ?, Slug = ? WHERE KategoriID = ?")
	statement.Exec(mdl.Ad, mdl.Slug, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
