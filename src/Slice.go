package main

/*
#include <Stdlib.h>
*/
import "C"
import (
    "unsafe"
    "fmt"
)

// 导入c语言库，上面的方式要连在一起不要有空格
// 自定义切片
type Slice struct {
    Data unsafe.Pointer // 万能指针类型，对应C语言中的void*
    len  int            // 有效长度
    cap  int            // 有效的容量
}

const TAG = 8

// 创建切片(长度 容量 数据)
func (s *Slice) Create(l, c int, data ...int) {
    // 如果数据为空，返回
    if len(data) == 0 {
        return
    }

    // 长度小于0 容量小于0 长度大于容量 数据长度大于长度
    if l < 0 || c < 0 || l > c || len(data) > l {
        return
    }

    // ulonglong unsigned long long  无符号的长长整型
    // 通过C语言代码开辟空间 存储数据
    // 如果堆空间开辟失败 返回值为NULL 相当于nil 内存地址编号为0的空间
    s.Data = C.malloc(C.ulonglong(c) * 8)
    s.len = l
    s.cap = c

    // 转成可以计算的指针类型
    p := uintptr(s.Data)
    // 循环添加数据存储
    for _, v := range data {
        // 数据存储
        *(*int)(unsafe.Pointer(p)) = v
        // 指针偏移
        p += TAG
    }
}

// 打印切片
func (s *Slice) Print() {
    if s == nil {
        return
    }
    p := uintptr(s.Data)
    // 将万能指针转成可以计算的指针
    for i := 0; i < s.len; i++ {
        // 获取内存中的数据
        fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
        p += TAG
    }
}

// 切片追加
func (s *Slice) Append(data ...int) {
    if s == nil || len(data) == 0 {
        return
    }

    if s.len+len(data) > s.cap {
        // 扩充容量
        // C.realloc(指针,字节大小)
        s.Data = C.realloc(s.Data, C.ulonglong(s.cap)*2*8)
        // 改变容量的值
        s.cap = s.cap * 2
    }

    p := uintptr(s.Data)

    for i := 0; i < s.len; i++ {
        // 指针偏移
        p += TAG
    }

    // 添加数据
    for _, v := range data {
        *(*int)(unsafe.Pointer(p)) = v
        p += TAG
    }
    // 更新有效数据（长度）
    s.len = s.len + len(data)
}

// 获取元素 GetData(下标)  返回值为int 元素
func (s *Slice) GetData(index int) int {
    if s == nil {
        return -1
    }
    if index < 0 || index > s.len-1 {
        return -1
    }

    p := uintptr(s.Data)
    for i := 0; i < index; i++ {
        // 指针偏移
        p += TAG
    }
    return *(*int)(unsafe.Pointer(p))
}

// 查找元素 Search(元素)返回值为int 下标
func (s *Slice) Search(data int) int {
    if s == nil || s.Data == nil {
        return -1
    }
    p := uintptr(s.Data)

    for i := 0; i < s.len; i++ {
        // 查找数据  返回第一次元素出现的位置
        if *(*int)(unsafe.Pointer(p)) == data {
            return i
        }

        // 指针偏移
        p += TAG
    }
    return -1
}

// 删除元素 Delete(下标)
func (s *Slice) Delete(index int) {
    if s == nil || s.Data == nil {
        return
    }

    if index < 0 || index > s.len-1 {
        return
    }
    // 将指针指向需要删除的下标位置
    p := uintptr(s.Data)
    for i := 0; i < index; i++ {
        p += TAG
    }
    // 用下一个指针对应的值为当前指针对应的值进行赋值
    temp := p
    for i := index; i < s.len; i++ {
        temp += TAG
        *(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
        p += TAG
    }
    s.len --
}

// 插入元素 Insert(下标 元素)
func (s *Slice) Insert(index, data int) {
    if s == nil || s.Data == nil {
        return
    }

    if index < 0 || index > s.len {
        return
    }
    // 如果插入数据是最后一个，调用追加方法
    if index == s.len-1 {
        s.Append(data)
        return
    }

    // 获取插入数据的位置
    p := uintptr(s.Data)
    for i := 0; i < index; i++ {
        p += TAG
    }

    // 获取末尾的指针位置
    temp := uintptr(s.Data)
    temp += TAG * uintptr(s.len)
    for i := s.len; i > index; i-- {
        // 用前一个数据为当前数据赋值
        *(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
        temp -= TAG
    }

    // 修改插入下标的数据
    *(*int)(unsafe.Pointer(p)) = data
    s.len++
}

// 销毁切片
func (s *Slice) Destroy() {
    // 调用C语言  适释放堆空间
    C.free(s.Data)
    s.Data = nil
    s.len = 0
    s.cap = 0
    s = nil
}
