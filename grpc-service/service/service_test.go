package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kaznowski-brothers/blog-backend/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Creating a new service initiates the database", func(t *testing.T) {
		service := New()
		assert.NotNil(t, service.db)
	})
}

func TestPing(t *testing.T) {
	t.Run("Returns a valid response", func(t *testing.T) {
		service := New()
		response, err := service.Ping(context.TODO(), &empty.Empty{})

		assert.NoError(t, err)

		expectedResponse := &blogapi.Pong{
			Answer: "Pong answer 1",
		}
		assert.Equal(t, response, expectedResponse)
	})
}
