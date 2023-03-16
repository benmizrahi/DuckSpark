package contract

import "github.com/benmizrahi/godist/protos"

type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	PlanRead() []*protos.IPartition
	//set configs
	Configs(map[string]string) IPluginContract
}
