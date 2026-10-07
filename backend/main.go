package main

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	_ "modernc.org/sqlite"
)

type application struct {
	echo *echo.Echo
	questionModel *QuestionModel
	commentModel  *CommentModel
}

func newApplication(db *sql.DB, e *echo.Echo) *application {
	return &application{
		echo: e,
		questionModel: &QuestionModel{db: db},
		commentModel:  &CommentModel{db: db},
	}
}

func (app *application) registerRoutes() {
	app.echo.GET("/questions", app.ListQuestionsHandler)
	app.echo.GET("/questions/:id", app.GetQuestionHandler)
	app.echo.GET("/questions/:id/comments", app.ListQuestionCommentsHandler)
	app.echo.POST("/new", app.NewQuestionHandler)
	app.echo.POST("/questions/:id/new", app.NewQuestionCommentHandler)
}

func main() {
	// Echo instance
	e := echo.New()
	
	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS("http://localhost:5173"))
	
	db, err := sql.Open("sqlite", "local.db")
	
	if err != nil {
		e.Logger.Error("Failed to connect to database", "error", err)
	}

	app := newApplication(db, e)
	app.registerRoutes()


	sc := echo.StartConfig{Address: ":1323"}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
