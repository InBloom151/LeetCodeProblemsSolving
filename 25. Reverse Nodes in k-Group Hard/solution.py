from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def reverseKGroup(head: Optional[ListNode], k: int) -> Optional[ListNode]:
    def reverseLinkedList(head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        current = head
        while current:
            next_node = current.next
            current.next = prev
            prev = current
            current = next_node
        return prev
    
    def countNodes(node: Optional[ListNode]) -> int:
        count = 0
        while node:
            count += 1
            node = node.next
        return count
    
    dummy = ListNode(0)
    dummy.next = head
    prev_group_end = dummy
    
    nodes_count = countNodes(head)
    
    while nodes_count >= k:
        group_start = prev_group_end.next
        current = group_start
        for _ in range(k - 1):
            current = current.next
        
        group_end = current
        next_group_start = group_end.next
        
        group_end.next = None
        prev_group_end.next = reverseLinkedList(group_start)
        group_start.next = next_group_start
        
        prev_group_end = group_start
        nodes_count -= k
    
    return dummy.next

head1 = ListNode(1)
head1.next = ListNode(2)
head1.next.next = ListNode(3)
head1.next.next.next = ListNode(4)
head1.next.next.next.next = ListNode(5)

k1 = 2
print("Input: [1,2,3,4,5], k =", k1)
result1 = reverseKGroup(head1, k1)
while result1:
    print(result1.val, end=" ")
    result1 = result1.next
print()

head2 = ListNode(1)
head2.next = ListNode(2)
head2.next.next = ListNode(3)
head2.next.next.next = ListNode(4)
head2.next.next.next.next = ListNode(5)

k2 = 3
print("Input: [1,2,3,4,5], k =", k2)
result2 = reverseKGroup(head2, k2)
while result2:
    print(result2.val, end=" ")
    result2 = result2.next
print()