package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //testSlice()
    //testLinkNode()
    //testBothwayLinkNode()
    //testCirLinkNode()
    //JosephRing()
    //testStackLink()
    //testQueueLink()
    //testBinaryNode()
    testSort()
    //testInsertSort()
}

func testSlice() {
    var s Slice

    s.Create(5, 5, 1, 2, 3, 4, 5, )
    s.Print()
    fmt.Println()
    s.Append(6, 7, 8, 9, 10, 11)
    s.Print()
    fmt.Println()
    fmt.Println("根据下标查找：", s.GetData(3))
    fmt.Println("数据所在下标：", s.Search(8))
    s.Delete(3)
    s.Print()
    fmt.Println()
    s.Insert(0, 42)
    s.Print()
    fmt.Println()
    s.Destroy()
    fmt.Println(s)
}

func testLinkNode() {
    var node LinkNode
    node.Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
    node.Print()
    fmt.Println()

    len := node.GetLen()
    fmt.Println(len)

    node.InsertByHead(21)
    node.Print()
    fmt.Println()

    node.InsertByTail(22)
    node.Print()
    fmt.Println()

    node.InsertByIndex(8, 33)
    node.Print()
    fmt.Println()

    node.DeleteByIndex(3)
    node.Print()
    fmt.Println()

    node.DeleteByData(5)
    node.Print()
    fmt.Println()

    fmt.Println(node.Search(34))
}

func testBothwayLinkNode() {
    var node BothwayLinkNode
    node.Create("a", "b", "c", "d")
    node.Print()
    fmt.Println()

    fmt.Println(node.GetLen())
    node.PrintTwo()

    node.InsertByIndex(2, "f")
    fmt.Println()
    node.Print()

    node.Delete(3)
    fmt.Println()
    node.Print()
}

func testCirLinkNode() {
    var node CirLinkList
    node.Create("A", "B", "C", "D", "E", "F", "G")
    node.Print()
    fmt.Println()

    fmt.Println(node.GetLen())
    node.Insert(4, "M")
    node.Print()
    fmt.Println()
    node.Delete(6)
    node.Print()

}

func testStackLink() {
    s := CreateStack(1, 2, 3, 4, 5, 6)
    PrintStack(s)
    fmt.Println(GetLen(s))

    //入栈
    s = Push(s, 7)
    s = Push(s, 8)
    PrintStack(s)

    //出栈
    s = Pop(s)
    fmt.Println("\n出栈后：")
    PrintStack(s)

    //出栈
    s = Pop(s)
    fmt.Println("\n出栈后：")
    PrintStack(s)

    // 清空链栈
    s = Clear(s)
    fmt.Println("\n清空栈：")
    PrintStack(s)
}

func testQueueLink() {
    var queue QueueNode
    queue.Create("A", "B", "C", "D", "E")
    queue.Print()
    queue.Push("G")
    queue.Print()
    queue.Push("M")
    queue.Print()
    queue.Pop()
    queue.Print()
    queue.Pop()
    queue.Print()
}

func testBinaryNode() {
    // 创建树
    tree := new(BinaryNode)

    // 创建节点
    node1 := BinaryNode{1, nil, nil}
    node2 := BinaryNode{2, nil, nil}
    node3 := BinaryNode{3, nil, nil}
    node4 := BinaryNode{4, nil, nil}
    node5 := BinaryNode{5, nil, nil}
    node6 := BinaryNode{6, nil, nil}
    //node7 := BinaryNode{7, nil, nil}

    /*
                0
            1       2
         3     4 5     6
      7
    */
    // 树创建节点
    tree.Data = 0
    tree.LChild = &node1
    tree.RChile = &node2

    node1.LChild = &node3
    node1.RChile = &node4

    node2.LChild = &node5
    node2.RChile = &node6

    //node3.LChild = &node7

    // 先序遍历
    tree.PreOrder()
    fmt.Println()
    // 中序遍历
    tree.MidOrder()
    fmt.Println()
    // 后序遍历
    tree.RearOrder()
    fmt.Println()

    fmt.Println("树的高度：", tree.TreeHeight())

    num := 0
    tree.LeafCount(&num)
    fmt.Println("叶子节点个数：", num)

    tree.Search(5)

    //tree.Destroy()
    //fmt.Println(tree)
    fmt.Println("反转前：")
    tree.PreOrder()

    // 如果想反转二叉树 二叉树必须是一个满二叉树
    tree.Reverse()

    fmt.Println("\n反转后：")
    tree.PreOrder()

    tree1 := tree.Copy()
    fmt.Println("新树内容：")
    // 修改新树的值，不会影响旧树
    tree1.LChild.Data = 777
    tree1.PreOrder()
}

func testSort() {
    //arr := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}
    //BubbleSort(arr)
    //fmt.Println("冒泡排序：",arr)

    //SelectSort(arr)
    //fmt.Println("选择排序：",arr)

    //InsertSort(arr)
    //fmt.Println("插入排序：", arr)

    //ShellSort(arr)
    //fmt.Println("希尔排序：", arr)

    //QuickSort(arr,0,len(arr)-1)
    //fmt.Println("快速排序：", arr)

    //HeapInit(arr)
    //fmt.Println("堆排序：", arr)

    // 二分查找
    arr := []int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,18,19,21,23,25,28}
    num := 16
    index := BinarySearch(arr, num)
    fmt.Println("二分查找到的下标：", index)
}

// 变相排序  基于大量重复 在某一个范围内
func testInsertSort() {
    // 随机数种子
    rand.Seed(time.Now().UnixNano())
    s := make([]int, 0)
    for i := 0; i < 10000; i++ {
        s = append(s, rand.Intn(1000))
    }

    // 统计数据集合中数据出现的次数
    m := make(map[int]int)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }

    fmt.Println(m)
    // 排序
    for i := 0; i < 1000; i++ {
        for j := 0; j < m[i]; j++ {
            fmt.Print(i, " ")
        }
    }
}
