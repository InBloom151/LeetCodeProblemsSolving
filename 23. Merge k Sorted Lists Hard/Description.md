You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> lists = [[1,4,5],[1,3,4],[2,6]]
<b>Output:</b> [1,1,2,3,4,4,5,6]
<b>Explanation:</b> The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> lists = []
<b>Output:</b> []
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> lists = [[]]
<b>Output:</b> []
</code>
</pre>
 

#### Constraints:

- `k == lists.length`
- `0 <= k <= 104`
- `0 <= lists[i].length <= 500`
- `-104 <= lists[i][j] <= 104`
- `lists[i]` is sorted in **ascending order**.
- The sum of `lists[i].length` will not exceed `10^4`.
