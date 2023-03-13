package main

import (
	"flag"

	"github.com/benmizrahi/godistodist/master"
	"github.com/benmizrahi/godistodist/worker"
)

func main() {

	typeOf := *flag.String("TypeOf", "Worker", "# Type of instance Master/Worker")
	isLocal := *flag.Bool("TypeOf", true, "# Run locally with processes K8S/Local")
	masterHost := *flag.String("MasterHost", "localhost", "# Where the master is hosted")
	masterPort := *flag.Int("MasterPort", 9999, "# Where the master port")

	switch typeOf {
	case "Master":
		master.NewMaster(isLocal).Init()
	default:
		worker.NewWorker(masterHost, masterPort).Init()
	}
}
