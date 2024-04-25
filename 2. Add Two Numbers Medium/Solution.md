### This approach has a time complexity of `O(1)`


- Initialize a `variable` to store the `carry` from the previous addition. Initially, it's set to `0`.
- Initialize a `dummy node` as the head of the result linked list.
- Traverse both linked lists simultaneously until both lists reach the end.
- At each step:
    1. Add the values of the current nodes in both linked lists along with the `carry`.
    2. Update the `carry` for the next iteration.
    3. Create a new node with the `sum % 10` and append it to the result linked list.
    4. Move to the next nodes in both linked lists.
- After the traversal, if there's any remaining carry, create a new node with the carry and append it to the result linked list.
- Return the next node of the dummy node, which is the actual head of the result linked list.


### 1. Python

**Runtime:** `50 ms` faster than `83.06%` submissions  
**Memory usage:** `16.6 MB` less than `48.76%` submissions  

``` python
from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def addTwoNumbers(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
    result = ListNode()
    current = result
    carry = 0
    
    while l1 or l2 or carry:
        sum_val = carry
        if l1:
            sum_val += l1.val
            l1 = l1.next
        if l2:
            sum_val += l2.val
            l2 = l2.next
        
        carry, val = divmod(sum_val, 10)
        current.next = ListNode(val)
        current = current.next
    
    return result.next
```

### 2. TypeScript

**Runtime:** `101 ms` faster than `92.51%` submissions  
**Memory usage:** `57.4 MB` less than `62.35%` submissions  

``` typescript
class ListNodeClass {
    val: number;
    next: ListNodeClass | null;

    constructor(val: number = 0, next: ListNodeClass | null = null) {
        this.val = val;
        this.next = next;
    }
}

function addTwoNumbers(l1: ListNodeClass | null, l2: ListNodeClass | null): ListNodeClass | null {
    let result: ListNodeClass | null = null;
    let tail: ListNodeClass | null = null;
    let carry: number = 0;

    while (l1 !== null || l2 !== null || carry) {
        const sumVal: number = (l1 ? l1.val : 0) + (l2 ? l2.val : 0) + carry;
        carry = sumVal >= 10 ? 1 : 0;

        if (tail === null) {
            result = tail = new ListNodeClass(sumVal % 10);
        } else {
            tail.next = new ListNodeClass(sumVal % 10);
            tail = tail.next;
        }

        if (l1) l1 = l1.next;
        if (l2) l2 = l2.next;
    }

    return result;
}
```

### 3. GO

**Runtime:** `4 ms` faster than `85.34%` submissions  
**Memory usage:** `4.5 MB` less than `14.36%` submissions  

``` go
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sumVal := carry
		if l1 != nil {
			sumVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sumVal += l2.Val
			l2 = l2.Next
		}

		carry = sumVal / 10
		current.Next = &ListNode{Val: sumVal % 10}
		current = current.Next
	}

	return dummy.Next
}
```

### 4. Java

**Runtime:** `1 ms` faster than `100.00%%` submissions  
**Memory usage:** `44.6 MB` less than `19.14%` submissions  

``` java
class ListNode {
    int val;
    ListNode next;
    
    ListNode() {}
    
    ListNode(int val) {
        this.val = val;
    }
    
    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

public class solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummyHead = new ListNode(0);
        ListNode current = dummyHead;
        int carry = 0;
        
        while (l1 != null || l2 != null || carry != 0) {
            if (l1 != null) {
                carry += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                carry += l2.val;
                l2 = l2.next;
            }
            
            current.next = new ListNode(carry % 10);
            carry /= 10;
            current = current.next;
        }
        
        return dummyHead.next;
    }
}
```