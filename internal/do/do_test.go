package do

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	t.Run("do sum", func(t *testing.T) {
		assert.Equal(t, Sum(1, 2), 3)
	})

}
