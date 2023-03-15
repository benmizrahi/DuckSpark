package contract

import "github.com/benmizrahi/godist/protos"

type IPartition struct {
	Tasks []*protos.Task
}
type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	PlanRead() []IPartition
	//set configs
	Configs(map[string]string) IPluginContract
}
