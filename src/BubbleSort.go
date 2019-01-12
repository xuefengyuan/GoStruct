package main

func BubbleSort(arr []int) {
    // 冒泡排序优化
    flag := false
    // 外层控制行
    for i := 0; i < len(arr)-1; i++ {
        // 内层控制列
        for j := 0; j < len(arr)-1-i; j++ {
            // 两个值相比较
            if arr[j] > arr[j+1] {
                // 交换数据
                arr[j], arr[j+1] = arr[j+1], arr[j]
                flag = true
            }
        }
        // 判断数据是否是有序
        if !flag {
            return
        } else {
            flag = false
        }
    }
}

// 冒泡排序，优化前
func BubbleSort1(arr []int) {
    // 外层控制行
    for i := 0; i < len(arr)-1; i++ {
        // 内层控制列
        for j := 0; j < len(arr)-1-i; j++ {
            // 两个值相比较
            if arr[j] > arr[j+1] {
                // 交换数据
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
