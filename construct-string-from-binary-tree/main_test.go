package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree2Str(t *testing.T) {
	assert.Equal(t, "1(2(4))(3)", tree2str(generateTree()))
}
