package service

import (
	"database/sql"
	"eticaretapi/config"
	"eticaretapi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func TodoAdd(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	mdl := &model.Todo{}
	c.Bind(mdl)
	statement, _ := db.Prepare("INSERT INTO Todo (Baslik, Icerik) VALUES (?, ?)")
	statement.Exec(mdl.Baslik, mdl.Icerik)
	defer statement.Close()
	return c.JSON(http.StatusCreated, mdl)
}

func TodoList(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	rows, _ := db.Query("SELECT TodoID, Baslik, Icerik FROM Todo")
	defer rows.Close()
	mdl := []model.Todo{}
	for rows.Next() {
		item := model.Todo{}
		rows.Scan(&item.TodoID, &item.Baslik, &item.Icerik)
		mdl = append(mdl, item)
	}
	return c.JSON(http.StatusOK, mdl)
}

func TodoGet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := model.Todo{}
	statement, _ := db.Prepare("SELECT TodoID, Baslik, Icerik FROM Todo WHERE TodoID = ?")
	err = statement.QueryRow(id).Scan(&mdl.TodoID, &mdl.Baslik, &mdl.Icerik)
	defer statement.Close()
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mdl)
}

func TodoDelete(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("DELETE FROM Todo WHERE TodoID = ?")
	statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func TodoSet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", config.DB_NAME)
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	mdl := &model.Todo{}
	c.Bind(mdl)
	statement, _ := db.Prepare("UPDATE Todo SET Baslik = ?, Icerik = ? WHERE TodoID = ?")
	statement.Exec(mdl.Baslik, mdl.Icerik, id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
