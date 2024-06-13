package common

import "github.com/benmizrahi/duckspark/internal/protos"

type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	Plan(args ...interface{}) []*protos.Task
	//set configs
	Configs(map[string]string) IPluginContract
	//runtime
	Execute(*protos.Task) *protos.TaskResult
}

type Options struct {
}
