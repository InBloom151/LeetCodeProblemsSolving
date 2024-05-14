class ListNode {
    val: number;
    next: ListNode | null;

    constructor(val: number = 0, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function swapPairs(head: ListNode | null): ListNode | null {
    let dummy: ListNode = new ListNode(0);
    dummy.next = head;
    let prev = dummy;

    while (prev.next !== null && prev.next?.next !== null) {
        let [first, second] = [prev.next, prev.next.next];
        [prev.next, first.next, second.next] = [second, second.next, first];

        prev = first;
    }

    return dummy.next;
}

function createLinkedList(lst: number[]): ListNode | null {
    let dummy: ListNode = new ListNode(0);
    let current = dummy;
    for (let val of lst) {
        current.next = new ListNode(val);
        current = current.next;
    }
    return dummy.next;
}

function linkedListToList(head: ListNode | null): Number[] {
    let lst: number[] = [];
    let current: ListNode | null = head;
    while (current !== null) {
        lst.push(current.val);
        current = current.next;
    }
    return lst;
}

const head1 = createLinkedList([1, 2, 3, 4]);
console.log(linkedListToList(swapPairs(head1)));

const head2 = createLinkedList([]);
console.log(linkedListToList(swapPairs(head2)));

const head3 = createLinkedList([1]);
console.log(linkedListToList(swapPairs(head3)));