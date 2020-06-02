package demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackMockRecurse(t *testing.T) {
	assert.Equal(t, 15, StackMockRecurse(5))
}

func TestFeibo(t *testing.T) {
	// 1,1, 2,3,5,8,13
	assert.Equal(t, 13, Feibo(7))
	assert.Equal(t, 5, Feibo(5))
}

func TestStackFeibo(t *testing.T) {
	assert.Equal(t, 13, StackFeibo(7))
	assert.Equal(t, 5, StackFeibo(5))
}
