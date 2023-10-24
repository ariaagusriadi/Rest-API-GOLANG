package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"project_modul_name/models"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	DB *sql.DB
}

func InitArticle(db *sql.DB) ArticleHandler {
	return ArticleHandler{
		DB: db,
	}
}

func (h ArticleHandler) FetchArticles(c echo.Context) (err error) {
	data := make([]models.Article, 0)
	query := `SELECT id, title, body FROM article`

	rows, err := h.DB.Query(query)
	if err != nil {
		resp :=
			ErrorResponse{
				Message: err.Error(),
			}
		return c.JSON(http.StatusInternalServerError, resp)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Article
		err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Body,
		)
		if err != nil {
			resp := ErrorResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, resp)
		}
		data = append(data, item)
	}

	return c.JSON(http.StatusOK, data)
}

func (h ArticleHandler) Insert(c echo.Context) (err error) {
	var item models.Article
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}
	query := `INSERT article SET title=?, body=?`

	dbRes, err := h.DB.Exec(query, item.Title, item.Body)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}
