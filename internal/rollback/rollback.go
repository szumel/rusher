package rollback

import "rusher/internal/platform/container"

const AliasRollbacker = "step.rollbacker"

func init() {
	r := Pool{Rollbackers: []Rollbacker{}}
	container.Set(AliasRollbacker, &r)
}

//there should be needed some context or data
type Rollbacker interface {
	Rollback() error
	Code() string
}

type Pool struct {
	Rollbackers []Rollbacker
}

func Subscribe(r Rollbacker) error {
	rollbacker, err := container.Get(AliasRollbacker)
	if err != nil {
		return err
	}

	rollbacker.(*Pool).Rollbackers = append(rollbacker.(*Pool).Rollbackers, r)

	return nil
}
