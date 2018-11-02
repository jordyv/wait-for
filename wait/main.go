package wait

import (
	"context"
	"errors"
	"github.com/Sirupsen/logrus"
	"time"
)

type Condition interface {
	Run(ctx context.Context, channel chan error, options ConditionOptions)
}

type ConditionOptions map[string]interface{}

func Wait(condition Condition, options ConditionOptions, timeout time.Duration) error {
	responseChan := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go condition.Run(ctx, responseChan, options)

	select {
	case err := <-responseChan:
		logrus.Debugf("condition response received; error '%v'", err)
		return err
	case <-time.After(timeout):
		logrus.Debugln("timeout exceeded")
		return errors.New("timeout exceeded")
	}
}
