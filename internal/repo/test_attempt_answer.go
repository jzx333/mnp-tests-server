package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type TestAttemptAnswerRepo struct {
	db *sqlx.DB
}

func NewTestAttemptAnswerRepo(db *sqlx.DB) *TestAttemptAnswerRepo {
	return &TestAttemptAnswerRepo{db: db}
}

func (r *TestAttemptAnswerRepo) Create(ans *dto.TestAttemptAnswer) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO test_attempt_answers (attempt_id, question_id, option_id, text, is_correct, answered_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id",
		ans.AttemptID, ans.QuestionID, ans.OptionID, ans.Text, ans.IsCorrect, ans.AnsweredAt,
	).Scan(&id)
	return id, err
}

func (r *TestAttemptAnswerRepo) GetByAttempt(attemptID int) ([]*dto.TestAttemptAnswer, error) {
	var answers []*dto.TestAttemptAnswer
	err := r.db.Select(&answers, "SELECT * FROM test_attempt_answers WHERE attempt_id=$1 ORDER BY answered_at", attemptID)
	return answers, err
}

func (r *TestAttemptAnswerRepo) Update(ans *dto.TestAttemptAnswer) error {
	_, err := r.db.Exec(
		"UPDATE test_attempt_answers SET option_id=$1, text=$2, is_correct=$3, answered_at=$4 WHERE id=$5",
		ans.OptionID, ans.Text, ans.IsCorrect, ans.AnsweredAt, ans.ID,
	)
	return err
}

func (r *TestAttemptAnswerRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM test_attempt_answers WHERE id=$1", id)
	return err
}
