package main

import (
	"fmt"
	"lai/go-lecture/tree"
)

// 想要扩展类型或者别人的结构
// 可以通过组合或者别名来扩展，组合myTreeNode  ,别名Queue

// 组合的方式扩展结构
type myTreeNode struct {
	node *tree.Node
}

// 组合的方式扩展方法
func (myNode *myTreeNode) postOrder() {

	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse() //0 2 3 4 5
	fmt.Println()
	myNode := myTreeNode{&root}
	myNode.postOrder()

	//root.Right.Left.print()
	//
	//root.setValue(100)
	//root.print()

	// 复制变量
	//pRoot:=&root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()
	//root.print()

	// 赋值nil 值
	//var pRoot *tree.Node
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	//nodes := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, nil},
	//}
	//fmt.Println(nodes)

}
