package main

import "fmt"

// 链栈
type StackNode struct {
    Data interface{} // 数据
    Next *StackNode  // 下一个节点
}

// 创建链栈
func CreateStack(data ...interface{}) *StackNode {
    if len(data) == 0 {
        return nil
    }
    s := new(StackNode)

    // 记录下一个节点
    var nextNode *StackNode = nil
    for _, v := range data {
        // 创建新节点存储数据
        newNode := new(StackNode)
        newNode.Data = v
        s = newNode

        s.Next = nextNode
        // 下一个节点为当前节点
        nextNode = s
    }
    return s
}

// 打印链栈
func PrintStack(s *StackNode) {
    if s == nil {
        return
    }
    for s != nil {
        fmt.Print(s.Data, " ")
        s = s.Next
    }
    fmt.Println()
}

// 栈链个数
func GetLen(s *StackNode) int {
    if s == nil {
        return -1
    }

    i := 0
    // 循环计算链栈个数
    for s != nil {
        i++
        s = s.Next
    }
    return i
}

// 入栈
func Push(s *StackNode, data interface{}) *StackNode {
    if s == nil || data == nil {
        return nil
    }
    // 新建节点
    newNode := new(StackNode)
    newNode.Data = data
    newNode.Next = s

    return newNode
}

// 出栈
func Pop(s *StackNode) *StackNode {
    if s == nil {
        return nil
    }
    // 下一个节点
    nextNode := s.Next
    // 当前节点清空
    s.Next = nil

    return nextNode
}

func Clear(s *StackNode) *StackNode  {

    for s != nil {
        s = Pop(s)
    }

    return nil
}