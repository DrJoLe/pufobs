package cmd

import (
	"fmt"
	"github.com/drjole/pufobs/pkg"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: fmt.Sprintf("List all available %s episodes", pkg.PUFO),
		Long:  fmt.Sprintf("List all available %s episodes", pkg.PUFO),
		Run: func(cmd *cobra.Command, args []string) {
			episodes := pkg.GetEpisodes()
			for _, e := range episodes {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), e.Title)
			}
		},
	}
}
