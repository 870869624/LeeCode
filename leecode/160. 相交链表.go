package main

// Definition for singly-linked list.
type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func getIntersectionNode(headA, headB *ListNode1) *ListNode1 {
	hashMap := make(map[*ListNode1]struct{}, 0)

	for headA != nil {
		hashMap[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// # 代码解释
// 该函数的功能是找到两个单链表 `headA` 和 `headB` 的第一个交点节点，如果没有交点则返回 `nil`。具体逻辑如下：
// 1. **空链表检查**：如果 `headA` 或 `headB` 为 `nil`，直接返回 `nil`，表示不存在交点。
// 2. **初始化指针**：定义两个指针 `pa` 和 `pb`，分别指向 `headA` 和 `headB` 的头节点。
// 3. **循环遍历**：
//    - 在每次循环中，判断 `pa` 和 `pb` 是否相等。若相等，则找到了交点，退出循环并返回该节点。
//    - 如果 `pa` 到达链表末尾（即 `pa == nil`），将 `pa` 指向 `headB` 的头节点；否则，将 `pa` 移动到下一个节点。
//    - 同理，如果 `pb` 到达链表末尾（即 `pb == nil`），将 `pb` 指向 `headA` 的头节点；否则，将 `pb` 移动到下一个节点。
// 4. **结束条件**：当 `pa` 和 `pb` 相遇时，退出循环。如果两链表无交点，则最终 `pa` 和 `pb` 都为 `nil`。

// ---

// # 控制流图
// ```mermaid
// flowchart TD
//     A[开始] --> B{任一链表为空?}
//     B -->|是| C[返回nil]
//     B -->|否| D[初始化pa和pb]
//     D --> E{pa不等于pb?}
//     E -->|是| F{pa为nil?}
//     F -->|是| G[pa指向headB]
//     F -->|否| H[pa移动到下一节点]
//     E -->|否| I[返回pa]
//     J{pb为nil?} -->|是| K[pb指向headA]
//     J -->|否| L[pb移动到下一节点]
//     G --> E
//     H --> E
//     K --> E
//     L --> E
// ```

// ### 流程图说明
// - **节点A**：流程开始。
// - **节点B**：检查是否有空链表，若有则直接返回 `nil`。
// - **节点D**：初始化指针 `pa` 和 `pb`。
// - **节点E**：判断 `pa` 和 `pb` 是否相等，若相等则返回结果。
// - **节点F/G/H**：处理 `pa` 的移动逻辑，包括切换到另一链表或继续遍历当前链表。
// - **节点J/K/L**：处理 `pb` 的移动逻辑，逻辑与 `pa` 类似。
// - **节点I**：返回交点节点或 `nil`。

func getIntersectionNode1(headA, headB *ListNode1) *ListNode1 {
	if headA == nil || headB == nil {
		return nil
	}

	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa

}
