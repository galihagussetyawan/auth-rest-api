package repository

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestGetUsers(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("SUCCESS", func(t *mtest.T) {
		t.Log("asdasd")
	})
}
