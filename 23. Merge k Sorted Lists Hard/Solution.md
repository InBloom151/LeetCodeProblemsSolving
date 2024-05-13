### This approach has a time complexity of `O(n*log(k))`

### 1. Python

**Runtime:** `79 ms` faster than `55.37%` submissions  
**Memory usage:** `19.4 MB` less than `74.69%` submissions  

``` python
from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def mergeTwoLists(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
    dummy = ListNode(0)
    current = dummy

    while l1 and l2:
        if l1.val < l2.val:
            current.next = l1
            l1 = l1.next
        else:
            current.next = l2
            l2 = l2.next
        current = current.next

    current.next = l1 if l1 else l2

    return dummy.next

def mergeKLists(lists: list[Optional[ListNode]]) -> Optional[ListNode]:
    if not lists:
        return None

    while len(lists) > 1:
        merged_lists = []
        for i in range(0, len(lists), 2):
            if i + 1 < len(lists):
                merged_lists.append(mergeTwoLists(lists[i], lists[i + 1]))
            else:
                merged_lists.append(lists[i])
        lists = merged_lists

    return lists[0]
```

### 2. TypeScript

**Runtime:** `95 ms` faster than `65.05%` submissions  
**Memory usage:** `56.5 MB` less than `85.29%` submissions  

``` typescript
class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val: number = 0, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function mergeTwoLists(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    const dummy = new ListNode(0);
    let current = dummy;

    while (l1 !== null && l2 !== null) {
        if (l1.val < l2.val) {
            current.next = l1;
            l1 = l1.next;
        } else {
            current.next = l2;
            l2 = l2.next;
        }
        current = current.next!;
    }

    current.next = l1 !== null ? l1 : l2;

    return dummy.next;
}

function mergeKLists(lists: (ListNode | null)[]): ListNode | null {
    if (lists.length === 0) return null;

    while (lists.length > 1) {
        const mergedLists: (ListNode | null)[] = [];
        for (let i = 0; i < lists.length; i += 2) {
            if (i + 1 < lists.length) {
                mergedLists.push(mergeTwoLists(lists[i], lists[i + 1]));
            } else {
                mergedLists.push(lists[i]);
            }
        }
        lists = mergedLists;
    }

    return lists[0];
}
```

### 3. GO

**Runtime:** `7 ms` faster than `79.54%` submissions  
**Memory usage:** `5.1 MB` less than `45.51%` submissions  

``` go
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedLists := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, lists[i])
			}
		}
		lists = mergedLists
	}

	return lists[0]
}
```