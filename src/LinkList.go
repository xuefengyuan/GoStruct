package main

import (
    "fmt"
    "reflect"
)

// 单向链表
// 通过结构体嵌套本结构体指针来实现链表
type LinkNode struct {
    Data interface{}
    Next *LinkNode
}

// 创建链表 Create(数据)
func (node *LinkNode) Create(data ...interface{}) {
    if node == nil {
        return
    }
    // 头节点
    head := node

    for i := 0; i < len(data); i++ {
        // 创建一个新的节点
        newNode := new(LinkNode)
        newNode.Data = data[i]
        newNode.Next = nil
        // 将新节点作为当前节点的下一个节点
        node.Next = newNode
        node = node.Next
    }
    node = head
}

// 打印链表
func (node *LinkNode) Print() {
    if node == nil {
        return
    }
    // 打印数据
    if node.Data != nil {
        fmt.Print(node.Data, " ")
    }
    // 使用递归遍历下一个数据
    node.Next.Print()
}

// 获取链表长度
func (node *LinkNode) GetLen() int {
    if node == nil {
        return -1
    }

    i := 0
    // 遍历查找下一个节点是否为nil
    for node.Next != nil {
        i++
        node = node.Next
    }
    return i
}

// 头部插入数据
func (node *LinkNode) InsertByHead(data interface{}) {
    if node == nil || data == nil {
        return
    }
    // 创建新节点
    newNode := new(LinkNode)
    // 新节点赋值
    newNode.Data = data
    newNode.Next = node.Next
    //  将新节点放在当前节点后面
    node.Next = newNode
}

// 尾部插入数据
func (node *LinkNode) InsertByTail(data interface{}) {

    if node == nil || data == nil {
        return
    }
    // 查找链表的末尾位置
    for node.Next != nil {
        node = node.Next
    }

    // 创建新节点
    newNode := new(LinkNode)
    // 新节点赋值
    newNode.Data = data
    newNode.Next = nil
    // 将新节点放在链表末尾
    node.Next = newNode
}

// 下标插入数据
func (node *LinkNode) InsertByIndex(index int, data interface{}) {
    if node == nil {
        return
    }
    if index < 0 || node.GetLen() < index {
        return
    }
    // 记录上一个节点
    preNode := node

    for i := 0; i < index; i++ {
        preNode = node
        // 如果超出链表个数 直接返回
        if node == nil {
            return
        }
        node = node.Next
    }
    // 创建一个新节点
    newNode := new(LinkNode)
    newNode.Data = data
    newNode.Next = node
    // 上一个节点链接当前节点
    preNode.Next = newNode
}

// 根据下标删除数据
func (node *LinkNode) DeleteByIndex(index int) {
    if node == nil || index < 0 {
        return
    }
    // 记录上一个链表节点
    preNode := node
    for i := 0; i < index; i++ {
        preNode = node
        if node == nil {
            return
        }
        node = node.Next
    }
    // 将上一个节点指针域结点指向node的下一个节点
    preNode.Next = node.Next
    // 销毁当前节点
    node.Data = nil
    node.Next = nil
    node = nil
}

// 删除节点数据
func (node *LinkNode) DeleteByData(data interface{}) {
    if node == nil || data == nil {
        return
    }
    preNode := node

    for node.Next != nil {
        preNode = node
        node = node.Next
        //判断interface存储的数据类型是否相同
        //reflect.DeepEqual()
        if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && node.Data == data {
            preNode.Next = node.Next
            // 销毁数据
            node.Data = nil
            node.Next = nil
            node = nil

            // 如果添加return 表示删除第一个相同的数据
            // 如果不添加return 表示删除所有相同的数据
            return
        }
    }
}

// 查找数据下标
func (node *LinkNode) Search(data interface{}) int {
    if node == nil || data == nil {
        return -1
    }

    i := 0

    for node.Next != nil {
        i++
        // 比较两个接口中的内容是否相同
        // reflect.DeepEqual()
        if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && node.Data == data {
            return i - 1
        }
        node = node.Next
    }

    return -1
}

// 销毁链表
func (node *LinkNode)Destroy()  {
    if node == nil {
        return
    }
    // 通过递归毁销毁链表
    node.Next.Destroy()

    node.Data = nil
    node.Next = nil
    node = nil
}