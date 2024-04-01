package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.NoError(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(-1000)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)

}
