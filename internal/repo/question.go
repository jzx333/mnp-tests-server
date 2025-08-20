package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type QuestionRepo struct {
	db *sqlx.DB
}

func NewQuestionRepo(db *sqlx.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

func (r *QuestionRepo) Create(q *dto.Question) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO questions (pool_id, text, question_type, score, position, media_url) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id",
		q.PoolID, q.Text, q.QuestionType, q.Score, q.Position, q.MediaURL,
	).Scan(&id)
	return id, err
}

func (r *QuestionRepo) GetByID(id int) (*dto.Question, error) {
	q := &dto.Question{}
	err := r.db.Get(q, "SELECT * FROM questions WHERE id=$1", id)
	return q, err
}

func (r *QuestionRepo) GetByPool(poolID int) ([]*dto.Question, error) {
	var qs []*dto.Question
	err := r.db.Select(&qs, "SELECT * FROM questions WHERE pool_id=$1 ORDER BY position", poolID)
	return qs, err
}

func (r *QuestionRepo) Update(q *dto.Question) error {
	_, err := r.db.Exec(
		"UPDATE questions SET pool_id=$1, text=$2, question_type=$3, score=$4, position=$5, media_url=$6 WHERE id=$7",
		q.PoolID, q.Text, q.QuestionType, q.Score, q.Position, q.MediaURL, q.ID,
	)
	return err
}

func (r *QuestionRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM questions WHERE id=$1", id)
	return err
}
