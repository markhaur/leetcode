package main

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := generateNAryTree()
	fmt.Printf("maxDepth: %d\n", maxDepth(root))
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	mDepth := 1
	for _, child := range root.Children {
		mDepth = max(mDepth, maxDepth(child)+1)
	}

	return mDepth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func generateNAryTree() *Node {
	leaf := Node{Val: 14}

	// third level
	eleven := Node{Val: 11, Children: []*Node{&leaf}}
	tweleve := Node{Val: 12}
	thirteen := Node{Val: 13}

	// second level
	six := Node{Val: 6}
	seven := Node{Val: 7, Children: []*Node{&eleven}}
	eight := Node{Val: 8, Children: []*Node{&tweleve}}
	nine := Node{Val: 9, Children: []*Node{&thirteen}}
	ten := Node{Val: 10}

	// first level
	two := Node{Val: 2}
	three := Node{Val: 3, Children: []*Node{&six, &seven}}
	four := Node{Val: 4, Children: []*Node{&eight}}
	five := Node{Val: 5, Children: []*Node{&nine, &ten}}

	// root node
	root := Node{Val: 1, Children: []*Node{&two, &three, &four, &five}}

	return &root

}
