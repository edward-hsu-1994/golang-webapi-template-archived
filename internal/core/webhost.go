package core

type WebHost struct {
	container *Container
	runFunc   any
}

func (this *WebHost) Run() error {
	err := this.container.Invoke(this.runFunc)
	if err != nil {
		return err
	}
	return nil
}
