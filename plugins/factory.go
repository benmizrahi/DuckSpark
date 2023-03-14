package plugins

import (
	"github.com/benmizrahi/godist/plugins/contract"
	"github.com/benmizrahi/godist/plugins/impl"
)

func MakeBuildIns() []contract.IPluginContract {
	internals := []contract.IPluginContract{}
	internals = append(internals, impl.NewFSPlugin())
	return internals
}
