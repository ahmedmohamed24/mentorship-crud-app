/*
Copyright © 2025 Ahmed Mohamed <ahmedmohamed24.dev@gmail.com>
*/
package cmd

import (
	"github.com/ahmedmohamed24/mentorship-crud-app/data/seeders"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seeding documents",
	Long:  `seeding documents`,
	Run: func(cmd *cobra.Command, args []string) {
		err := seeders.SeedDocuments()
		if err != nil {
			log.Errorln("Seeding failed:", err)
			return
		}
		log.Infoln("Seeding completed successfully")

	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
