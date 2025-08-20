package dto

import "time"

type QuestionPool struct {
	ID               int       `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	Description      string    `db:"description" json:"description"`
	TimeLimitSeconds int       `db:"time_limit_seconds" json:"time_limit_seconds"`
	CreatedBy        string    `db:"created_by" json:"created_by"`
	OwnerID          *string   `db:"owner_id" json:"owner_id"` // nullable
	UnitID           int       `db:"unit_id" json:"unit_id"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
