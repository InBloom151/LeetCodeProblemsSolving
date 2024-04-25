Given a string `s`, find the length of the **longest substring** without repeating characters.

 

#### Example 1:

<pre>
<code>
<b>Input:</b> s = "abcabcbb"
<b>Output:</b> 3
<b>Explanation:</b> The answer is "abc", with the length of 3.
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> s = "bbbbb"
<b>Output:</b> 1
<b>Explanation:</b> The answer is "b", with the length of 1.
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> s = "pwwkew"
<b>Output:</b> 3
<b>Explanation:</b> The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
</code>
</pre>
 

Constraints:

- `0 <= s.length <= 5 * 104`
- `s` consists of English letters, digits, symbols and spaces.
