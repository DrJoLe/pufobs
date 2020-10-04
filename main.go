package main

import (
	"github.com/drjole/pufobs/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})
	log.SetLevel(log.InfoLevel)
}

func main() {
	cmd.Execute()
}
