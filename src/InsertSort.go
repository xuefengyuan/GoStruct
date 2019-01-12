package main

import (
    "math/rand"
    "time"
    "fmt"
)

// 插入排序
func InsertSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        // 如果当前数据小于有序数据
        if arr[i] < arr[i-1] {

            j := i - 1
            // 获取有效数据
            temp := arr[i]

            // 依次比较有序数据
            for j >= 0 && arr[j] > temp {
                arr[j+1] = arr[j]
                j--
            }
            arr[j+1] = temp
        }
    }
}

// 变相排序  基于大量重复 在某一个范围内
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
