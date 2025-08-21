package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type OptionRepo struct {
	db *sqlx.DB
}

func NewOptionRepo(db *sqlx.DB) *OptionRepo {
	return &OptionRepo{db: db}
}

// Create
func (r *OptionRepo) Create(o *dto.Option) (int, error) {
	var id int
	err := r.db.QueryRowx(
		`INSERT INTO options (question_id, text, is_correct) VALUES ($1,$2,$3) RETURNING id`,
		o.QuestionID, o.Text, o.IsCorrect,
	).Scan(&id)
	return id, err
}

// GetByID
func (r *OptionRepo) GetByID(id int) (*dto.Option, error) {
	var o dto.Option
	err := r.db.Get(&o, `SELECT * FROM options WHERE id=$1`, id)
	return &o, err
}

// Update
func (r *OptionRepo) Update(o *dto.Option) error {
	_, err := r.db.Exec(
		`UPDATE options SET question_id=$1, text=$2, is_correct=$3 WHERE id=$4`,
		o.QuestionID, o.Text, o.IsCorrect, o.ID,
	)
	return err
}

// Delete
func (r *OptionRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM options WHERE id=$1`, id)
	return err
}
