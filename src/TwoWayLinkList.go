package main

import "fmt"

// 双向链表
type BothwayLinkNode struct {
    Data interface{}      // 数据
    Prev *BothwayLinkNode // 上一个指针节点
    Next *BothwayLinkNode // 下一个指针节点
}

// 创建双向链表 Create(数据集合)
func (node *BothwayLinkNode) Create(data ...interface{}) {
    if node == nil || len(data) == 0 {
        return
    }
    // 记录头节点
    head := node
    for _, v := range data {
        // 创建新节点
        newNode := new(BothwayLinkNode)
        newNode.Data = v
        // 新节点，指向上一个节点
        newNode.Prev = node
        // 当前节点的下一个节点是新节点
        node.Next = newNode
        // 当前节点为下一个节点
        node = node.Next
    }

    node = head
}

// 打印双向链表
func (node *BothwayLinkNode) Print() {
    if node == nil {
        return
    }
    for node != nil {
        if node.Data != nil {
            fmt.Print(node.Data, " ")
        }
        node = node.Next
    }
}

// 打印双向链表 倒序
func (node *BothwayLinkNode) PrintTwo() {
    if node == nil {
        return
    }

    // 指向链表末尾
    for node.Next != nil {
        node = node.Next
    }

    // 从后向前打印数据
    for node.Prev != nil {
        if node.Data != nil {
            fmt.Print(node.Data, " ")
        }
        node = node.Prev
    }
}

// 数据长度 返回值 个数
func (node *BothwayLinkNode) GetLen() int {
    if node == nil {
        return -1
    }
    i := 0
    for node.Next != nil {
        i++
        node = node.Next
    }
    return i
}

// 插入数据 （下标 数据）Insert
func (node *BothwayLinkNode) InsertByIndex(index int, data interface{}) {
    if node == nil || index < 0 || data == nil {
        return
    }
    // 记录上一个节点
    preNode := node
    // 循环找到插入的节点
    for i := 0; i < index; i++ {
        preNode = node
        if node == nil {
            return
        }
        node = node.Next
    }
    // 创建新节点
    newNode := new(BothwayLinkNode)
    // 将新节点的指针域分别指向上一个节点和下一个节点
    newNode.Data = data
    newNode.Next = node
    newNode.Prev = preNode

    // 上一个节点的下一个节点为新节点
    preNode.Next = newNode
    // 下一个节点的上一个节点为新节点
    node.Prev = newNode
}

func (node *BothwayLinkNode) Delete(index int) {

    if node == nil || index < 0 {
        return
    }

    preNode := node
    for i := 0; i < index; i++ {
        preNode = node
        if node == nil {
            return
        }
        node = node.Next
    }
    // 删除节点
    preNode.Next = node.Next
    // 当前节点的下一个节点的上一个节点为上一个节点
    node.Next.Prev = preNode

    // 销毁当前节点
    node.Data = nil
    node.Next = nil
    node.Prev = nil
    node = nil
}

// 链表销毁
func (node *BothwayLinkNode) Destroy() {
    if node == nil {
        return
    }
    // 递归调用销毁链表
    node.Next.Destroy()

    node.Data = nil
    node.Next = nil
    node.Prev = nil
    node = nil
}
