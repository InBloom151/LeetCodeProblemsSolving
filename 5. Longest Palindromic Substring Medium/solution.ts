function preprocess(s: string): string {
    return "#" + s.split("").join("#") + "#";
}

function palindromeSub(s: string): string {
    if (s.length < 1) {
        return "";
    };

    let processedS = preprocess(s);
    let n = processedS.length;
    let p = new Array(n).fill(0);
    let c = 0;
    let r = 0;

    for (let i = 1; i < n - 1; i++) {
        let mirror = 2 * c - i;

        if (i < r) {
            p[i] = Math.min(r - i, p[mirror]);
        };

        while (i + p[i] < n && i - (1 + p[i]) >= 0 && processedS[i + (1 + p[i])] == processedS[i - (1 + p[i])]) {
            p[i] += 1;
        };

        if (i + p[i] > r) {
            c = i;
            r = i + p[i];
        };
    };

    let maxLen = Math.max.apply(null, p);
    let centerIndex = p.indexOf(maxLen);

    let startIndex = Math.ceil((centerIndex - maxLen) / 2);
    return s.slice(startIndex, startIndex + maxLen);
};

console.log(palindromeSub("babad"));
console.log(palindromeSub("cbbd"));