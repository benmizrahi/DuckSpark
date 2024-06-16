package plugins

import "github.com/benmizrahi/duckspark/internal/common"

type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	Plan(args ...interface{}) common.Maplan
	//set configs
	Configs(map[string]string) IPluginContract
}

type Options struct {
}
