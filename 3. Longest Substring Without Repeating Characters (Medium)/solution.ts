function longestSubstringLength(s: string): number {
    const lastOccurrence: number[] = Array(256).fill(-1);
    let start = 0;
    let maxLength = 0;

    for (let i = 0; i < s.length; i++) {
        const charCode = s.charCodeAt(i);
        if (lastOccurrence[charCode] >= start) {
            start = lastOccurrence[charCode] + 1;
        }
        lastOccurrence[charCode] = i;
        maxLength = Math.max(maxLength, i - start + 1);
    }

    return maxLength;
}

console.log(longestSubstringLength("abcabcbb"));
console.log(longestSubstringLength("bbbbb"));
console.log(longestSubstringLength("pwwkew"));