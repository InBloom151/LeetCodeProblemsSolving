function countAndSay(n: number): string {
    if (n === 1) return "1";

    function nextSequence(sequence: string): string {
        let result: string[] = [];
        let count = 1;
        let char = sequence.charAt(0);

        for (let i = 1; i < sequence.length; i++) {
            if (sequence.charAt(i) === char) {
                count++;
            } else {
                result.push(count.toString());
                result.push(char);
                char = sequence.charAt(i);
                count = 1;
            }
        }
        result.push(count.toString());
        result.push(char);

        return result.join("");
    }

    let currentSequence = "1";
    for (let i = 1; i < n; i++) {
        currentSequence = nextSequence(currentSequence);
    }
    return currentSequence;
}

console.log(countAndSay(4))
console.log(countAndSay(1))