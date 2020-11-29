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
7.如果父节点为红，叔叔为红，爷爷节点为黑（必然），则父和叔叔变为黑，爷爷节点变为红。
  如爷爷节点为根节点，再转变为黑色，如果爷爷节点不是根节点则不符合规则，此时需要将爷爷节点作为新增节点递归判断
8.如果父节点是红色，且新增节点没有叔叔节点/叔叔节点为黑节点，视情况做上面四种旋转操作
*/

/**
删除操作（不考虑平衡）：
	1.如果删除的节点无叶子节点（非Nil）：直接删除
	2.如果删除的节点有单个叶子节点（非Nil）：直接删除，叶子节点变为自己的颜色，并代替自己挂在父节点上
	3.如果删除的节点有两个叶子节点（非Nil）：找到待删除节点的前继节点/后继节点，删除掉待删除节点，将前继几点/后继节点代替删除节点的位置
		具体操作如下（最省事）：
			将待删除节点和前继节点/后继节点的值互换，最后删除掉被更换的前继节点/后继节点（即通过一些转换，使原本删除拥有左右子树的节点
		的操作变为1.操作）
	总的来说，红黑树的实际删除操作只有1.和2.两种
删除时的平衡规则：
	1.被删除节点为红节点，且没有叶子节点（红色的删除1操作）：
		删除1后不需要做其他操作，直接删除
	2.被删除节点为黑节点，且有单个叶子节点（黑色的删除2操作）：
		删除2之后不需要做其他操作，TODO 注意平衡2情况不存在红节点，因为红黑树规定任一节点到其每个叶子的所有路径都包含相同数目的黑色节点
	3.被删除节点为黑节点，且没有叶子节点（黑色的删除1操作）：
	done	3.1 黑节点是左子节点，其兄弟节点也是黑色，兄弟节点的右子节点是红色：
			父节点与兄弟节点颜色互换，兄弟右子节点变为黑色，兄弟节点以父节点为圆心做左旋操作，删除黑节点
	done	3.2 黑节点是左子节点，其兄弟节点也是黑色，兄弟节点的左子节点是红色：
			兄弟节点和兄弟左子节点颜色互换，兄弟左子节点以兄弟节点为圆心做右旋操作，此时情况会变成3.1，再按照3.1步骤平衡，删除黑节点
			TODO 如果3.1和3.2同时存在，则优先按3.1处理
	done	3.3 黑节点的兄弟节点是黑色，父节点是红色，且兄弟节点也是叶子节点：
			兄弟节点与父节点互换颜色，删除黑节点，
	done	3.4 黑节点是右子节点，其兄弟节点也是黑色，兄弟节点的左子节点是红色：
			父节点与兄弟节点颜色互换，兄弟左子节点变为黑色，兄弟节点以父节点为圆心做右旋操作，删除黑节点
	done	3.5 黑节点是右子节点，其兄弟节点也是黑色，兄弟节点的右子节点是红色：
			兄弟节点和兄弟右子节点颜色互换，兄弟右子节点以兄弟节点为圆心做左旋操作，此时情况会变成3.4，再按照3.4步骤平衡，删除黑节点
			TODO 如果3.4和3.5同时存在，则优先按3.4处理
		3.6	兄弟节点和父亲节点都为黑节点，且兄弟节点没有红子节点：
			兄弟节点设为红，删除黑节点，将父节点作为待删除节点进行递归判断（此时不要管待删除节点是否有子节点了）
			（TODO 注意！递归判断里不能删除待删除节点）
	done	3.7 黑节点是左子节点，兄弟节点是红色：
			将父节点和兄弟节点颜色互换，兄弟节点以父节点为圆心做左旋操作，此时情况会变成3.3，再按3.3步骤操作，删除黑节点
	done	3.8 黑节点是右子节点，兄弟节点是红色:
			将父节点和兄弟节点颜色互换，兄弟节点以父节点为圆心做右旋操作，此时情况会变成3.3，再按3.3步骤操作，删除黑节点



*/

