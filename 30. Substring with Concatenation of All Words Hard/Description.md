You are given a string `s` and an array of strings `words`. All the strings of `words` are of **the same length**.

A **concatenated string** is a string that exactly contains all the strings of any permutation of `words` concatenated.

For example, if `words = ["ab","cd","ef"]`, then `"abcdef"`, `"abefcd"`, `"cdabef"`, `"cdefab"`, `"efabcd"`, and `"efcdab"` are all concatenated strings. `"acdbef"` is not a concatenated string because it is not the concatenation of any permutation of `words`.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in **any order**.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> s = "barfoothefoobarman", words = ["foo","bar"]
<b>Output:</b> [0,9]
<b>Explanation:</b> The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
<b>Output:</b> []
<b>Explanation:</b> There is no concatenated substring.
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
<b>Output:</b> [6,9,12]
<b>Explanation:</b> The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].
</code>
</pre>
 

#### Constraints:

- `1 <= s.length <= 104`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 30`
- s and `words[i]` consist of lowercase English letters.
