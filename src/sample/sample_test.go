package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Num_Add(t *testing.T) {
	numID := Num(1)
	actual := numID.Add(1)
	assert.Equal(t, Num(2), actual)
}
