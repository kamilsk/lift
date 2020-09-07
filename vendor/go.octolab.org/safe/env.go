package safe

import (
	"os"
	"sync"

	"github.com/pkg/errors"
)

var env sync.Mutex

// SetEnvs calls the os.Setenv under a mutex to guarantee consistency between other calls.
//
//  func Test(t *testing.T) {
//  	t.Run("first case", func(t *testing.T) {
//  		release, err := safe.SetEnvs(
//  			"ENV_VAR1", "value_1",
//  			"ENV_VAR2", "value_2",
//  		)
//  		require.NoError(t, err)
//  		defer release()
//
//  		...
//  	})
//
//  	t.Run("second case", func(t *testing.T) {
//  		release, err := safe.SetEnvs(
//  			"ENV_VAR1", "value_3",
//  			"ENV_VAR2", "value_4",
//  		)
//  		require.NoError(t, err)
//  		defer release()
//
//  		...
//  	})
//  }
//
func SetEnvs(envs ...string) (func(func(error)), error) {
	total := len(envs)
	if total == 0 || total%2 != 0 {
		return nil, ErrBadCall
	}

	env.Lock()

	before := make([]*string, total)
	for i := 0; i < total; i += 2 {
		key, value := envs[i], envs[i+1]

		var prev *string
		if value, present := os.LookupEnv(key); present {
			prev = &value
		}
		before[i], before[i+1] = &key, prev

		if err := os.Setenv(key, value); err != nil {

			env.Unlock()

			return nil, errors.Wrapf(
				err,
				"cannot set environment variable %s=%q",
				key, value,
			)
		}
	}

	return func(handler func(error)) {
		defer env.Unlock()

		for i := 0; i < total; i += 2 {
			key, value := before[i], before[i+1]
			if value == nil {
				Do(func() error { return os.Unsetenv(*key) }, handler)
			} else {
				Do(func() error { return os.Setenv(*key, *value) }, handler)
			}
		}
	}, nil
}
