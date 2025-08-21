package repo

import (
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

// Assign test to group
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

// GetByID
func (r *TestGroupAssignmentRepo) GetByID(id int) (*dto.TestGroupAssignment, error) {
	var a dto.TestGroupAssignment
	err := r.db.Get(&a, `SELECT * FROM test_group_assignments WHERE id=$1`, id)
	return &a, err
}

// Update
func (r *TestGroupAssignmentRepo) Update(a *dto.TestGroupAssignment) error {
	_, err := r.db.Exec(
		`UPDATE test_group_assignments 
		 SET test_id=$1, group_id=$2, assigned_by=$3, assigned_at=$4, deadline=$5 
		 WHERE id=$6`,
		a.TestID, a.GroupID, a.AssignedBy, a.AssignedAt, a.Deadline, a.ID,
	)
	return err
}

// Delete
func (r *TestGroupAssignmentRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM test_group_assignments WHERE id=$1`, id)
	return err
}
