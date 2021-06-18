package main

import "fmt"

type Node struct {
    val   int
    left  *Node
    right *Node
}

func InorderRecursive(root *Node) {
    if root == nil {
        return
    }

    InorderRecursive(root.left)
    fmt.Printf("%d ", root.val)
    InorderRecursive(root.right)
}

func main() {
    /*
                10
               /  \
             20    30
            / \      \
           40  50     60
          /
         70
    */

    root := &Node{22, nil, nil}
    root.left = &Node{28, nil, nil}
    root.right = &Node{31, nil, nil}
    root.left.left = &Node{23, nil, nil}
    root.left.right = &Node{22, nil, nil}
    root.right.right = &Node{40, nil, nil}
    root.left.left.left = &Node{12, nil, nil}
	root.left.left.right = &Node{80, nil, nil}
	root.left.right.left = &Node{25, nil, nil}
	root.left.right.right = &Node{27, nil, nil}
	root.right.right.left = &Node{70, nil, nil}
	root.right.right.right = &Node{40, nil, nil}
	root.left.left.left.left = &Node{17, nil, nil}
	root.left.left.left.right = &Node{12, nil, nil}
	root.left.left.right.left = &Node{22, nil, nil}
	root.left.left.right.right = &Node{72, nil, nil}










    fmt.Println("Inorder Traversal - recursive solution : ")
    InorderRecursive(root)
}