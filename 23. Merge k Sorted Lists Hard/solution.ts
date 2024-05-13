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

const lists: number[][] = [[1, 4, 5], [1, 3, 4], [2, 6]];

const linkedLists: (ListNode | null)[] = [];
for (const lst of lists) {
    let head: ListNode | null = new ListNode(lst[0]);
    let current: ListNode | null = head;
    for (let i = 1; i < lst.length; i++) {
        current.next = new ListNode(lst[i]);
        current = current.next;
    }
    linkedLists.push(head);
}

const result: ListNode | null = mergeKLists(linkedLists);

const output: number[] = [];
let tempResult: ListNode | null = result;
while (tempResult !== null) {
    output.push(tempResult.val);
    tempResult = tempResult.next;
}
console.log(output);