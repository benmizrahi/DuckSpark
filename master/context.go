package master

import "github.com/benmizrahi/godist/plugins/contract"

type Context struct {
	session *Master
	plugins map[string]contract.IPluginContract
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
	plan := c.plugins[from].PlanRead()

	return c
}

func (c *Context) Transform(job string) *Context {
	return c
}

func (c *Context) Load(job string) *Context {
	return c
}