//先调整，最后返回要删除的节点（有可能递归调整但不删除）
func (tree *Tree) balanceBeforeDelete(node *TreeNode) *TreeNode {
	father := node.Father
	var brother *TreeNode
	if father.LeftSon == node {
		brother = father.RightSon
		if brother.RightSon != nil && brother.IsBlack == true && brother.RightSon.IsBlack == false {
			//3.1
			return tree.brotherBlackFarSonRed(node, true)
		}
		if brother.LeftSon != nil && brother.IsBlack == true && brother.LeftSon.IsBlack == false {
			//3.2
			pending := tree.brotherBlackNearSonRed(node, true)
			//3.2 → 3.1
			return tree.brotherBlackFarSonRed(pending, true)
		}
		//3.3
		if brother.IsBlack == true && father.IsBlack == false {
			return tree.brotherBlackFatherRed(node, true)
		}
		//3.7
		if brother.IsBlack == false {
			pending := tree.brotherRed(node, true)
			//3.7 → 3.3
			return tree.brotherBlackFatherRed(pending, true)
		}
		//3.6
		if brother.IsBlack == true && father.IsBlack == true {
			pending := tree.brotherBlackFatherBlack(node, false)
			//直接将父节点作为待删除节点递归调整
			tree.balanceBeforeDelete(pending.Father)
			return pending
		}
	} else {
		brother = father.LeftSon
		if brother.LeftSon != nil && brother.IsBlack == true && brother.LeftSon.IsBlack == false {
			//3.4
			return tree.brotherBlackFarSonRed(node, false)
		}
		if brother.RightSon != nil && brother.IsBlack == true && brother.RightSon.IsBlack == false {
			//3.5
			pending := tree.brotherBlackNearSonRed(node, false)
			//3.5 → 3.4
			return tree.brotherBlackFarSonRed(pending, false)
		}
		//3.3
		if brother.IsBlack == true && father.IsBlack == false {
			return tree.brotherBlackFatherRed(node, false)
		}
		//3.8
		if brother.IsBlack == false {
			pending := tree.brotherRed(node, false)
			// 3.8 → 3.3
			return tree.brotherBlackFatherRed(pending, false)
		}
		//3.6
		if brother.IsBlack == true && father.IsBlack == true {
			pending := tree.brotherBlackFatherBlack(node, false)
			//直接将父节点作为待删除节点递归判断
			tree.balanceBeforeDelete(pending.Father)
			return pending
		}

	}
	return nil
}

//对应3.1、3.4情况
func (tree *Tree) brotherBlackFarSonRed(node *TreeNode, isLeft bool) *TreeNode {
	father := node.Father
	fatherColor := father.IsBlack
	if isLeft {
		//3.1
		brother := father.RightSon
		father.IsBlack = brother.IsBlack
		brother.IsBlack = fatherColor
		brother.RightSon.IsBlack = true
		//兄弟节点绕父节点左旋
		tree.Left(brother.RightSon)
	} else {
		brother := father.LeftSon
		father.IsBlack = brother.IsBlack
		brother.IsBlack = fatherColor
		brother.LeftSon.IsBlack = true
		//兄弟节点绕父节点右旋
		tree.Right(brother.LeftSon)
	}
	return node
}

/**
对应3.2、3.5情况
isLeft:待删除节点是否为左子节点
*/
func (tree *Tree) brotherBlackNearSonRed(node *TreeNode, isLeft bool) *TreeNode {
	father := node.Father
	var brotherColor bool
	//3.2
	if isLeft {
		brother := father.RightSon
		brotherColor = brother.IsBlack
		son := brother.LeftSon
		brother.IsBlack = son.IsBlack
		son.IsBlack = brotherColor
		//这个右旋比较特殊，是直接将左子节点右旋过去，下面写的Right()是基于左子节点的子节点进行判断，在这里使用有可能NPE
		//注意父节点与祖父节点是否对应上了，不然会死循环
		tree.FatherRight(son)
	} else {
		//3.5
		brother := father.LeftSon
		brotherColor = brother.IsBlack
		son := brother.RightSon
		brother.IsBlack = son.IsBlack
		son.IsBlack = brotherColor
		//这个左旋比较特殊，是直接将右子节点右旋过去，下面写的Left()是基于右子节点的子节点进行判断，在这里使用有可能NPE
		//注意父节点与祖父节点是否对应上了，不然会死循环
		tree.FatherLeft(son)
	}
	return node
}

//对应3.3，兄弟黑父亲红
//兄弟节点与父节点互换颜色
func (tree *Tree) brotherBlackFatherRed(node *TreeNode, isLeft bool) *TreeNode {
	var brother *TreeNode
	father := node.Father
	if isLeft {
		brother = father.RightSon
	} else {
		brother = father.LeftSon
	}
	father.IsBlack = true
	brother.IsBlack = false
	return node
}

//对应3.7，3.8
func (tree *Tree) brotherRed(node *TreeNode, isLeft bool) *TreeNode {
	father := node.Father
	var brother *TreeNode
	if isLeft {
		brother = father.RightSon
		brother.IsBlack = father.IsBlack
		father.IsBlack = false
		tree.FatherLeft(brother)
		return node
	} else {
		brother = father.LeftSon
		brother.IsBlack = father.IsBlack
		father.IsBlack = false
		tree.FatherRight(brother)
		return node
	}
}

