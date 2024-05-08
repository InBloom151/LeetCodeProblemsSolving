class ListNode {
    val: number;
    next: ListNode | null;

    constructor(val: number, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
    const dummy: ListNode = new ListNode();
    let current: ListNode | null = dummy;

    while (list1 !== null && list2 !== null) {
        if (list1.val < list2.val) {
            current.next = list1;
            list1 = list1.next;
        } else {
            current.next = list2;
            list2 = list2.next;
        }
        current = current.next;
    }

    if (list1 !== null) {
        current.next = list1;
    }

    if (list2 !== null) {
        current.next = list2;
    }

    return dummy.next;
}

function printList(node: ListNode | null): void {
    while (node !== null) {
        console.log(node.val);
        node = node.next;
    }
}

const list1: ListNode = new ListNode(1, new ListNode(2, new ListNode(4)));
const list2: ListNode = new ListNode(1, new ListNode(3, new ListNode(4)));

const mergedList: ListNode | null = mergeTwoLists(list1, list2);
printList(mergedList);
