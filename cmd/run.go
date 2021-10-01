package cmd

import (
	"github.com/go-uniform/uniform"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
)

func init() {
	var level string
	var rate int
	var limit int
	var test bool
	// todo: add custom flag variables here

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run " + service.AppName + " service",
		Long:  "Run " + service.AppName + " service",
		Run: func(cmd *cobra.Command, args []string) {
			service.Execute(level, rate, limit, test, _base.NatsUri, _base.CompileNatsOptions(), uniform.M{
				// todo: link custom flags to arg values here, example: "custom": custom,
			})
		},
	}

	// set the service's environment configurations via many command-line-interface (CLI) arguments
	runCmd.Flags().StringVarP(&level, "lvl", "l", "trace", "The logging level ['trace', 'debug', 'info', 'notice', 'warning', 'error', 'fatal'] that service is running in")
	runCmd.Flags().IntVarP(&rate, "rate", "r", 1000, "The sample rate of the trace logs used for performance auditing [set to -1 to log every trace]")
	runCmd.Flags().IntVarP(&limit, "limit", "x", 1000, "The messages per second that each topic worker will be limited to [set to 0 or less for maximum throughput]")
	runCmd.Flags().BoolVar(&test, "test", false, "A flag indicating if service should enter into test mode")
	// todo: add custom CLI flags here

	// todo: add custom CLI flag validations here

	_base.RootCmd.AddCommand(runCmd)
}
