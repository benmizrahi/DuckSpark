package main

import (
	"flag"

	"github.com/benmizrahi/godist/master"
	"github.com/benmizrahi/godist/worker"
)

func main() {

	typeOf := flag.String("type", "Worker", "# Type of instance Master/Worker")
	host := flag.String("host", "localhost", "# Host to listen")
	port := flag.Int("port", 9999, "#port to listen")
	isLocal := flag.Bool("isLocal", true, "# Run locally with processes K8S/Local")
	flag.Parse()

	switch *typeOf {
	case "Master":
		master.NewMaster(*isLocal, *host, *port).Init()
	default:
		worker.NewWorker(*host, *port).Init()
	}
}
