Given an input string `s` and a pattern `p`, implement regular expression matching with support for `'.'` and `'*'` where:

- `'.'` Matches any single character.​​​​
- `'*'` Matches zero or more of the preceding element.

The matching should cover the **entire** input string (not partial).
 

#### Example 1:

<pre>
<code>
<b>Input:</b> s = "aa", p = "a"
<b>Output:</b> false
<b>Explanation:</b> "a" does not match the entire string "aa".
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> s = "aa", p = "a*"
<b>Output:</b> true
<b>Explanation:</b> '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> s = "ab", p = ".*"
<b>Output:</b> true
<b>Explanation:</b> ".*" means "zero or more (*) of any character (.)".
</code>
</pre>
 

#### Constraints:

- `1 <= s.length <= 20`
- `1 <= p.length <= 20`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'.'`, and `'*'`.
- It is guaranteed for each appearance of the character `'*'`, there will be a previous valid character to match.
