package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/jordyv/wait-for/wait"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var (
	statusCode int
	forceHttps bool

	defaultPort = 80
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Wait for an HTTP connection",
	Long: `Wait till an HTTP connection returns a status code.

Example:
  wait-for -t 5 -s 200 http localhost 8080         Wait till http://localhost:8080 returns a 200 status code.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			return nil
		}
		if len(args) == 1 {
			return nil
		}

		return errors.New("expecting at least 1 argument")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		host := args[0]
		var port int
		if len(args) == 2 {
			port, _ = strconv.Atoi(args[1])
		}
		if port == 0 {
			port = defaultPort
		}

		options := wait.ConditionOptions{
			"host":       host,
			"port":       port,
			"statusCode": statusCode,
			"https":      forceHttps,
		}

		logrus.Infof("Waiting for HTTP status code %d on %s at port %v", statusCode, host, port)

		if err := wait.Wait(wait.ConditionHTTP{}, options, timeout); err != nil {
			logrus.Errorln("condition failed:", err)
			os.Exit(ExitCodeTimeoutExceeded)
		}
		logrus.Infoln("success")
		os.Exit(ExitCodeSuccess)
	},
}

func init() {
	httpCmd.Flags().IntVarP(&statusCode, "status", "c", 200, "Expected HTTP status code")
	httpCmd.Flags().BoolVarP(&forceHttps, "https", "s", false, "Force use of HTTPS request")

	rootCmd.AddCommand(httpCmd)
}
