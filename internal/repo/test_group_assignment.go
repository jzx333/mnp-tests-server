package repo

import (
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
	err := r.db.QueryRow(
		"INSERT INTO test_group_assignments (test_id, group_id, assigned_by, assigned_at, deadline) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		a.TestID, a.GroupID, a.AssignedBy, a.AssignedAt, a.Deadline,
	).Scan(&id)
	return id, err
}

func (r *TestGroupAssignmentRepo) GetByID(id int) (*dto.TestGroupAssignment, error) {
	a := &dto.TestGroupAssignment{}
	err := r.db.Get(a, "SELECT * FROM test_group_assignments WHERE id=$1", id)
	return a, err
}

func (r *TestGroupAssignmentRepo) GetByTest(testID int) ([]*dto.TestGroupAssignment, error) {
	var assignments []*dto.TestGroupAssignment
	err := r.db.Select(&assignments, "SELECT * FROM test_group_assignments WHERE test_id=$1 ORDER BY assigned_at", testID)
	return assignments, err
}

func (r *TestGroupAssignmentRepo) Update(a *dto.TestGroupAssignment) error {
	_, err := r.db.Exec(
		"UPDATE test_group_assignments SET group_id=$1, assigned_by=$2, assigned_at=$3, deadline=$4 WHERE id=$5",
		a.GroupID, a.AssignedBy, a.AssignedAt, a.Deadline, a.ID,
	)
	return err
}

func (r *TestGroupAssignmentRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM test_group_assignments WHERE id=$1", id)
	return err
}
