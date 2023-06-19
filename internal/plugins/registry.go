package plugins

import "github.com/benmizrahi/godist/internal/common"

var pluginRegistry = map[string]*common.IPluginContract{}

func init() {
	// add the plugin initializer to the registry
	pluginRegistry = append(pluginRegistry, NewFSPlugin)
}

func GetPlugin(name string) *common.IPluginContract {
	return pluginRegistry[name]
}
