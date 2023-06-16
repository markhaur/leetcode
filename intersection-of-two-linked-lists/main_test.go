package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	headA, headB, intersectNode := generateListOfNodes()

	assert.Equal(t, getIntersectionNode(headA, headB), intersectNode)
}
