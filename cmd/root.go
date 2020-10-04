package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "pufobs",
	Short: "A \"DAS PODCAST UFO\" archiving tool.",
	Long:  "pufobs is a small tool to list and download \"DAS PODCAST UFO\" episodes.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	downloadCmd.PersistentFlags().BoolVarP(&Overwrite, "overwrite", "o", false, "overwrite existing files")

	RootCmd.AddCommand(countCmd)
	RootCmd.AddCommand(currentCmd)
	RootCmd.AddCommand(downloadCmd)
	RootCmd.AddCommand(listCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
