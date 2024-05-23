Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with `O(log n)` runtime complexity.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> nums = [1,3,5,6], target = 5
<b>Output:</b> 2
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> nums = [1,3,5,6], target = 2
<b>Output:</b> 1
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> nums = [1,3,5,6], target = 7
<b>Output:</b> 4
</code>
</pre>
 

#### Constraints:

- `1 <= nums.length <= 104`
- `-104 <= nums[i] <= 104`
- `nums contains distinct values sorted in ascending order.`
- `-104 <= target <= 104`
