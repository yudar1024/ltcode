package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = merge(left.Next, right)
		return left
	}
	right.Next = merge(left, right.Next)
	return right
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找到中点
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 中点断开
	next := slow.Next
	slow.Next = nil
	// 递归排序
	left := sortList(head)
	right := sortList(next)
	// 合并有序链表
	return merge(left, right)
}

func sortLinkList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var nodes []int = nil
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	sort.Ints(nodes)
	// 构建有序链表
	dummy := &ListNode{Val: 0, Next: nil}
	cur := dummy
	for _, v := range nodes {
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return merge(left, right)

}

func main() {
	// 构建链表
	head := &ListNode{4, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{1, nil}
	head.Next.Next.Next = &ListNode{3, nil}
	// 排序
	sorted := sortLinkList(head)
	// 打印
	for sorted != nil {
		println(sorted.Val)
		sorted = sorted.Next
	}
}
