package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type TestRepo struct {
	db *sqlx.DB
}

func NewTestRepo(db *sqlx.DB) *TestRepo {
	return &TestRepo{db: db}
}

func (r *TestRepo) Create(t *dto.Test) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO tests (title, description, created_by) VALUES ($1,$2,$3) RETURNING id",
		t.Title, t.Description, t.CreatedBy,
	).Scan(&id)
	return id, err
}

func (r *TestRepo) GetByID(id int) (*dto.Test, error) {
	t := &dto.Test{}
	err := r.db.Get(t, "SELECT id, title, description, created_by, created_at FROM tests WHERE id=$1", id)
	return t, err
}

func (r *TestRepo) GetAll() ([]*dto.Test, error) {
	var tests []*dto.Test
	err := r.db.Select(&tests, "SELECT id, title, description, created_by, created_at FROM tests ORDER BY id")
	return tests, err
}

func (r *TestRepo) Update(t *dto.Test) error {
	_, err := r.db.Exec("UPDATE tests SET title=$1, description=$2 WHERE id=$3",
		t.Title, t.Description, t.ID)
	return err
}

func (r *TestRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tests WHERE id=$1", id)
	return err
}
