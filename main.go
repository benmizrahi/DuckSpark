package main

import (
	"flag"

	"github.com/benmizrahi/godist/master"
	log "github.com/sirupsen/logrus"
)

func main() {

	host := flag.String("host", "localhost", "# Host to listen")
	port := flag.Int("port", 9999, "#port to listen")
	isLocal := flag.Bool("isLocal", true, "# Run locally with processes K8S/Local")
	flag.Parse()
	log.Info("GoDist Status: Starting")
	godist := master.NewMaster(*isLocal, *host, *port, 2)

	log.Info("******************")
	log.Info("Status: Ready!")
	log.Info("your wish is my command... lets GO!!")
	log.Info("******************")

	godist.
		Context().
		Extract("fsplugin", map[string]string{"path": ".extra/data/", "format": "csv", "parallelism": "5"}).
		Show()
}
