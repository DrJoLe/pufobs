package cmd

import (
	"fmt"
	"github.com/drjole/pufobs/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: fmt.Sprintf("List all available %s episodes", pkg.PUFO),
	Long:  fmt.Sprintf("List all available %s episodes", pkg.PUFO),
	Run: func(cmd *cobra.Command, args []string) {
		episodes := pkg.GetEpisodes()
		for _, e := range episodes {
			fmt.Println(e.Title)
		}
	},
}
