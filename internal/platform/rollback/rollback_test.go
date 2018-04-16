package rollback

import (
	"github.com/szumel/rusher/internal/platform/container"
	"testing"
)

func TestSubscribe(t *testing.T) {
	r := rollbackerMock{}
	err := Subscribe(&r)
	if err != nil {
		t.Fatal("Rollbacker subscribe failed.")
	}

	pool, err := container.Get(AliasRollbacker)
	if err != nil {
		t.Fatal("Could not retrieve rollbackers pool from container.")
	}

	concretePool, ok := pool.(*Pool)
	if !ok {
		t.Fatal("Subscribe rollbacker failed. Could not assert to expected type.")
	}

	expected := concretePool.Rollbackers[0]
	if expected.Code() != r.Code() {
		t.Fatalf("Subscribe rollbacker failed. Expected code %s have %s", expected.Code(), r.Code())
	}
}

type rollbackerMock struct{}

func (r *rollbackerMock) Rollback() error {
	return nil
}

func (*rollbackerMock) Code() string {
	return "test"
}
