package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	_ "modernc.org/sqlite"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	db, err := sql.Open("sqlite", "local.db")

	if err != nil {
		e.Logger.Error("Failed to connect to database", "error", err)
	}

	e.GET("/", func(c *echo.Context) error {
		questionModel := &QuestionModel{db: db}
		q, err := questionModel.GetAllQuestions()
		if err != nil {
			c.Logger().Error("An error occured", "error", err)
			c.JSON(http.StatusInternalServerError, struct{ message string }{message: "An error occured on our end"})
			return err
		}

		return c.JSON(http.StatusOK, q)
	})

	e.GET("/comments", func(c *echo.Context) error {
		commentModel := &CommentModel{db: db}
		cm, err := commentModel.GetAllComments(2)
		if err != nil {
			c.Logger().Error("An error occured", "error", err)
			c.JSON(http.StatusInternalServerError, struct{ message string }{message: "An error occured on our end"})
			return err
		}

		return c.JSON(http.StatusOK, cm)
	})

	sc := echo.StartConfig{Address: ":1323"}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
