package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type UserGroupRepo struct {
	db *sqlx.DB
}

func NewUserGroupRepo(db *sqlx.DB) *UserGroupRepo {
	return &UserGroupRepo{db: db}
}

func (r *UserGroupRepo) Create(g *dto.UserGroup) (int, error) {
	var id int
	err := r.db.QueryRowx(`INSERT INTO user_groups (name) VALUES ($1) RETURNING id`, g.Name).Scan(&id)
	return id, err
}

func (r *UserGroupRepo) GetByID(id int) (*dto.UserGroup, error) {
	var g dto.UserGroup
	err := r.db.Get(&g, `SELECT * FROM user_groups WHERE id=$1`, id)
	return &g, err
}

func (r *UserGroupRepo) GetAll() ([]dto.UserGroup, error) {
	var groups []dto.UserGroup
	err := r.db.Select(&groups, `SELECT * FROM user_groups`)
	return groups, err
}

func (r *UserGroupRepo) Update(g *dto.UserGroup) error {
	_, err := r.db.Exec(`UPDATE user_groups SET name=$1 WHERE id=$2`, g.Name, g.ID)
	return err
}

func (r *UserGroupRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM user_groups WHERE id=$1`, id)
	return err
}
