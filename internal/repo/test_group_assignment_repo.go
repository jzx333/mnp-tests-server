package repo

import (
	"github.com/google/uuid"
	"time"

	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type TestGroupAssignmentRepo struct {
	db *sqlx.DB
}

func NewTestGroupAssignmentRepo(db *sqlx.DB) *TestGroupAssignmentRepo {
	return &TestGroupAssignmentRepo{db: db}
}

func (r *TestGroupAssignmentRepo) Create(a *dto.TestGroupAssignment) (int, error) {
	var id int
	if a.AssignedAt.IsZero() {
		a.AssignedAt = time.Now()
	}
	err := r.db.QueryRowx(
		`INSERT INTO test_group_assignments (test_id, group_id, assigned_by, assigned_at, deadline)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		a.TestID, a.GroupID, a.AssignedBy, a.AssignedAt, a.Deadline,
	).Scan(&id)
	return id, err
}

func (r *TestGroupAssignmentRepo) GetByID(id int) (*dto.TestGroupAssignment, error) {
	var a dto.TestGroupAssignment
	err := r.db.Get(&a, `SELECT * FROM test_group_assignments WHERE id=$1`, id)
	return &a, err
}

func (r *UserRepo) GetByCuratorID(curatorID uuid.UUID) ([]dto.User, error) {
	var users []dto.User
	err := r.db.Select(&users, `SELECT * FROM users WHERE curator_id=$1`, curatorID)
	return users, err
}

func (r *TestGroupAssignmentRepo) Update(a *dto.TestGroupAssignment) error {
	_, err := r.db.Exec(
		`UPDATE test_group_assignments 
		 SET test_id=$1, group_id=$2, assigned_by=$3, assigned_at=$4, deadline=$5 
		 WHERE id=$6`,
		a.TestID, a.GroupID, a.AssignedBy, a.AssignedAt, a.Deadline, a.ID,
	)
	return err
}

func (r *TestGroupAssignmentRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM test_group_assignments WHERE id=$1`, id)
	return err
}

func (r *TestGroupAssignmentRepo) GetAll() ([]dto.TestGroupAssignment, error) {
	var assignments []dto.TestGroupAssignment
	err := r.db.Select(&assignments, `SELECT * FROM test_group_assignments`)
	return assignments, err
}
