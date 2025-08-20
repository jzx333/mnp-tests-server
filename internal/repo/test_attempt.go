package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type TestAttemptRepo struct {
	db *sqlx.DB
}

func NewTestAttemptRepo(db *sqlx.DB) *TestAttemptRepo {
	return &TestAttemptRepo{db: db}
}

func (r *TestAttemptRepo) Create(attempt *dto.TestAttempt) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO test_attempts (test_id, assignment_id, user_id, started_at, status, score, max_score) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id",
		attempt.TestID, attempt.AssignmentID, attempt.UserID, attempt.StartedAt, attempt.Status, attempt.Score, attempt.MaxScore,
	).Scan(&id)
	return id, err
}

func (r *TestAttemptRepo) GetByID(id int) (*dto.TestAttempt, error) {
	attempt := &dto.TestAttempt{}
	err := r.db.Get(attempt, "SELECT * FROM test_attempts WHERE id=$1", id)
	return attempt, err
}

func (r *TestAttemptRepo) GetByUser(userID string) ([]*dto.TestAttempt, error) {
	var attempts []*dto.TestAttempt
	err := r.db.Select(&attempts, "SELECT * FROM test_attempts WHERE user_id=$1 ORDER BY started_at DESC", userID)
	return attempts, err
}

func (r *TestAttemptRepo) Update(attempt *dto.TestAttempt) error {
	_, err := r.db.Exec(
		"UPDATE test_attempts SET finished_at=$1, score=$2, max_score=$3, status=$4 WHERE id=$5",
		attempt.FinishedAt, attempt.Score, attempt.MaxScore, attempt.Status, attempt.ID,
	)
	return err
}

func (r *TestAttemptRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM test_attempts WHERE id=$1", id)
	return err
}
