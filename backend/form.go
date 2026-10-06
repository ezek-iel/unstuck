package main

type NewQuestion struct {
	Title   string `json:"title"`
	Details string `json:"details"`
	Tags    string `json:"tags"`
}

func (n *NewQuestion) ConvertToQuestion() (Question, map[string]string) {
	v := NewValidator()

	v.Check("title", "title length should be more than 10 characters", len(n.Title) > 10)
	v.Check("title", "title length should be less than 200 characters", len(n.Title) < 200)
	v.Check("details", "details should be more than 10 characters", len(n.Details) > 10)
	v.Check("details", "details should be less than 1000 characters", len(n.Details) < 1_000)
	return Question{Title: n.Title, Details: n.Details, Tags: n.Tags}, v.errors
}

type NewComment struct {
	Question int `json:"question"`
	Details  string `json:"details"`
}

func (c *NewComment) ConvertToComment() (Comment, map[string]string){
	v := NewValidator()

	v.Check("details", "details should be more than 10 characters",len(c.Details) > 10)
	v.Check("details", "details should be less than 1000 characters", len(c.Details) < 1_000)

	return Comment{Question: c.Question, Details: c.Details}, v.errors
}
