package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 结构体的方法定义，（node Node）是接受者
func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node . Ignored")
		return
	}
	node.Value = value
}

// 工厂函数，来构造函数,注意返回局部变量的地址
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
