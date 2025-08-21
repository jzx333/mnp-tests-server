package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type QuestionPoolRepo struct {
	db *sqlx.DB
}

func NewQuestionPoolRepo(db *sqlx.DB) *QuestionPoolRepo {
	return &QuestionPoolRepo{db: db}
}

// Create
func (r *QuestionPoolRepo) Create(p *dto.QuestionPool) (int, error) {
	var id int
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	err := r.db.QueryRowx(
		`INSERT INTO question_pools (name, description, time_limit_seconds, created_by, owner_id, unit_id, created_at)
		 VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		p.Name, p.Description, p.TimeLimitSeconds, p.CreatedBy, p.OwnerID, p.UnitID, p.CreatedAt,
	).Scan(&id)
	return id, err
}

// GetByID
func (r *QuestionPoolRepo) GetByID(id int) (*dto.QuestionPool, error) {
	var p dto.QuestionPool
	err := r.db.Get(&p, `SELECT * FROM question_pools WHERE id=$1`, id)
	return &p, err
}

// Update
func (r *QuestionPoolRepo) Update(p *dto.QuestionPool) error {
	_, err := r.db.Exec(
		`UPDATE question_pools SET name=$1, description=$2, time_limit_seconds=$3, owner_id=$4, unit_id=$5 WHERE id=$6`,
		p.Name, p.Description, p.TimeLimitSeconds, p.OwnerID, p.UnitID, p.ID,
	)
	return err
}

// Delete
func (r *QuestionPoolRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM question_pools WHERE id=$1`, id)
	return err
}
