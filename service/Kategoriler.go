package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func KategorilerAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Kategoriler{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Kategoriler (ad, slug) VALUES (?, ?)")
	statement.Exec(mdl.ad, mdl.slug)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func KategorilerList(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT KategorilerID, ad, slug FROM Kategoriler")
	defer rows.Close()
	mdl := []model.Kategoriler{}
	for rows.Next() {
		item := model.Kategoriler{}
		rows.Scan(&item.KategorilerID, &item.ad, &item.slug)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func KategorilerGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Kategoriler{}
	statement, _ := db.Prepare("SELECT KategorilerID, ad, slug FROM Kategoriler WHERE KategorilerID = ?")
	err = statement.QueryRow(id).Scan(&mdl.KategorilerID, &mdl.ad, &mdl.slug)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func KategorilerDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Kategoriler WHERE KategorilerID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func KategorilerSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Kategoriler{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Kategoriler SET ad = ?, slug = ? WHERE KategorilerID = ?")
	statement.Exec(mdl.ad, mdl.slug, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
