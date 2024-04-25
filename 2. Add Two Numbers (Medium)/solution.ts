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

const l1: ListNodeClass = new ListNodeClass(2, new ListNodeClass(4, new ListNodeClass(3)));
const l2: ListNodeClass = new ListNodeClass(5, new ListNodeClass(6, new ListNodeClass(4)));
let result: ListNodeClass | null = addTwoNumbers(l1, l2);
while (result !== null) {
    console.log(result.val);
    result = result.next;
}