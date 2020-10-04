package cmd

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/drjole/pufobs/internal"
	"github.com/drjole/pufobs/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path"
)

var Overwrite bool

var downloadCmd = &cobra.Command{
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

		page, err := http.Get(episode.URL)
		defer func() {
			if err := page.Body.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		doc, err := goquery.NewDocumentFromReader(page.Body)
		if err != nil {
			log.Fatal(err)
		}
		href, ok := doc.Find(".powerpress_link_d").First().Attr("href")
		if !ok {
			log.Fatalf("error parsing podcast-ufo.fail")
		}

		if len(args) >= 2 {
			filepath = args[1]
		} else {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			filename, ok := doc.Find(".powerpress_link_d").First().Attr("download")
			if !ok {
				log.Fatalf("error parsing podcast-ufo.fail")
			}
			filepath = path.Join(wd, filename)
		}

		media, err := http.Get(href)
		defer func() {
			if err := page.Body.Close(); err != nil {
				log.Fatal(err)
			}
		}()

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
		fmt.Printf("%d bytes written", n)
	},
}
