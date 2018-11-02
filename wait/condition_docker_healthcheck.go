package wait

import (
	"context"
	"github.com/docker/docker/client"
	"time"
)

type ConditionDockerHealthcheck struct{}

func getContainerHealthStatus(ctx context.Context, dockerClient *client.Client, container string) (string, error) {
	containerInfo, err := dockerClient.ContainerInspect(ctx, container)
	if err != nil {
		return "", err
	}
	health := containerInfo.State.Health
	return health.Status, nil
}

func (condition ConditionDockerHealthcheck) Run(ctx context.Context, channel chan error, options ConditionOptions) {
	container := options["container"]

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		channel <- err
		return
	}

	for {
		healthStatus, err := getContainerHealthStatus(ctx, dockerClient, container.(string))
		if err != nil {
			channel <- err
			return
		}
		if healthStatus == "healthy" {
			channel <- nil
		}
		time.Sleep(500 * time.Millisecond)
	}
}
