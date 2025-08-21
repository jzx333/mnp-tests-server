package dto

import (
	"github.com/google/uuid"
	"time"
)

type TestGroupAssignment struct {
	ID         int        `db:"id" json:"id"`
	TestID     int        `db:"test_id" json:"test_id"`
	GroupID    int        `db:"group_id" json:"group_id"`
	AssignedBy uuid.UUID  `db:"assigned_by" json:"assigned_by"`
	AssignedAt time.Time  `db:"assigned_at" json:"assigned_at"`
	Deadline   *time.Time `db:"deadline" json:"deadline,omitempty"`
}
