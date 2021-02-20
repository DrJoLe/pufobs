package cmd

import (
	"fmt"
	"github.com/drjole/pufobs/pkg"
	"github.com/spf13/cobra"
	"sort"
)

func NewCurrentCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "current",
		Short: fmt.Sprintf("Print the name of the current %s episode", pkg.PUFO),
		Long:  fmt.Sprintf("Print the name of the current %s episode", pkg.PUFO),
		Run: func(cmd *cobra.Command, args []string) {
			episodes := pkg.GetEpisodes()
			sort.Slice(episodes, func(i, j int) bool {
				return episodes[i].Published.After(*episodes[j].Published)
			})
			fmt.Println(episodes[0])
		},
	}
}
