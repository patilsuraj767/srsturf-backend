package cmd

import (
	"fmt"

	"github.com/patilsuraj767/turf/turf/config"
	"github.com/patilsuraj767/turf/turf/db"
	"github.com/patilsuraj767/turf/turf/model"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting database migration")
		config.InitConfig()
		db.InitDatabase()
		db.DBConn.AutoMigrate(&model.Customer{}, &model.Booking{})
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed database",
	Long:  "seed database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seed database")
	},
}

var dbCmd = &cobra.Command{Use: "db", Short: "Use to migrate or seed database"}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(migrateCmd)
	dbCmd.AddCommand(seedCmd)
}
