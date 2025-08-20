package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type UnitRepo struct {
	db *sqlx.DB
}

func NewUnitRepo(db *sqlx.DB) *UnitRepo {
	return &UnitRepo{db: db}
}

func (r *UnitRepo) Create(unit *dto.Unit) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO units (name) VALUES ($1) RETURNING id", unit.Name).Scan(&id)
	return id, err
}

func (r *UnitRepo) GetByID(id int) (*dto.Unit, error) {
	unit := &dto.Unit{}
	err := r.db.Get(unit, "SELECT id, name FROM units WHERE id=$1", id)
	return unit, err
}

func (r *UnitRepo) GetAll() ([]*dto.Unit, error) {
	var units []*dto.Unit
	err := r.db.Select(&units, "SELECT id, name FROM units ORDER BY id")
	return units, err
}

func (r *UnitRepo) Update(unit *dto.Unit) error {
	_, err := r.db.Exec("UPDATE units SET name=$1 WHERE id=$2", unit.Name, unit.ID)
	return err
}

func (r *UnitRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM units WHERE id=$1", id)
	return err
}
