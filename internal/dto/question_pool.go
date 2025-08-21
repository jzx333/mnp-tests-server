package dto

import (
	"github.com/google/uuid"
	"time"
)

type QuestionPool struct {
	ID               int        `db:"id" json:"id"`
	Name             string     `db:"name" json:"name"`
	Description      string     `db:"description" json:"description"`
	TimeLimitSeconds int        `db:"time_limit_seconds" json:"time_limit_seconds"`
	CreatedBy        uuid.UUID  `db:"created_by" json:"created_by"`
	OwnerID          *uuid.UUID `db:"owner_id" json:"owner_id,omitempty"`
	UnitID           int        `db:"unit_id" json:"unit_id"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
}
