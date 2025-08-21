package repo

import (
	"time"

	_ "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type TestRepo struct {
	db *sqlx.DB
}

func NewTestRepo(db *sqlx.DB) *TestRepo {
	return &TestRepo{db: db}
}

// Create
func (r *TestRepo) Create(t *dto.Test) (int, error) {
	var id int
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	err := r.db.QueryRowx(
		`INSERT INTO tests (title, description, created_by, created_at) 
		 VALUES ($1, $2, $3, $4) RETURNING id`,
		t.Title, t.Description, t.CreatedBy, t.CreatedAt,
	).Scan(&id)
	return id, err
}

// GetByID
func (r *TestRepo) GetByID(id int) (*dto.Test, error) {
	var t dto.Test
	err := r.db.Get(&t, `SELECT * FROM tests WHERE id=$1`, id)
	return &t, err
}

// Update
func (r *TestRepo) Update(t *dto.Test) error {
	_, err := r.db.Exec(
		`UPDATE tests SET title=$1, description=$2 WHERE id=$3`,
		t.Title, t.Description, t.ID,
	)
	return err
}

// Delete
func (r *TestRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM tests WHERE id=$1`, id)
	return err
}
