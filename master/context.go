package master

type Context struct {
	session *Master
}

func NewContext(master *Master) *Context {
	return &Context{
		session: master,
	}
}

func (c *Context) Extract(job string) *Context {
	return c
}

func (c *Context) Transform(job string) *Context {
	return c
}

func (c *Context) Load(job string) *Context {
	return c
}
