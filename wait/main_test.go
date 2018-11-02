package wait

import (
	"context"
	"github.com/pkg/errors"
	"testing"
	"time"
)

var resultSuccess = 1
var resultError = 2
var resultTimeout = 3

type condition struct {
	Result int
}

func (c condition) Run(ctx context.Context, channel chan error, options ConditionOptions) {
	switch c.Result {
	case resultError:
		channel <- errors.New("error occurred")
		return
	case resultSuccess:
		channel <- nil
		return
	}
}

func TestWait_Success(t *testing.T) {
	condition := condition{resultSuccess}
	err := Wait(condition, ConditionOptions{}, time.Second)

	if err != nil {
		t.Error(err)
	}
}

func TestWait_Error(t *testing.T) {
	condition := condition{resultError}
	err := Wait(condition, ConditionOptions{}, time.Second)

	if err == nil {
		t.Error(err)
	}
	if err.Error() != "error occurred" {
		t.Error("expecting error occurred")
	}
}

func TestWait_Timeout(t *testing.T) {
	condition := condition{resultTimeout}
	err := Wait(condition, ConditionOptions{}, time.Second)

	if err == nil {
		t.Error(err)
	}
	if err.Error() != "timeout exceeded" {
		t.Error("expecting timeout exceeded")
	}
}
