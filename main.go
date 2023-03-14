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
	sc := master.NewMaster(*isLocal, *host, *port, 2).Context()

	log.Info("******************")
	log.Info("******************")
	log.Info("Status: Read!")
	log.Info("your wish is my command... lets GO!!")
	log.Info("******************")
	log.Info("******************")
	sc.
		Extract("").
		Transform("").
		Load("")
}
