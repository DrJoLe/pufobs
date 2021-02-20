package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pufobs",
		Short: "A \"DAS PODCAST UFO\" archiving tool.",
		Long:  "pufobs is a small tool to list and download \"DAS PODCAST UFO\" episodes.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
		},
	}
}

func init() {
	downloadCmd := NewDownloadCmd()
	downloadCmd.PersistentFlags().BoolVarP(&Overwrite, "overwrite", "o", false, "overwrite existing files")

	rootCmd.AddCommand(NewCountCmd())
	rootCmd.AddCommand(NewCurrentCmd())
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(NewListCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
