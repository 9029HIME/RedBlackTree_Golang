package main

import (
	"fmt"
	"log"
)


/**
1.root必须为黑色
2.Nil为黑节点
3.所有叶子节点为黑色（包括Nil）
4.红节点必须要有两个黑节点（即不能有红节点连红节点）
5.默认新增的节点为红色
*/

/**
新增时的平衡规则：
1.左旋（插入节点为右树右节点时）:
	父节点绕着祖父节点逆时针旋转，使父节点变为祖父节点，原祖父节点变为父节点的左子节点，且原左子节点连接到原祖父节点的右子节点
	最后父节点变为黑色，祖父节点变为红色
2.右旋（插入节点为左树左节点时）：
	父节点绕着祖父节点顺时针旋转，使父节点变为祖父节点，原祖父节点变为父节点的右子节点，且原右子节点连接到原祖父节点的左子节点
	最后父节点变为黑色，祖父节点变为红色
3.左内旋/RL旋转（插入节点为右树左节点时）：新节点与祖父节点互换位置，此时祖父节点变为新节点左子节点，父节点变为新节点右子节点
	最后新节点变为黑色，祖父节点变为红色
4.右内旋/LR旋转（插入节点为左树右节点时）：新节点与祖父节点互换位置，此时祖父节点变为新节点右子节点，父节点变为新节点左子节点
	最后新节点变为黑色，祖父节点变为红色
5.插入第一个节点时，红变黑
6.如果父节点为黑，直接插入
7.如果父节点为红，叔叔为红，爷爷节点为黑（必然），则父和叔叔变为黑，爷爷节点变为红。如爷爷节点为根节点，再转变为黑色，
如果爷爷节点不是根节点则不符合规则，此时需要将爷爷节点作为新增节点递归判断
8.如果父节点是红色，且新增节点没有叔叔节点/叔叔节点为黑节点，视情况做上面四种旋转操作
*/

/**
此方法用于插入数据后平衡，需要在方法体内判断用哪种平衡方式
*/
func (tree *Tree) Balance(node *TreeNode) {
	var L bool = false
	var	R bool = false
	var RL bool = false
	var LR bool = false

	//如果是第一个节点
	if tree.root == node {
		log.Println("第一个节点而已，直接变黑")
		node.IsBlack = true
		return
	}
	//获取父、叔、祖父节点
	father := node.Father
	grandFather := father.Father
	var uncle *TreeNode
	if grandFather != nil && grandFather.LeftSon == father {
		uncle = grandFather.RightSon
		if node.Value > father.Value{
			LR = true
		}else{
			R = true
		}
	} else if grandFather != nil && grandFather.RightSon == father {
		uncle = grandFather.LeftSon
		if node.Value > father.Value{
			L = true
		}else{
			RL = true
		}
	}
	//7.
	if uncle != nil && uncle.IsBlack == false && (uncle.IsBlack == father.IsBlack) {
		log.Println("准备进入第7点")
		uncle.IsBlack = true
		father.IsBlack = true
		grandFather.IsBlack = false
		//将爷爷节点作为新增节点递归判断
		tree.Balance(grandFather)
		return
	}
	//8.
	//记得，进行旋转的大前提是有祖父
	if grandFather!=nil && father.IsBlack == false && (uncle == nil || uncle.IsBlack == true){
		log.Println("准备旋转")
		if L{
			log.Println("进入了L模式")
			tree.Left(node)
			return
		}else if R{
			log.Println("进入了R模式")
			tree.Right(node)
			return
		}else if LR{
			log.Println("进入了LR模式")
			tree.LeftRight(node)
			return
		}else if RL{
			log.Println("进入了RL模式")
			tree.RightLeft(node)
			return
		}
	}

	//如果父节点为黑色，直接插入
	log.Println("父节点为黑，直接插入")
}

func (tree *Tree) Left(node *TreeNode) {
	/**
	左旋（插入节点为右树右节点时）:
	父节点绕着祖父节点逆时针旋转，使父节点变为祖父节点，原祖父节点变为父节点的左子节点，且原左子节点连接到原祖父节点的右子节点
	最后父节点变为黑色，祖父节点变为红色
	 */
	father := node.Father
	grandFather := father.Father
	brother := father.LeftSon
	greatGrandFather := grandFather.Father
	//更改记得要成对出现，不然会信息不对等（如"我的儿子的父亲不是我"）
	father.LeftSon = grandFather
	grandFather.Father = father

	grandFather.RightSon = brother
	if brother != nil{
		brother.Father = grandFather
	}

	if greatGrandFather!=nil && greatGrandFather.LeftSon == grandFather{
		greatGrandFather.LeftSon = father
	}else if greatGrandFather!=nil && greatGrandFather.RightSon == grandFather{
		greatGrandFather.RightSon = father
	}else if greatGrandFather==nil{
		//说明grandFather是root，需要将root指针重新指定
		tree.root = father
	}
	father.Father = greatGrandFather
	//插入node后再进行balance，要注意node的grandFather也要变
	node.GrandFather = father.Father

	father.IsBlack = true
	grandFather.IsBlack = false

}

