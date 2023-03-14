package master

type Context struct {
	session *Master
}

func NewContext(master *Master) *Context {
	return &Context{
		session: master,
	}
}

func (c *Context) Extract(from string, config map[string]string) *Context {
	c.session.Plugins[from].Configs(config)
	return c
}

func (c *Context) Transform(job string) *Context {
	return c
}

func (c *Context) Load(job string) *Context {
	return c
}
