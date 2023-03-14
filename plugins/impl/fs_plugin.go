package impl

import (
	"io/ioutil"
	"log"

	"github.com/benmizrahi/godist/plugins/contract"
)

type FSPlugin struct {
	Path   string
	Format string
}

// Configs implements contract.IPluginContract
func (p FSPlugin) Configs(conf map[string]string) contract.IPluginContract {
	p.Format = conf["format"]
	p.Path = conf["path"]
	return p
}

// Distrbute implements plugins.IPluginContract
func (p FSPlugin) Distrbute(dist contract.IDistrbution, task contract.ITask) contract.ITaskResult {
	panic("unimplemented")
}

// Name implements plugins.IPluginContract
func (p FSPlugin) Name() string {
	return "fsplugin"
}

// Plan implements plugins.IPluginContract
func (p FSPlugin) PlanRead() []contract.IDistrbution {
	files, err := ioutil.ReadDir(p.Path)
	if err != nil {
		log.Fatal(err)
	}
	distribution := []contract.IDistrbution{}
	for _, file := range files {
		distribution = append(distribution, contract.IDistrbution{Dynamic: file})
	}
	return distribution
}

// Name must be New + struct name
func NewFSPlugin() contract.IPluginContract {
	return FSPlugin{}
}