func (tree *Tree)Right(node *TreeNode) {
	/**
	右旋（插入节点为左树左节点时）：
	父节点绕着祖父节点顺时针旋转，使父节点变为祖父节点，原祖父节点变为父节点的右子节点，且原右子节点连接到原祖父节点的左子节点
	最后父节点变为黑色，祖父节点变为红色
	 */
	father := node.Father
	grandFather := father.Father
	brother := father.RightSon
	greatGrandFather := grandFather.Father
	//更改记得要成对出现，不然会信息不对等（如"我的儿子的父亲不是我"）
	father.RightSon = grandFather
	grandFather.Father = father

	grandFather.LeftSon = brother
	if brother != nil{
		brother.Father = grandFather
	}

	if greatGrandFather!=nil && greatGrandFather.LeftSon == grandFather{
		greatGrandFather.LeftSon = father
	}else if greatGrandFather!=nil && greatGrandFather.RightSon == grandFather{
		greatGrandFather.RightSon = father
	}else if greatGrandFather==nil{
		//说明grandFather是root，需要将root指针重新指定
		tree.root = father
	}
	father.Father = greatGrandFather
	//插入node后再进行balance，要注意node的grandFather也要变
	node.GrandFather = father.Father

	father.IsBlack = true
	grandFather.IsBlack = false
}

func (tree *Tree)LeftRight(node *TreeNode) {
	/**
	新节点与祖父节点互换位置，此时祖父节点变为新节点右子节点，父节点变为新节点左子节点
	最后新节点变为黑色，祖父节点变为红色
	 */
	father := node.Father
	grandFather := father.Father
	greatGrandFather := grandFather.Father
	node.LeftSon = father
	node.RightSon = grandFather
	father.Father = node
	grandFather.Father = node
	//记得去掉父与祖父的子节点，不然会形成环路
	father.RightSon = nil
	grandFather.LeftSon = nil


	if greatGrandFather!=nil && greatGrandFather.LeftSon == grandFather{
		greatGrandFather.LeftSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	}else if greatGrandFather!=nil && greatGrandFather.RightSon == grandFather{
		greatGrandFather.RightSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	}else if greatGrandFather==nil{
		//说明grandFather是root，需要将root指针重新指定
		//TODO 很奇怪，如果这里传的是root的指针会无效，必须要传tree.root
		tree.root = node
		node.Father = nil
		node.GrandFather = nil
	}


	node.IsBlack = true
	grandFather.IsBlack = false

}

func (tree *Tree)RightLeft(node *TreeNode) {
	/**
	新节点与祖父节点互换位置，此时祖父节点变为新节点左子节点，父节点变为新节点右子节点
	最后新节点变为黑色，祖父节点变为红色
	*/
	father := node.Father
	grandFather := father.Father
	greatGrandFather := grandFather.Father
	node.RightSon = father
	node.LeftSon = grandFather
	father.Father = node
	grandFather.Father = node
	//记得去掉父与祖父的子节点，不然会形成环路
	father.LeftSon = nil
	grandFather.RightSon = nil


	if greatGrandFather!=nil && greatGrandFather.LeftSon == grandFather{
		greatGrandFather.LeftSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	}else if greatGrandFather!=nil && greatGrandFather.RightSon == grandFather{
		greatGrandFather.RightSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	}else if greatGrandFather==nil{
		//说明grandFather是root，需要将root指针重新指定
		//TODO 很奇怪，如果这里传的是root的指针会无效，必须要传tree.root
		tree.root = node
		node.Father = nil
		node.GrandFather = nil
	}

	node.IsBlack = true
	grandFather.IsBlack = false
}


type TreeNode struct {
	Value       int64
	LeftSon     *TreeNode
	RightSon    *TreeNode
	Father      *TreeNode
	GrandFather *TreeNode
	//默认新节点是红色，刚好符合默认值是false
	IsBlack bool
}

type Tree struct {
	root   *TreeNode
	length int
}

func (tree *Tree) LDR(root *TreeNode) {
	if root.LeftSon != nil {
		leftSon := root.LeftSon
		tree.LDR(leftSon)
	}
	var color string
	if root.IsBlack {
		color = "黑"
	} else {
		color = "红"
	}
	fmt.Printf("%v(%v) ->", root.Value, color)
	if root.RightSon != nil {
		rightSon := root.RightSon
		tree.LDR(rightSon)
	}
}

func (tree *Tree) Add(value int64) {
	pendingBalance := tree.doAdd(value)
	tree.Balance(pendingBalance)
}

func (tree *Tree) doAdd(value int64) *TreeNode {
	node := &TreeNode{
		Value: value,
	}
	father := tree.root
	if father == nil {
		tree.root = node
		return node
	}
	for {
		if value >= father.Value {
			if father.RightSon == nil {
				father.RightSon = node
				node.Father = father
				node.GrandFather = father.Father
				break
			} else {
				father = father.RightSon
			}
		} else {
			if father.LeftSon == nil {
				father.LeftSon = node
				node.Father = father
				node.GrandFather = father.Father

				break
			} else {
				father = father.LeftSon
			}
		}
	}
	return node
}
