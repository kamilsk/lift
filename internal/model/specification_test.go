package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestSpecification_Merge(t *testing.T) {
	t.Run("nil specification", func(t *testing.T) {
		var spec *Specification
		assert.NotPanics(t, func() { spec.Merge(&Specification{Name: "test"}) })
		assert.Nil(t, spec)
	})
}
