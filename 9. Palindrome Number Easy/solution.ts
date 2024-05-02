function isPalindrome(x: number): boolean {
    return x.toString() === x.toString().split("").reverse().join("");
}


console.log(isPalindrome(121))
console.log(isPalindrome(-121))
console.log(isPalindrome(10))