package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 6, 10, 18, 22, 35}

	i := BinarySearch(a, 10)
	assert.Equal(t, 2, i)

	i = BinarySearch(a, 20)
	assert.Equal(t, -1, i)
}
