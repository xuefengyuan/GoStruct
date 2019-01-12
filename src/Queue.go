package main

import "fmt"

// 链队列
type QueueNode struct {
    Data interface{}
    Next *QueueNode
}

// 创建链列（数据）
func (queue *QueueNode) Create(data ...interface{}) {
    if queue == nil || len(data) == 0 {
        return
    }

    // 创建链列
    for _, v := range data {
        newNode := new(QueueNode)
        newNode.Data = v

        queue.Next = newNode
        queue = queue.Next
    }
}

// 打印链列
func (queue *QueueNode) Print() {
    if queue == nil {
        return
    }
    for queue != nil {
        if queue.Data != nil {
            fmt.Print(queue.Data, " ")
        }
        queue = queue.Next
    }
    fmt.Println()
}

// 链列个数
func (queue *QueueNode) GetLen() int {
    if queue == nil {
        return -1
    }
    i := 0
    for queue != nil {
        i++
        queue = queue.Next
    }
    return i
}

// 入列(insert)
func (queue *QueueNode) Push(data interface{}) {
    if queue == nil || data == nil {
        return
    }
    // 找到队列末尾
    for queue.Next != nil {
        queue = queue.Next
    }
    // 创建新节点 将新节点加入队列末尾
    newNode := new(QueueNode)
    newNode.Data = data

    queue.Next = newNode
}

// 出列
func (queue *QueueNode) Pop() {
    if queue == nil {
        return
    }
    // 队头出列
    queue.Next = queue.Next.Next
}
