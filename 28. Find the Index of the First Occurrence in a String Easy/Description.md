Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if needle is not part of `haystack`.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> haystack = "sadbutsad", needle = "sad"
<b>Output:</b> 0
<b>Explanation:</b> "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> haystack = "leetcode", needle = "leeto"
<b>Output:</b> -1
<b>Explanation:</b> "leeto" did not occur in "leetcode", so we return -1.
</code>
</pre>
 

#### Constraints:

- `1 <= haystack.length, needle.length <= 104`
- `haystack` and `needle` consist of only lowercase English characters.
