function longestCommonPrefix(strs: string[]): string {
    if (!strs) return "";

    strs.sort();

    let first = strs[0];
    let last = strs[strs.length - 1];

    for(let i = 0; i < first.length; i++) {
        if (first[i] != last[i]) return first.substring(0, i);
    }

    return first;
}

console.log(longestCommonPrefix(["flower","flow","flight"]))
console.log(longestCommonPrefix(["dog","racecar","car"]))