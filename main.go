package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

func main() {

	// host := flag.String("host", "localhost", "# Host to listen")
	// port := flag.Int("port", 9999, "#port to listen")
	// isLocal := flag.Bool("isLocal", true, "# Run locally with processes K8S/Local")

	flag.Parse()

	log.Info("gobig Status: Starting")
	log.Info("******************")
	log.Info("Status: Ready!")
	log.Info("******************")

}
