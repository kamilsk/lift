package model

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Environment returns a DecodeHookFunc that converts map[string]interface{}
// into EnvironmentVariables.
func Environment() mapstructure.DecodeHookFunc {
	return mapstructure.DecodeHookFuncType(func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		var out EnvironmentVariables
		if f.Kind() != reflect.Map || !t.AssignableTo(reflect.TypeOf(out)) {
			return data, nil
		}
		env := data.(map[string]interface{})
		out = make(EnvironmentVariables, 0, len(env))
		for name, raw := range env {
			value, is := raw.(string)
			if !is {
				return nil, errors.Errorf("decode: %q is not a string: %#v", name, raw)
			}
			out = append(out, EnvironmentVariable{Name: name, Value: value})
		}
		return out, nil
	})
}
