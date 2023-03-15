package contract

type IPluginContract interface {
	//Plugin Name
	Name() string
	//Master Read planning
	PlanRead() []IPartition
	//set configs
	Configs(map[string]string) IPluginContract
}
