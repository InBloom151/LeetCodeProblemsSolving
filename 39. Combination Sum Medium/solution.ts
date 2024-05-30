function combinationSum(candidates: number[], target: number): number[][] {
    let result: number[][] = [];

    function backtrack(remain: number, combo: number[], start: number): void {
        if (remain === 0) {
            result.push(Array.from(combo));
            return;
        } else if (remain < 0) {
            return;
        }

        for (let i = start; i < candidates.length; i++) {
            combo.push(candidates[i]);
            backtrack(remain - candidates[i], combo, i);
            combo.pop();
        }
    }

    backtrack(target, [], 0);
    return result;
}

let candidates1 = [2, 3, 6, 7];
let target1 = 7;
console.log(combinationSum(candidates1, target1));
let candidates2 = [2, 3, 5];
let target2 = 8;
console.log(combinationSum(candidates2, target2));
let candidates3 = [2];
let target3 = 1;
console.log(combinationSum(candidates3, target3));