package core

import (
	"reflect"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
)

type Configuration map[string]any

const configurationPathSplitSymbol = ":"

func processPath(path string) []string {
	if len(path) == 0 {
		return []string{}
	}

	return strings.Split(path, configurationPathSplitSymbol)
}

func (this Configuration) GetValue(path ...string) any {
	return GetValue(this, path...)
}

func GetValue(config map[string]any, path ...string) any {
	// https://gobyexample.com/generics
	// https://github.com/ahmetb/go-linq
	if From(path).AnyWithT(func(subPath string) bool {
		return strings.Contains(subPath, configurationPathSplitSymbol)
	}) == false {
		if len(path) == 0 {
			return nil
		}

		value := config[path[0]]

		if len(path) == 1 {
			return value
		}
		if value == nil && len(path) > 1 {
			return nil
		}

		if reflect.TypeOf(value).Implements(reflect.TypeOf((Configuration)(nil)).Elem()) {
			return GetValue(value.(map[string]any), path[1:]...)
		} else {
			return nil
		}
	}

	var processedPath []string
	From(path).SelectManyT(func(item string) Query {
		return From(processPath(item))
	}).ToSlice(&processedPath)

	return GetValue(config, processedPath...)
}
