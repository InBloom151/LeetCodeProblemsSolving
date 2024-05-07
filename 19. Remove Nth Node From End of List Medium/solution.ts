class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val: number = 0, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
    let first: ListNode | null = head;
    let second: ListNode | null = head;

    for (let i = 0; i < n; i++) {
        if (first === null) {
            return head;
        }
        first = first.next;
    }

    if (first === null) {
        return head !== null ? head.next : null;
    }

    while (first.next !== null) {
        first = first.next;
        second = second!.next;
    }

    if (second !== null && second.next !== null) {
        second.next = second.next.next;
    }

    return head;
}