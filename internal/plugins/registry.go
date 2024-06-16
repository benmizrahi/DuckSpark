package plugins

var pluginRegistry = map[string]IPluginContract{}

func init() {
	// add the plugin initializer to the registry
	fs := NewFSPlugin()
	pluginRegistry[fs.Name()] = fs
}

func GetPlugin(name string) IPluginContract {
	return pluginRegistry[name]
}
