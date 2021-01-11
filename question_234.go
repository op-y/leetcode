package leetcode

import (
    "fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    // 空链表认为满足回文条件
    if head == nil {
        return true
    }

    length := 0
    cursor := head
    // 获取链表长度
    for cursor != nil {
        length++
        cursor = cursor.Next
    }
    fmt.Printf("length of slice: %d\n", length)

    if length <= 1 {
        return true
    }

    // 获取头指针
    h1 := head
    h2 := head
    cursor = head
    count := 0
    for cursor != nil {
        if count == length / 2 - 1 {
            h1 = cursor
            fmt.Printf("h1 of slice: %d\n", count)
        }
        if count == length / 2 + length % 2 {
            h2 = cursor
            fmt.Printf("h2 of slice: %d\n", count)
            break
        }
        count++
        cursor = cursor.Next
    }

    // 反转前半部分链表
    p2 := head
    p1 := p2.Next
    for p2 != h1  {
        tp := p1.Next
        p1.Next = p2
        p2 = p1
        p1 = tp
        fmt.Println("loop...")
    }
    head.Next = nil

    p1 = h1
    p2 = h2
    fmt.Println("输出h1链表")
    for p1 != nil {
        fmt.Printf("%d->", p1.Val)
        p1 = p1.Next
    }
    fmt.Println("")
    fmt.Println("输出h2链表")
    for p2 != nil {
        fmt.Printf("%d->", p2.Val)
        p2 = p2.Next
    }
    fmt.Println("")

    // 从中间往两头比较
    for h1 != nil && h2 != nil {
        if h1.Val != h2.Val {
            return false
        }
        h1 = h1.Next
        h2 = h2.Next
    }
    // 原则上不存在两边长度不相同的情况
    if h1 != nil || h2 != nil {
        return false
    }
    return true
}
