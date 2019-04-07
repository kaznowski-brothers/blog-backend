package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Instantiates with 0 cals made", func(t *testing.T) {
		db := New()
		assert.Equal(t, uint64(0), db.numberOfCalls)
	})
}

func TestDummyCall(t *testing.T) {
	t.Run("Returns the number of calls", func(t *testing.T) {
		db := New()

		response := db.DummyCall()
		assert.Equal(t, uint64(1), response)
	})

	t.Run("Increases the number of calls with each call", func(t *testing.T) {
		db := New()

		response := db.DummyCall()
		assert.Equal(t, uint64(1), response)

		response = db.DummyCall()
		assert.Equal(t, uint64(2), response)
	})
}
