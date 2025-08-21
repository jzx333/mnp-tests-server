package dto

import (
	"github.com/google/uuid"
	"time"
)

type TestAttempt struct {
	ID           int        `db:"id" json:"id"`
	TestID       int        `db:"test_id" json:"test_id"`
	AssignmentID int        `db:"assignment_id" json:"assignment_id"`
	UserID       uuid.UUID  `db:"user_id" json:"user_id"`
	StartedAt    time.Time  `db:"started_at" json:"started_at"`
	FinishedAt   *time.Time `db:"finished_at" json:"finished_at"`
	Score        *int       `db:"score" json:"score"`
	MaxScore     *int       `db:"max_score" json:"max_score"`
	Status       string     `db:"status" json:"status"` // enum attempt_status
}
