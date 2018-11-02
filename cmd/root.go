package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var ExitCodeSuccess = 0
var ExitCodeTimeoutExceeded = 1

var timeout time.Duration
var verbose bool

var rootCmd = &cobra.Command{
	Use: "wait-for",
	Long: `
Wait for

Examples:
  wait-for tcp localhost 8080             Wait till TCP port 8080 at localhost gets up
  wait-for http localhost 8080            Wait till http://localhost:8080 returns 200
  wait-for docker-healthcheck mysql       Wait the Docker healthcheck for container 'mysql' returns healthy
`,
}

func init() {
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "Timeout")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
