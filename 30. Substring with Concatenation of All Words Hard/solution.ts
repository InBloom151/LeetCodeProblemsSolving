function findSubstring(s: string, words: string[]): number[] {
    if (!words[0]) {
        return [];
    }

    const wordLength = words[0].length;
    const wordCount = words.length;
    let wordFrequency: {[key: string]: number} = {};

    for (let word of words) {
        if (wordFrequency[word]) {
            wordFrequency[word]++;
        } else {
            wordFrequency[word] = 1;
        }
    }

    let result: number[] = [];

    for (let i = 0; i < wordLength; i++) {
        let left = i;
        let right = i;
        let currentCount = 0;
        let currentFrequency: {[key: string]: number} = {};

        while (right + wordLength <= s.length) {
            let word = s.slice(right, right + wordLength);
            right += wordLength;

            if (wordFrequency[word]) {
                if (currentFrequency[word]) {
                    currentFrequency[word]++;
                } else {
                    currentFrequency[word] = 1;
                }

                currentCount++;

                while (currentFrequency[word] > wordFrequency[word]) {
                    let leftWord = s.slice(left, left + wordLength);
                    currentFrequency[leftWord]--;
                    left += wordLength;
                    currentCount--;
                }

                if (currentCount === wordCount) {
                    result.push(left);
                }
            } else {
                currentFrequency = {};
                currentCount = 0;
                left = right;
            }
        }

    }
    return result;
}

console.log(findSubstring("barfoothefoobarman", ["foo","bar"]))
console.log(findSubstring("wordgoodgoodgoodbestword", ["word","good","best","word"]))
console.log(findSubstring("barfoofoobarthefoobarman", ["bar","foo","the"]))
console.log(findSubstring("wordgoodgoodgoodbestword", ["word","good","best","good"]))