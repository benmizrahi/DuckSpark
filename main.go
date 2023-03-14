package main

import (
	"flag"

	"github.com/benmizrahi/godist/common"
	"github.com/benmizrahi/godist/master"
	"github.com/benmizrahi/godist/worker"
)

func main() {

	typeOf := flag.String("type", "Worker", "# Type of instance Master/Worker")
	host := flag.String("host", "localhost", "# Host to listen")
	port := flag.Int("port", 9999, "#port to listen")
	masterPath := flag.String("master", "localhost", "# Master host")
	isLocal := flag.Bool("isLocal", true, "# Run locally with processes K8S/Local")
	flag.Parse()

	var godist *master.Master = nil
	switch *typeOf {
	case "Master":
		godist = master.NewMaster(*isLocal, *host, *port)
	default:
		worker.NewWorker(*host, *port, *masterPath).Init()
	}

	session := common.Exec[master.Master]
					(func() master.Master { return *godist.Start() }).Await()
	session.
		Extract("").
		Transform("").
		Load("")
}
