package db

type DB struct {
	numberOfCalls uint64
}

func New() *DB {
	return &DB{numberOfCalls: 0}
}

func (db *DB) DummyCall() uint64 {
	db.numberOfCalls++
	return db.numberOfCalls
}
