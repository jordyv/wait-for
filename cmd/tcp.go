package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/jordyv/wait-for/wait"
	"github.com/spf13/cobra"
	"os"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Wait for TCP connection",
	Long: `Wait till a TCP connection gets op.

Example:
  wait-for -t 5 tcp localhost 8080         Wait till port 8080 on localhost gets up for max 5 seconds
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		host := args[0]
		port := args[1]
		options := wait.ConditionOptions{
			"host": host,
			"port": port,
		}

		logrus.Infof("Waiting for TCP connection to %s on port %v", host, port)

		if err := wait.Wait(wait.ConditionTCP{}, options, timeout); err != nil {
			logrus.Errorln("condition failed:", err)
			os.Exit(ExitCodeTimeoutExceeded)
		}
		logrus.Infoln("success")
		os.Exit(ExitCodeSuccess)
	},
}

func init() {
	rootCmd.AddCommand(tcpCmd)
}
