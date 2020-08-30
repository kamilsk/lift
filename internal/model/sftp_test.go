package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestSFTP_Merge(t *testing.T) {
	t.Run("nil sftp", func(t *testing.T) {
		var sftp *SFTP
		assert.NotPanics(t, func() { sftp.Merge(&SFTP{Size: "small"}) })
		assert.Nil(t, sftp)
	})
}
