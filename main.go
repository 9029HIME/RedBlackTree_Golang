package main

import (
	"fmt"
	"time"
)

func main() {
	tree := new(Tree)
	tree.Add(70)
	tree.Add(60)
	tree.Add(50)
	tree.Add(40)
	tree.Add(55)
	tree.Add(58)
	tree.Add(56)
	time.Sleep(time.Second * 1)
	tree.LDR(tree.root)
	fmt.Println("根节点是", tree.root.Value)
	fmt.Println(tree.Get(56))
}
