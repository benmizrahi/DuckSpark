package plugins

import (
	"github.com/benmizrahi/godist/internal/plugins/contract"
	"github.com/benmizrahi/godist/internal/plugins/impl"
)

func MakeBuildIns() map[string]func() contract.IPluginContract {
	mapOfPlugins := map[string]func() contract.IPluginContract{}
	//map all internal
	mapOfPlugins["fsplugin"] = impl.NewFSPlugin
	return mapOfPlugins
}
