package dto

type Option struct {
	ID         int    `db:"id" json:"id"`
	QuestionID int    `db:"question_id" json:"question_id"`
	Text       string `db:"text" json:"text"`
	IsCorrect  bool   `db:"is_correct" json:"is_correct"`
}
