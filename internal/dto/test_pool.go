package dto

type TestPool struct {
	TestID int `db:"test_id" json:"test_id"`
	PoolID int `db:"pool_id" json:"pool_id"`
}
