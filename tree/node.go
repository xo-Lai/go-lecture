package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 结构体的方法定义，（node treeNode）是接受者
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node . Ignored")
		return
	}
	node.value = value
}

// 工厂函数，来构造函数,注意返回局部变量的地址
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)

	root.traverse() //0 2 3 4 5
	//root.right.left.print()
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
	//var pRoot *treeNode
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, nil},
	//}
	//fmt.Println(nodes)

}
