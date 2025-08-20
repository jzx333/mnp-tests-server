package dto

import "time"

type TestAttemptAnswer struct {
	ID         int       `db:"id" json:"id"`
	AttemptID  int       `db:"attempt_id" json:"attempt_id"`
	QuestionID int       `db:"question_id" json:"question_id"`
	OptionID   *int      `db:"option_id" json:"option_id"` // nullable
	Text       *string   `db:"text" json:"text"`           // для текстовых вопросов
	IsCorrect  *bool     `db:"is_correct" json:"is_correct"`
	AnsweredAt time.Time `db:"answered_at" json:"answered_at"`
}
