package cmd

import (
	"fmt"

	"github.com/patilsuraj767/turf/turf"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start echo server for turf backend",
	Long:  "start echo server for turf backend",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting echo server")
		turf.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
