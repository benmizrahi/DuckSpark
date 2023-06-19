package main

import (
	"flag"

	"github.com/benmizrahi/godist/internal/master"
	log "github.com/sirupsen/logrus"
)

func main() {

	host := flag.String("host", "localhost", "# Host to listen")
	port := flag.Int("port", 9999, "#port to listen")
	isLocal := flag.Bool("isLocal", true, "# Run locally with processes K8S/Local")
	flag.Parse()

	log.Info("Godist Status: Starting")
	log.Info("******************")
	log.Info("Status: Ready!")
	log.Info("your wish is my command... lets GO!!")
	log.Info("******************")

	master.
		NewMaster(*isLocal, *host, *port, 2).
		Parallelize("ID,Name \n 1,BEN").
		Show()
}
