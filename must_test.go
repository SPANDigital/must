package must

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMust(t *testing.T) {
	t.Run("Test panic", func(t *testing.T) {
		assert.Panics(t, func() {
			Must("errors", errors.New("error"))
		})
	})
	t.Run("Test not panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Must("no error", nil)
		})
	})
}
