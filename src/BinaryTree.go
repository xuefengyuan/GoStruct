package main

import (
    "fmt"
    "reflect"
)

// 二叉树
type BinaryNode struct {
    // 数据
    Data interface{}
    // 左子树
    LChild *BinaryNode
    // 右子树
    RChile *BinaryNode
}

// 创建二叉树
func (node *BinaryNode) Create() {
    node = new(BinaryNode)
}

// DLR — 先序遍历，即先根再左再右
func (node *BinaryNode) PreOrder() {
    if node == nil {
        return
    }

    fmt.Print(node.Data, " ")
    // 递归遍历左子树
    node.LChild.PreOrder()
    // 递归遍历右子树
    node.RChile.PreOrder()
}

// LDR — 中序遍历，即先左再根再右
func (node *BinaryNode) MidOrder() {
    if node == nil {
        return
    }
    // 递归遍历左子树
    node.LChild.MidOrder()

    fmt.Print(node.Data, " ")

    // 递归遍历右子树
    node.RChile.MidOrder()
}

// LRD — 后序遍历，即先左再右再根
func (node *BinaryNode) RearOrder() {
    if node == nil {
        return
    }
    // 递归遍历左子树
    node.LChild.RearOrder()
    // 递归遍历右子树
    node.RChile.RearOrder()

    fmt.Print(node.Data, " ")
}

// 二叉树高度 深度
func (node *BinaryNode) TreeHeight() int {
    if node == nil {
        // 这里不能返回0，不然数量会减1个
        return 0
    }
    // 进入下一层遍历
    lh := node.LChild.TreeHeight()
    rh := node.RChile.TreeHeight()
    if lh > rh {
        lh ++
        return lh
    } else {
        rh ++
        return rh
    }
}

// 二叉树叶子节点个数
// 叶子节点是没有后继的节点
func (node *BinaryNode) LeafCount(num *int) {
    if node == nil {
        return
    }
    // 判断没有后继的节点为叶子节点
    if node.LChild == nil && node.RChile == nil {
        (*num)++
    }

    node.LChild.LeafCount(num)
    node.RChile.LeafCount(num)
}

// 二叉树数据查找
func (node *BinaryNode) Search(data interface{}) {

    if node == nil {
        return
    }

    // == 不支持slice和map比较
    // 也可以用 reflect.DeepEqual() 比较
    // 通过反射确认类型比较
    if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && node.Data == data {
        fmt.Println("找到了数据：", node.Data)
        return
    }

    // 左子节点递归查找
    node.LChild.Search(data)
    // 右子节点递归查找
    node.RChile.Search(data)
}

// 二叉树销毁
func (node *BinaryNode) Destroy() {
    if node == nil {
        return
    }

    // 左子节点递归销毁
    node.LChild.Destroy()
    node.LChild = nil
    // 右子节点递归销毁
    node.RChile.Destroy()
    node.RChile = nil
    node.Data = nil
    node= nil
}

// 二叉树反转，如果想反转二叉树 二叉树必须是一个满二叉树
func (node *BinaryNode)Reverse()  {
    if node == nil {
        return
    }
    // 交换节点  多重赋值
    node.LChild,node.RChile = node.RChile,node.LChild

    // 递归调用反转
    node.LChild.Reverse()
    node.RChile.Reverse()
}

// 二叉树拷贝
func (node *BinaryNode)Copy() *BinaryNode  {
    if node == nil {
        return nil
    }

    // 递归copy
    lChild := node.LChild.Copy()
    rChile := node.RChile.Copy()

    // 创建新节点
    newNode := new(BinaryNode)
    newNode.Data = node.Data

    newNode.LChild= lChild
    newNode.RChile = rChile

    return newNode

    return nil
}