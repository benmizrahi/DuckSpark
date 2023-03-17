package plugins

import (
	"github.com/benmizrahi/godist/plugins/contract"
	"github.com/benmizrahi/godist/plugins/impl"
)

func MakeBuildIns() map[string]func() contract.IPluginContract {
	mapOfPlugins := map[string]func() contract.IPluginContract{}
	//map all internal
	mapOfPlugins["fsplugin"] = impl.NewFSPlugin
	return mapOfPlugins
}
