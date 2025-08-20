package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type DepartmentRepo struct {
	db *sqlx.DB
}

func NewDepartmentRepo(db *sqlx.DB) *DepartmentRepo {
	return &DepartmentRepo{db: db}
}

func (r *DepartmentRepo) Create(dept *dto.Department) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO departments (name) VALUES ($1) RETURNING id", dept.Name).Scan(&id)
	return id, err
}

func (r *DepartmentRepo) GetByID(id int) (*dto.Department, error) {
	dept := &dto.Department{}
	err := r.db.Get(dept, "SELECT id, name FROM departments WHERE id=$1", id)
	return dept, err
}

func (r *DepartmentRepo) GetAll() ([]*dto.Department, error) {
	var depts []*dto.Department
	err := r.db.Select(&depts, "SELECT id, name FROM departments ORDER BY id")
	return depts, err
}

func (r *DepartmentRepo) Update(dept *dto.Department) error {
	_, err := r.db.Exec("UPDATE departments SET name=$1 WHERE id=$2", dept.Name, dept.ID)
	return err
}

func (r *DepartmentRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM departments WHERE id=$1", id)
	return err
}
