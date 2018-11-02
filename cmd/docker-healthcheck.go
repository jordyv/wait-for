package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/jordyv/wait-for/wait"
	"github.com/spf13/cobra"
	"os"
)

var dockerHealthcheckCmd = &cobra.Command{
	Use:   "docker-healthcheck",
	Short: "Wait for a Docker container to get healthy",
	Long: `Wait till a Docker container's healthcheck returns 'healthy'.

Example:
  wait-for -t 5 docker-healthcheck mysql        Wait till container with name 'mysql' gets healthy
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		container := args[0]
		options := wait.ConditionOptions{
			"container": container,
		}

		logrus.Infof("Waiting for Docker healthcheck of container %s", container)

		if err := wait.Wait(wait.ConditionDockerHealthcheck{}, options, timeout); err != nil {
			logrus.Errorln("condition failed:", err)
			os.Exit(ExitCodeTimeoutExceeded)
		}
		logrus.Infoln("success")
		os.Exit(ExitCodeSuccess)
	},
}

func init() {
	rootCmd.AddCommand(dockerHealthcheckCmd)
}
