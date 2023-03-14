package main

import (
	"flag"

	"github.com/benmizrahi/godist/master"
)

func main() {

	host := flag.String("host", "localhost", "# Host to listen")
	port := flag.Int("port", 9999, "#port to listen")
	isLocal := flag.Bool("isLocal", true, "# Run locally with processes K8S/Local")
	flag.Parse()

	sc := master.NewMaster(*isLocal, *host, *port, 2).Context()
	sc.
		Extract("").
		Transform("").
		Load("")
}
