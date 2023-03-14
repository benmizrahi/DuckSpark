package plugins

type IDistrbution struct{}

type ITask struct{}

type ITaskResult struct{}

type IPluginContract interface {
	Name() string
	//master
	Plan() []IDistrbution
	//worker method
	Distrbute(dist IDistrbution, task ITask) ITaskResult
}
