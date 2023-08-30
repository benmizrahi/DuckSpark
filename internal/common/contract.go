package common

import (
	"github.com/benmizrahi/gobig/internal/domains"
)

type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	PlanRead() []*domains.IPartition
	//set configs
	Configs(map[string]string) IPluginContract
	//runtime
	Execute(*domains.Task) *domains.TaskResult
}
