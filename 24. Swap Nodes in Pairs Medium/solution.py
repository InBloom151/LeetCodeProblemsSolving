class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def swapPairs(head: ListNode) -> ListNode:
    dummy = ListNode(0)
    dummy.next = head
    prev = dummy

    while prev.next and prev.next.next:
        first, second = prev.next, prev.next.next
        prev.next, first.next, second.next = second, second.next, first

        prev = first

    return dummy.next

def createLinkedList(lst: list) -> ListNode:
    dummy = ListNode(0)
    current = dummy
    for val in lst:
        current.next = ListNode(val)
        current = current.next
    return dummy.next

def linkedListToList(head: ListNode) -> list:
    lst = []
    current = head
    while current:
        lst.append(current.val)
        current = current.next
    return lst

head1 = createLinkedList([1, 2, 3, 4])
print(linkedListToList(swapPairs(head1)))

head2 = createLinkedList([])
print(linkedListToList(swapPairs(head2)))

head3 = createLinkedList([1])
print(linkedListToList(swapPairs(head3)))
