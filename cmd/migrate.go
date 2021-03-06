package cmd

import (
	"log"

	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/upload"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs migrations and ensures admin user existance",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.DropTableIfExists(&upload.File{})
		db.AutoMigrate(&upload.File{})
	},
}
