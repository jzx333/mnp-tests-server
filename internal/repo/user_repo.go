package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create
func (r *UserRepo) Create(u *dto.User) (uuid.UUID, error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	_, err := r.db.NamedExec(`
		INSERT INTO users (id, role_id, curator_id, unit_id, department_id, full_name, practice_start, practice_end)
		VALUES (:id, :role_id, :curator_id, :unit_id, :department_id, :full_name, :practice_start, :practice_end)
	`, u)
	return u.ID, err
}

// GetByID
func (r *UserRepo) GetByID(id uuid.UUID) (*dto.User, error) {
	var u dto.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE id=$1`, id)
	return &u, err
}

// Update
func (r *UserRepo) Update(u *dto.User) error {
	_, err := r.db.NamedExec(`
		UPDATE users
		SET role_id=:role_id, curator_id=:curator_id, unit_id=:unit_id, department_id=:department_id, 
		    full_name=:full_name, practice_start=:practice_start, practice_end=:practice_end
		WHERE id=:id
	`, u)
	return err
}

// Delete
func (r *UserRepo) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
