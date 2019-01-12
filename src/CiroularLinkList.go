package main

import (
    "fmt"
)

// 循环链表
type CirLinkList struct {
    Data interface{}
    Next *CirLinkList
}

// 创建循环链表
func (node *CirLinkList) Create(data ...interface{}) {
    if node == nil || len(data) == 0 {
        return
    }
    head := node
    for _, v := range data {
        // 创建一个新节点
        newNode := new(CirLinkList)
        newNode.Data = v

        // 当前节点的下一个节点为新节点
        node.Next = newNode
        node = node.Next
    }

    // 形成闭环
    // 最后一个节点的下一个节点 为第一个节点
    node.Next = head.Next
    node = head
}

// 打印链表
func (node *CirLinkList) Print() {
    if node == nil {
        return
    }
    // 记录循环链表的起始点
    start := node.Next
    // 打印链表
    for {
        node = node.Next

        if node.Data != nil {
            fmt.Print(node.Data, " ")
        }
        if start == node.Next {
            break
        }
    }
}

// 循环链表长度 返回值 个数
func (node *CirLinkList) GetLen() int {
    if node == nil {
        return -1
    }
    // 记录循环链表的起始点
    start := node.Next

    // 定义计数器
    i := 0;
    for {
        node = node.Next
        i++
        if start == node.Next {
            break
        }
    }
    return i
}

// 插入数据（下标 数据）
func (node *CirLinkList) Insert(index int, data interface{}) {
    // 判断位置是否越界
    if node == nil || index < 0 || node.GetLen()+1 < index || data == nil {
        return
    }
    // 记录链表第一个节点
    start := node.Next

    // 记录上一个节点
    preNode := node
    for i := 0; i < index; i++ {
        preNode = node
        node = node.Next
    }
    // 创建新节点
    newNode := new(CirLinkList)
    newNode.Data = data
    // 新节点的下一个节点为node
    newNode.Next = node
    // 上一个节点的下一个节点为新节点
    preNode.Next = newNode

    // 判断如果是第一个节点需要将最后一个节点指向新节点
    if index == 1 {
        for {
            // 判断最后一个节点
            if start == node.Next {
                break
            }
            node = node.Next
        }
        // 新节点为最后一个节点的下一个节点
        node.Next = newNode
    }
}

// 根据下标删除节点
func (node *CirLinkList) Delete(index int) {
    if node == nil || index < 0 {
        return
    }
    // 起始节点
    start := node.Next
    // 记录上一个节点
    preNode := node
    // 循环找到删除数据的位置
    for i := 0; i < index; i++ {
        preNode = node
        node = node.Next
    }

    // 判断如果删除的是第一个节点
    if index == 1 {
        // temp记录最后一个位置
        temp := start
        for {
            if start == temp.Next {
                break
            }
            // 将最后一个节点指向新节点
            temp = temp.Next
        }
        temp.Next = node.Next
    }

    // 删除数据
    preNode.Next = node.Next
    node = nil
}

// 约瑟夫环
func JosephRing() {
    list := new(CirLinkList)
    list.Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
        17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32)

    fmt.Println("原始数据：")
    list.Print()
    fmt.Println()
    fmt.Println("删除数据：")
    i := 1
    for list.GetLen() > 2 {
        i += 3
        if i > list.GetLen() {
            i = list.GetLen()%3 + 1
        }
        list.Delete(i)
        list.Print()

        fmt.Println("\ni = ", i)

    }
}
