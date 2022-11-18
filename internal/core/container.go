package core

import (
	"reflect"

	"go.uber.org/dig"
)

type Container struct {
	*dig.Container
	controllers []any
}

func NewContainer(container *dig.Container) *Container {
	instance := &Container{
		Container:   container,
		controllers: []any{},
	}

	return instance
}

func (this *Container) GetService(t any) error {
	serviceType := reflect.TypeOf(t).Elem()
	myfunc := reflect.FuncOf([]reflect.Type{serviceType}, []reflect.Type{}, false)

	var mfunType = reflect.MakeFunc(myfunc, func(args []reflect.Value) []reflect.Value {
		reflect.ValueOf(t).Elem().Set(args[0])
		return []reflect.Value{}
	})

	err := this.Invoke(mfunType.Convert(myfunc).Interface().(any))

	return err
}

func (this *Container) AddControllers(constructors ...interface{}) {
	if len(constructors) == 0 {
		return
	}

	for _, element := range constructors {
		this.Container.Provide(element)
		this.controllers = append(this.controllers, element)
	}
}
