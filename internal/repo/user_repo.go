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
func (r *UserRepo) Create(u *dto.User) (string, error) {
	id := uuid.New().String()
	_, err := r.db.Exec(`
		INSERT INTO users 
		(id, role_id, curator_id, unit_id, department_id, full_name, practice_start, practice_end) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		id, u.RoleID, u.CuratorID, u.UnitID, u.DepartmentID, u.FullName, u.PracticeStart, u.PracticeEnd)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetByID
func (r *UserRepo) GetByID(id string) (*dto.User, error) {
	u := &dto.User{}
	err := r.db.Get(u, `
		SELECT id, role_id, curator_id, unit_id, department_id, full_name, practice_start, practice_end
		FROM users WHERE id=$1`, id)
	return u, err
}

// GetAll
func (r *UserRepo) GetAll() ([]*dto.User, error) {
	var users []*dto.User
	err := r.db.Select(&users, `
		SELECT id, role_id, curator_id, unit_id, department_id, full_name, practice_start, practice_end
		FROM users ORDER BY full_name`)
	return users, err
}

// Update
func (r *UserRepo) Update(u *dto.User) error {
	_, err := r.db.Exec(`
		UPDATE users 
		SET role_id=$1, curator_id=$2, unit_id=$3, department_id=$4, full_name=$5, practice_start=$6, practice_end=$7
		WHERE id=$8`,
		u.RoleID, u.CuratorID, u.UnitID, u.DepartmentID, u.FullName, u.PracticeStart, u.PracticeEnd, u.ID)
	return err
}

// Delete
func (r *UserRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