//对应3.6
func (tree *Tree) brotherBlackFatherBlack(node *TreeNode, isLeft bool) *TreeNode {
	father := node.Father
	var brother *TreeNode
	if isLeft {
		brother = father.RightSon
	} else {
		brother = father.LeftSon
	}
	brother.IsBlack = false
	return node
}

/**
此方法用于删除数据的平衡
*/
func (tree *Tree) deleteBalance(node *TreeNode) {
	//平衡分3种：1.直接删除、2：删除后平衡、3：平衡后删除
	hasLeftKid := node.LeftSon != nil
	hasRightKid := node.RightSon != nil
	if !hasLeftKid && !hasRightKid {
		if !node.IsBlack {
			tree.deleteDirectly(node)
		} else {
			//TODO balanceBeforeDelete
			pending := tree.balanceBeforeDelete(node)
			tree.deleteDirectly(pending)
		}
	} else if hasLeftKid || hasRightKid {
		tree.deleteBeforeBalance(node)
	}
}

//适合情况：删除一个红叶子节点
func (tree *Tree) deleteDirectly(node *TreeNode) {
	father := node.Father
	if father.LeftSon == node {
		father.LeftSon = nil
	} else {
		father.RightSon = nil
	}
	node.Father = nil
}

//适合情况：删除只有单个叶子节点的黑节点
func (tree *Tree) deleteBeforeBalance(node *TreeNode) {
	var son *TreeNode
	father := node.Father
	if node.LeftSon != nil {
		son = node.LeftSon
	} else {
		son = node.RightSon
	}
	if father.LeftSon == node {
		father.LeftSon = son
	} else {
		father.RightSon = son
	}
	son.Father = father
	node.Father = nil
	node.LeftSon = nil
	node.RightSon = nil
	son.IsBlack = true
}

func (tree *Tree) Delete(value int64) {
	pendingBalance := tree.doDelete(value)
	if pendingBalance == nil {
		return
	}
	log.Printf("真正要删除的节点：%v\n", pendingBalance)
	//真正的删除
	tree.deleteBalance(pendingBalance)
}

/**
找到真正要删除的节点
*/
func (tree *Tree) doDelete(value int64) *TreeNode {
	//maybe nil
	deletedNode := tree.Get(value)
	if deletedNode == nil {
		return nil
	}
	hasLeftKid := deletedNode.LeftSon != nil
	hasRightKid := deletedNode.RightSon != nil
	if hasLeftKid && hasRightKid {
		//寻找前继节点和后继 节点（优先找红色的）
		pendingBalance := getLessAndBigger(deletedNode)
		log.Printf("找到了颜色为%v的替代节点，值是%v\n", pendingBalance.IsBlack, pendingBalance.Value)
		//互换值
		temp := pendingBalance.Value
		pendingBalance.Value = deletedNode.Value
		deletedNode.Value = temp
		return pendingBalance
	}
	return deletedNode
}

func getLessAndBigger(node *TreeNode) *TreeNode {
	less := node.LeftSon
	bigger := node.RightSon
	for less.RightSon != nil {
		less = less.RightSon
	}
	for bigger.LeftSon != nil {
		bigger = bigger.LeftSon
	}
	//只要其中一个不是红色，直接返回另一个就好了
	if less.IsBlack == false {
		return less
	} else {
		return bigger
	}
}

func (tree *Tree) Get(value int64) *TreeNode {
	root := tree.root
	for {
		if root == nil {
			return nil
		} else if root.Value == value {
			return root
		} else if root.Value < value {
			//应该在右节点寻找
			root = root.RightSon
		} else if root.Value > value {
			root = root.LeftSon
		}
	}
}

/**
此方法用于插入数据后平衡，需要在方法体内判断用哪种平衡方式
*/
func (tree *Tree) addBalance(node *TreeNode) {
	var L bool = false
	var R bool = false
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
		if node.Value > father.Value {
			LR = true
		} else {
			R = true
		}
	} else if grandFather != nil && grandFather.RightSon == father {
		uncle = grandFather.LeftSon
		if node.Value > father.Value {
			L = true
		} else {
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
		tree.addBalance(grandFather)
		return
	}
	//8.
	//记得，进行旋转的大前提是有祖父
	if grandFather != nil && father.IsBlack == false && (uncle == nil || uncle.IsBlack == true) {
		log.Println("准备旋转")
		if L {
			log.Println("进入了L模式")
			tree.Left(node)
			return
		} else if R {
			log.Println("进入了R模式")
			tree.Right(node)
			return
		} else if LR {
			log.Println("进入了LR模式")
			tree.LeftRight(node)
			return
		} else if RL {
			log.Println("进入了RL模式")
			tree.RightLeft(node)
			return
		}
	}

	//如果父节点为黑色，直接插入
	log.Println("父节点为黑，直接插入")
}

