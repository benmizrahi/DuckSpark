package master

import (
	"github.com/benmizrahi/godist/plugins/contract"
	"github.com/benmizrahi/godist/protos"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Context struct {
	session *Master
	plugins map[string]contract.IPluginContract
	plan    []protos.IPartition
}

func NewContext(master *Master) *Context {
	return &Context{
		session: master,
		plugins: map[string]contract.IPluginContract{},
	}
}

func (c *Context) Extract(from string, config map[string]string) *Context {
	plugin := c.session.Plugins[from]
	c.plugins[from] = plugin().Configs(config)
	c.plan = c.plugins[from].PlanRead()
	return c
}

func (c *Context) Transform() *Context {
	return c
}

func (c *Context) Load(job string) *Context {
	return c
}

func (c *Context) Show() *Context {
	for _, partition := range c.plan {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  []string{"println"},
			CreationTime: timestamppb.Now(),
		})
	}

	c.session.DoAction(c.plan)
	return c
}
