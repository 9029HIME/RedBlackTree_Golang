package main

import (
	"fmt"
	"time"
)

func main() {
	tree := new(Tree)
	var i int64
	for i = 1; i <= 10; i++ {
		tree.Add(i)
	}
	time.Sleep(time.Second * 1)
	tree.LDR(tree.root)
	fmt.Println("根节点是", tree.root.Value)
	time.Sleep(time.Second * 1)
	tree.Delete(1)
	time.Sleep(time.Second * 1)
	tree.LDR(tree.root)
}
