package wait

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"net"
)

type ConditionTCP struct{}

func (condition ConditionTCP) Run(ctx context.Context, channel chan error, options ConditionOptions) {
	host := options["host"]
	port := options["port"]

	addr := fmt.Sprintf("%s:%s", host, port)

	logrus.Debugf("Starting TCP request to %s", addr)

	select {
	case <-ctx.Done():
		return
	default:
		con, err := net.Dial("tcp", addr)
		if err == nil {
			defer con.Close()
			channel <- nil
		}
	}
}