//以插入节点为主观视觉进行父节点左旋
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
	if brother != nil {
		brother.Father = grandFather
	}

	if greatGrandFather != nil && greatGrandFather.LeftSon == grandFather {
		greatGrandFather.LeftSon = father
	} else if greatGrandFather != nil && greatGrandFather.RightSon == grandFather {
		greatGrandFather.RightSon = father
	} else if greatGrandFather == nil {
		//说明grandFather是root，需要将root指针重新指定
		tree.root = father
	}
	father.Father = greatGrandFather
	//插入node后再进行balance，要注意node的grandFather也要变
	node.GrandFather = father.Father

	father.IsBlack = true
	grandFather.IsBlack = false

}

//以插入节点为主观视觉进行父节点右旋
func (tree *Tree) Right(node *TreeNode) {
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
	if brother != nil {
		brother.Father = grandFather
	}

	if greatGrandFather != nil && greatGrandFather.LeftSon == grandFather {
		greatGrandFather.LeftSon = father
	} else if greatGrandFather != nil && greatGrandFather.RightSon == grandFather {
		greatGrandFather.RightSon = father
	} else if greatGrandFather == nil {
		//说明grandFather是root，需要将root指针重新指定
		tree.root = father
	}
	father.Father = greatGrandFather
	//插入node后再进行balance，要注意node的grandFather也要变
	node.GrandFather = father.Father

	father.IsBlack = true
	grandFather.IsBlack = false
}

//以旋转节点为主观视觉进行自身左旋
func (tree *Tree) FatherLeft(node *TreeNode) {
	//左旋，说明是右子节点
	father := node.Father
	brother := father.LeftSon
	grandFather := node.GrandFather
	leftSon := node.LeftSon
	rightSon := node.RightSon
	//圆点：node 圆心：father
	node.LeftSon = father
	father.Father = node
	father.GrandFather = grandFather
	father.RightSon = leftSon

	node.Father = grandFather
	grandFather.LeftSon = node
	node.GrandFather = grandFather.Father

	if brother != nil {
		brother.GrandFather = node
	}
	if rightSon != nil {
		rightSon.GrandFather = grandFather
	}
	if leftSon != nil {
		leftSon.Father = father
		leftSon.GrandFather = node
	}
}

//以旋转节点为主观视觉进行自身右旋
func (tree *Tree) FatherRight(node *TreeNode) {
	//右旋，说明是左子节点
	father := node.Father
	brother := father.RightSon
	grandFather := node.GrandFather
	leftSon := node.LeftSon
	rightSon := node.RightSon
	//圆点：node 圆心：father
	node.RightSon = father
	father.Father = node
	father.GrandFather = grandFather
	father.LeftSon = rightSon

	node.Father = grandFather
	grandFather.RightSon = node
	node.GrandFather = grandFather.Father

	if brother != nil {
		brother.GrandFather = node
	}
	if leftSon != nil {
		leftSon.GrandFather = grandFather
	}
	if rightSon != nil {
		rightSon.Father = father
		rightSon.GrandFather = node
	}
}

func (tree *Tree) LeftRight(node *TreeNode) {
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

	if greatGrandFather != nil && greatGrandFather.LeftSon == grandFather {
		greatGrandFather.LeftSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	} else if greatGrandFather != nil && greatGrandFather.RightSon == grandFather {
		greatGrandFather.RightSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	} else if greatGrandFather == nil {
		//说明grandFather是root，需要将root指针重新指定
		//TODO 很奇怪，如果这里传的是root的指针会无效，必须要传tree.root
		tree.root = node
		node.Father = nil
		node.GrandFather = nil
	}

	node.IsBlack = true
	grandFather.IsBlack = false

}

func (tree *Tree) RightLeft(node *TreeNode) {
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

	if greatGrandFather != nil && greatGrandFather.LeftSon == grandFather {
		greatGrandFather.LeftSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	} else if greatGrandFather != nil && greatGrandFather.RightSon == grandFather {
		greatGrandFather.RightSon = node
		node.Father = greatGrandFather
		//插入node后再进行balance，要注意node的grandFather也要变
		node.GrandFather = greatGrandFather.Father
	} else if greatGrandFather == nil {
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
	tree.addBalance(pendingBalance)
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
