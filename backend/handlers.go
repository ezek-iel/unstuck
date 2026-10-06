package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func (app *application) ListQuestionsHandler(c *echo.Context) error {
	q, err := app.questionModel.GetAllQuestions()
	if err != nil {
		c.Logger().Error("An error occured", "error", err)
		return c.JSON(http.StatusInternalServerError, struct{ message string }{message: "An error occured on our end"})
	}

	return c.JSON(http.StatusOK, q)
}

func (app *application) GetQuestionHandler(c *echo.Context) error {
	idValue := c.Param("id")
	id, err := strconv.Atoi(idValue)

	if err != nil {
		c.Logger().Error("Invalid Param Value", "error", err)
		return c.JSON(http.StatusBadRequest, struct{ message string }{message: "Invalid id value"})
	}

	question, err := app.questionModel.GetQuestionById(id)
	if err != nil {
		return reportServerError("Error when fetching question by id", err, c)
	}

	return c.JSON(http.StatusOK, question)
}

func (app *application) ListQuestionCommentsHandler(c *echo.Context) error {
	idValue := c.Param("id")
	id, err := strconv.Atoi(idValue)

	if err != nil {
		c.Logger().Error("Invalid Param Value", "error", err)
		return c.JSON(http.StatusBadRequest, struct{ message string }{message: "Invalid id value"})
	}

	comment, err := app.commentModel.GetAllComments(id)

	if err != nil {
		return reportServerError("An error occured when fetching comments for a particular question", err, c)
	}

	return c.JSON(http.StatusOK, comment)
}

func (app *application) NewQuestionHandler(c *echo.Context) error {
	var q NewQuestion
	err := c.Bind(&q)
	if err != nil {
		return reportClientError(map[string]string{"message": err.Error()}, c)
	}

	newQuestion, validationErrors := q.ConvertToQuestion()
	if len(validationErrors) != 0 {
		return reportClientError(validationErrors, c)
	}

	err = app.questionModel.InsertNewQuestion(newQuestion)
	if err != nil {
		return reportServerError("Error occured when inserting question value", err, c)
	}

	return c.JSON(http.StatusOK, struct{ message string }{message: "Operation Completed"})
}

func (app *application) NewQuestionCommentHandler(c *echo.Context) error {
	idValue := c.Param("id")
	id, err := strconv.Atoi(idValue)

	if err != nil {
		return reportClientError(map[string]string{"message": "invalid client value", "err": err.Error()}, c)
	}

	newComment := NewComment{Question: id}
	if err := c.Bind(&newComment); err != nil {
		return reportClientError(map[string]string{"message": "error occured while receiving request", "error": err.Error()}, c)
	}

	comment, validationError := newComment.ConvertToComment()
	if len(validationError) > 0 {
		return reportClientError(validationError, c)
	}

	if err = app.commentModel.InsertNewComment(id, comment); err != nil {
		return reportServerError("an error occured when inserting comment for a question", err, c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Operation completed successfully"})
}
