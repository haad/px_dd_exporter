package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/haad/px_dd_exporter/log"
)

var rootCmd = &cobra.Command{
	Use:   "px_dd_exporter",
	Short: "Pixxelfederation exproter to export definition data",
	Long:  `Prometheus exporter to export needed data from origin servers.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute Main command execution entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var debug bool
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enables debug logging")

	log.InitLogger(debug)
}
