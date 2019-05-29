package config_test

import (
	"testing"

	. "github.com/kamilsk/lift/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGoModule(t *testing.T) {
	mod, is := GoModule("testdata")
	assert.True(t, is)
	assert.Equal(t, "go/module/test", mod)
}

func TestGoPackage(t *testing.T) {
	pkg, is := GoPackage("testdata/src/go/package/test")
	assert.True(t, is)
	assert.Equal(t, "go/package/test", pkg)
}
