# GoStruct
Go的数据结构与算法，后续看情况补充

有些内容不是很完全，后面有时间再补充了

##### Slice：自定义切片，跟C语言交互

##### LinkList：单向链表

##### TwoWayLinkList：双向链表

##### CiroularLinkList：循环链表

##### Stack：栈链表

##### Queue：队列链表

##### BinaryTree：二叉树

```go
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

    fmt.Println("树的高度：",tree.TreeHeight())

    num := 0
    tree.LeafCount(&num)
    fmt.Println("叶子节点个数：",num)

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
```

##### BubbleSort：冒泡排序

- 冒泡排序优化

##### SelectSort：选择排序

##### InsertSort：插入排序

变相排序  基于大量重复 在某一个范围内

```go
func InsertSort1() {
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

    //fmt.Println(m)
    // 排序
    for i := 0; i < 1000; i++ {
        for j := 0; j < m[i]; j++ {
            fmt.Print(i, " ")
        }
    }
}
```

##### ShellSort：希尔排序

##### QuickSort：快速排序

##### HeapInit：堆排序

##### BinarySearch：二分查找

注意：二分查找的数据，必需要是有序的















