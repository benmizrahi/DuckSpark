package common

import "github.com/benmizrahi/gobig/internal/protos"

type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	PlanRead() []*protos.IPartition
	//set configs
	Configs(map[string]string) IPluginContract
	//runtime
	Execute(*protos.Task) *protos.TaskResult
}

type Options struct {
	Columns    []string
	Repartiton int
}
