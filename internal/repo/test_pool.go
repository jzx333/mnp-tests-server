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

func (r *TestPoolRepo) Add(tp *dto.TestPool) error {
	_, err := r.db.Exec("INSERT INTO test_pools (test_id, pool_id) VALUES ($1,$2)", tp.TestID, tp.PoolID)
	return err
}

func (r *TestPoolRepo) Remove(tp *dto.TestPool) error {
	_, err := r.db.Exec("DELETE FROM test_pools WHERE test_id=$1 AND pool_id=$2", tp.TestID, tp.PoolID)
	return err
}

func (r *TestPoolRepo) GetByTest(testID int) ([]*dto.TestPool, error) {
	var tps []*dto.TestPool
	err := r.db.Select(&tps, "SELECT test_id, pool_id FROM test_pools WHERE test_id=$1", testID)
	return tps, err
}

func (r *TestPoolRepo) GetByPool(poolID int) ([]*dto.TestPool, error) {
	var tps []*dto.TestPool
	err := r.db.Select(&tps, "SELECT test_id, pool_id FROM test_pools WHERE pool_id=$1", poolID)
	return tps, err
}
