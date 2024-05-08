Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> s = "()"
<b>Output:</b> true
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> s = "()[]{}"
<b>Output:</b> true
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> s = "(]"
<b>Output:</b> false
</code>
</pre>
 

#### Constraints:

- `1 <= s.length <= 104`
- `s` consists of parentheses only `'()[]{}'`.
