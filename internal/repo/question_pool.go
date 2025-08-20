package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type QuestionPoolRepo struct {
	db *sqlx.DB
}

func NewQuestionPoolRepo(db *sqlx.DB) *QuestionPoolRepo {
	return &QuestionPoolRepo{db: db}
}

func (r *QuestionPoolRepo) Create(pool *dto.QuestionPool) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO question_pools (name, description, time_limit_seconds, created_by, owner_id, unit_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id",
		pool.Name, pool.Description, pool.TimeLimitSeconds, pool.CreatedBy, pool.OwnerID, pool.UnitID,
	).Scan(&id)
	return id, err
}

func (r *QuestionPoolRepo) GetByID(id int) (*dto.QuestionPool, error) {
	pool := &dto.QuestionPool{}
	err := r.db.Get(pool, "SELECT * FROM question_pools WHERE id=$1", id)
	return pool, err
}

func (r *QuestionPoolRepo) GetAll() ([]*dto.QuestionPool, error) {
	var pools []*dto.QuestionPool
	err := r.db.Select(&pools, "SELECT * FROM question_pools ORDER BY id")
	return pools, err
}

func (r *QuestionPoolRepo) Update(pool *dto.QuestionPool) error {
	_, err := r.db.Exec(
		"UPDATE question_pools SET name=$1, description=$2, time_limit_seconds=$3, owner_id=$4, unit_id=$5 WHERE id=$6",
		pool.Name, pool.Description, pool.TimeLimitSeconds, pool.OwnerID, pool.UnitID, pool.ID,
	)
	return err
}

func (r *QuestionPoolRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM question_pools WHERE id=$1", id)
	return err
}
