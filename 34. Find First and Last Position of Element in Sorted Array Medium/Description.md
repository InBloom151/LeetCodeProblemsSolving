Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> nums = [5,7,7,8,8,10], target = 8
<b>Output:</b> [3,4]
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> nums = [5,7,7,8,8,10], target = 6
<b>Output:</b> [-1,-1]
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> nums = [], target = 0
<b>Output:</b> [-1,-1]
</code>
</pre>
 

#### Constraints:

- `0 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`
- `nums is a non-decreasing array.`
- `-109 <= target <= 109`
