package plugins

import "github.com/benmizrahi/godist/internal/common"

var pluginRegistry = map[string]*common.IPluginContract{}

func init() {
	// add the plugin initializer to the registry
	fs := NewFSPlugin()
	pluginRegistry[fs.Name()] = &fs
}

func GetPlugin(name string) *common.IPluginContract {
	return pluginRegistry[name]
}
