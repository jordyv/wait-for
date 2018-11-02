package wait

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
)

type ConditionHTTP struct{}

func (condition ConditionHTTP) Run(ctx context.Context, channel chan error, options ConditionOptions) {
	host := options["host"]
	port := options["port"]
	expectedStatusCode := options["statusCode"]
	forceHttps := options["https"]

	var protocol string
	if forceHttps == true || port == 443 {
		protocol = "https"
	} else {
		protocol = "http"
	}

	url := fmt.Sprintf("%s://%s:%d", protocol, host, port)

	logrus.Debugf("Starting HTTP request to %s", url)

	req, _ := http.NewRequest(http.MethodHead, url, nil)
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		channel <- err
		return
	}

	defer res.Body.Close()

	if res.StatusCode == expectedStatusCode {
		channel <- nil
		return
	} else {
		channel <- errors.New("unexpected status code")
	}
}
