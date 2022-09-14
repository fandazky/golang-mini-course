package controllers

import (
	"go-restful-demo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBookControllers(c echo.Context) error {
	books := []models.Book{
		{
			ID:        1,
			Title:     "Mengenal Alama Semesta",
			Author:    "Jojo Mahendra",
			Publisher: "Kalilembah (Yogyakarta)",
		},
		{
			ID:        2,
			Title:     "Azab Kubur",
			Author:    "Bambang Gentolet",
			Publisher: "Puja rambi (Malang)",
		},
		{
			ID:        3,
			Title:     "Desain Interior Rumah Minimalis",
			Author:    "Ader Herera",
			Publisher: "Gemilang Utama (Surabaya)",
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookByIdControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{
		ID:        uint(id),
		Title:     "Alarm Tubuh Tua",
		Author:    "Hanan Alqatiri",
		Publisher: "Pujangga Bintang (Yogyakarta)",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func CreateBookControllers(c echo.Context) error {
	var book models.Book
	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func UpdateBookControllers(c echo.Context) error {
	book := models.Book{
		Title:     "Perahu Kertas",
		Author:    "Dee",
		Publisher: "Bentang Pustaka (Yogyakarta)",
	}

	id, _ := strconv.Atoi(c.Param("id"))
	book.ID = uint(id)

	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func DeleteBookControllers(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "book with ID " + id + " has been deleted",
	})
}
