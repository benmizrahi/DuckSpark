package plugins

var pluginRegistry = map[string]IPluginContract{}

func init() {
	// add the plugin initializer to the registry
	fs := NewFSPlugin()
	duckdb := NewDuckPlugin()
	pluginRegistry[fs.Name()] = fs
	pluginRegistry[duckdb.Name()] = duckdb
}

func GetPlugin(name string) IPluginContract {
	return pluginRegistry[name]
}
