package impl

import "github.com/benmizrahi/godist/plugins/contract"

type FSPlugin struct {
}

// Configs implements contract.IPluginContract
func (FSPlugin) Configs(map[string]string) {
	panic("unimplemented")
}

// Distrbute implements plugins.IPluginContract
func (FSPlugin) Distrbute(dist contract.IDistrbution, task contract.ITask) contract.ITaskResult {
	panic("unimplemented")
}

// Name implements plugins.IPluginContract
func (FSPlugin) Name() string {
	return "godist-fsplugin"
}

// Plan implements plugins.IPluginContract
func (FSPlugin) PlanRead() []contract.IDistrbution {
	panic("unimplemented")
}

// Name must be New + struct name
func NewFSPlugin() contract.IPluginContract {
	return FSPlugin{}
}
