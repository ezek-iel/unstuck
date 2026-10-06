package main

import (
	"database/sql"
)

type Question struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Upvotes int    `json:"upvotes"`
	Tags    string `json:"tags"`
}

type QuestionModel struct {
	db *sql.DB
}

type Comment struct {
	Id       int    `json:"id"`
	Question int    `json:"question"`
	Details  string `json:"details"`
}

type CommentModel struct {
	db *sql.DB
}

func (q *QuestionModel) GetAllQuestions() ([]Question, error) {

	questions := []Question{}

	rows, err := q.db.Query("SELECT * FROM question")
	if err != nil {
		return questions, err
	}
	defer rows.Close()

	for rows.Next() {
		row := Question{}

		err := rows.Scan(&row.Id, &row.Title, &row.Details, &row.Upvotes, &row.Tags)

		if err != nil {
			return questions, nil
		}

		questions = append(questions, row)
	}

	if err := rows.Err(); err != nil {
		return questions, nil
	}

	return questions, nil
}

func (q *QuestionModel) GetQuestionById(id int) (Question, error) {
	row := q.db.QueryRow("SELECT * FROM question WHERE question.id = ?", id)
	
	var question Question
	err := row.Scan(&question.Id, &question.Title, &question.Details, &question.Upvotes, &question.Tags)

	return question, err
}

func (q *QuestionModel) InsertNewQuestion(question Question) error {
	result, err := q.db.Exec("INSERT INTO question (title, details, tags, upvotes) VALUES (? , ? , ?, 0)", question.Title, question.Details, question.Tags)
	
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
} 

func (c *CommentModel) GetAllComments(questionId int)([]Comment, error) {
	rows, err := c.db.Query("SELECT * FROM comment WHERE comment.question = ?", questionId)

	comments := []Comment{}
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		c := Comment{}
		err := rows.Scan(&c.Id, &c.Question, &c.Details)
		if err != nil {
			return comments, err
		}
		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return comments, err
	}
	
	return comments, nil
}

func (c *CommentModel) InsertNewComment(questionId int, comment Comment) error {
	result, err := c.db.Exec("INSERT INTO comment (question, details) VALUES (?, ?)", questionId, comment.Details)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}