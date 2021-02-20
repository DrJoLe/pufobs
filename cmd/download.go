package cmd

import (
	"errors"
	"fmt"
	"github.com/drjole/pufobs/internal"
	"github.com/drjole/pufobs/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

var Overwrite bool

func NewDownloadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "download [title] [filepath]",
		Short: fmt.Sprintf("Download a %s episode", pkg.PUFO),
		Long:  fmt.Sprintf("Download a %s episode", pkg.PUFO),
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) >= 1 {
				if _, err := pkg.GetEpisode(args[0]); err != nil {
					return err
				}
			}
			if len(args) >= 2 {
				if !Overwrite && internal.FileExists(args[1]) {
					return errors.New(fmt.Sprintf("file %s exists! (You may use --overwrite to overwrite it)", args[1]))
				}
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			var title string
			var filepath string

			if len(args) >= 1 {
				title = args[0]
			} else {
				title = pkg.GetLatestEpisode().Title
			}
			episode, err := pkg.GetEpisode(title)
			if err != nil {
				log.Fatal(err)
			}

			if len(args) >= 2 {
				filepath = args[1]
			} else {
				wd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				filename := strings.ToLower(strings.ReplaceAll(episode.Title, " ", "_")) + ".mp3"
				filepath = path.Join(wd, filename)
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Downloading episode \"%s\" to %s\n", episode.Title, filepath)

			media, err := http.Get(episode.URL)
			if err != nil {
				log.Fatal(err)
			}

			fp, err := os.Create(filepath)
			defer func() {
				if err := fp.Close(); err != nil {
					log.Fatal(err)
				}
			}()
			n, err := io.Copy(fp, media.Body)
			if err != nil {
				log.Fatal(err)
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%d bytes written", n)
		},
	}
}
