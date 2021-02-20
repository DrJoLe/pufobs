package cmd

import (
	"fmt"
	"github.com/drjole/pufobs/pkg"
	"github.com/spf13/cobra"
)

func NewCountCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "count",
		Short: fmt.Sprintf("Print the number of currently available %s episodes", pkg.PUFO),
		Long:  fmt.Sprintf("Print the number of currently available %s episodes", pkg.PUFO),
		Run: func(cmd *cobra.Command, args []string) {
			episodes := pkg.GetEpisodes()
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), len(episodes))
		},
	}
}
