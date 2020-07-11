package model_test

import (
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestEnvironment(t *testing.T) {
	tests := map[string]struct {
		from     reflect.Type
		to       reflect.Type
		data     interface{}
		expected interface{}
		error    string
	}{
		"bad kind": {
			from:  reflect.TypeOf(map[string]interface{}{}),
			to:    reflect.TypeOf(EnvironmentVariables{}),
			data:  map[string]interface{}{"good": "value", "bad": 0},
			error: `decode: "bad" is not a string: 0`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			hook := Environment().(mapstructure.DecodeHookFuncType)
			out, err := hook(test.from, test.to, test.data)
			if test.error != "" {
				assert.Nil(t, out)
				assert.EqualError(t, err, test.error)
				return
			}
			assert.Equal(t, test.expected, out)
			assert.NoError(t, err)
		})
	}
}
