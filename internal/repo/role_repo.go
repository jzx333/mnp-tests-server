package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type RoleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) Create(role *dto.Role) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO roles (name) VALUES ($1) RETURNING id", role.Name).Scan(&id)
	return id, err
}

func (r *RoleRepo) GetByID(id int) (*dto.Role, error) {
	role := &dto.Role{}
	err := r.db.Get(role, "SELECT id, name FROM roles WHERE id=$1", id)
	return role, err
}

func (r *RoleRepo) GetAll() ([]*dto.Role, error) {
	var roles []*dto.Role
	err := r.db.Select(&roles, "SELECT id, name FROM roles ORDER BY id")
	return roles, err
}

func (r *RoleRepo) Update(role *dto.Role) error {
	_, err := r.db.Exec("UPDATE roles SET name=$1 WHERE id=$2", role.Name, role.ID)
	return err
}

func (r *RoleRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM roles WHERE id=$1", id)
	return err
}
