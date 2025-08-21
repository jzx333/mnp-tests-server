package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type TestPoolRepo struct {
	db *sqlx.DB
}

func NewTestPoolRepo(db *sqlx.DB) *TestPoolRepo {
	return &TestPoolRepo{db: db}
}

// Add pool to test
func (r *TestPoolRepo) Add(pool *dto.TestPool) error {
	_, err := r.db.Exec(
		`INSERT INTO test_pools (test_id, pool_id) VALUES ($1, $2)`,
		pool.TestID, pool.PoolID,
	)
	return err
}

// Remove pool from test
func (r *TestPoolRepo) Remove(testID, poolID int) error {
	_, err := r.db.Exec(
		`DELETE FROM test_pools WHERE test_id=$1 AND pool_id=$2`,
		testID, poolID,
	)
	return err
}

// List pools of a test
func (r *TestPoolRepo) ListByTest(testID int) ([]dto.TestPool, error) {
	var pools []dto.TestPool
	err := r.db.Select(&pools, `SELECT * FROM test_pools WHERE test_id=$1`, testID)
	return pools, err
}
