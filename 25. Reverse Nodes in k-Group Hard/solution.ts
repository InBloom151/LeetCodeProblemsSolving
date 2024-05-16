class ListNode {
    val: number;
    next: ListNode | null;

    constructor(val: number, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function reverseKGroup(head: ListNode | null, k: number): ListNode | null {
    function reverseLinkedList(head: ListNode | null): ListNode | null {
        let prev: ListNode | null = null;
        let current: ListNode | null = head;
        while (current !== null) {
            let nextNode: ListNode | null = current.next;
            current.next = prev;
            prev = current;
            current = nextNode;
        }
        return prev;
    }

    function countNodes(node: ListNode | null): number {
        let count: number = 0;
        while (node !== null) {
            count++;
            node = node.next;
        }
        return count;
    }

    let dummy: ListNode = new ListNode(0);
    dummy.next = head;
    let prevGroupEnd: ListNode = dummy;

    let nodesCount: number = countNodes(head);

    while (nodesCount >= k) {
        let groupStart: ListNode | null = prevGroupEnd.next;
        let current: ListNode | null = groupStart;
        for (let i = 0; i < (k - 1); i++) {
            current = current!.next;
        }

        let groupEnd: ListNode | null = current;
        let nextGroupStart: ListNode | null = groupEnd!.next;

        groupEnd!.next = null;
        prevGroupEnd.next = reverseLinkedList(groupStart);
        groupStart.next = nextGroupStart;

        prevGroupEnd = groupStart;
        nodesCount -= k;
    }

    return dummy.next;
}

let head1: ListNode = new ListNode(1);
head1.next = new ListNode(2);
head1.next.next = new ListNode(3);
head1.next.next.next = new ListNode(4);
head1.next.next.next.next = new ListNode(5);

let k1: number = 2;
console.log(`Input: [1, 2, 3, 4, 5], k = ${k1}`);
let result1: ListNode | null = reverseKGroup(head1, k1);
while (result1 !== null) {
    console.log(result1.val);
    result1 = result1.next;
}


let head2: ListNode = new ListNode(1);
head2.next = new ListNode(2);
head2.next.next = new ListNode(3);
head2.next.next.next = new ListNode(4);
head2.next.next.next.next = new ListNode(5);

let k2: number = 3;
console.log(`Input: [1, 2, 3, 4, 5], k = ${k2}`);
let result2: ListNode | null = reverseKGroup(head2, k2);
while (result2 !== null) {
    console.log(result2.val);
    result2 = result2.next;
}
