Given an array `nums` of `n` integers, return an array of all the **unique** quadruplets `[nums[a], nums[b], nums[c], nums[d]]` such that:

- `0 <= a, b, c, d < n`
- `a`, `b`, `c`, and `d` are **distinct**.
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

You may return the answer in **any order**.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> nums = [1,0,-1,0,-2,2], target = 0
<b>Output:</b> [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> nums = [2,2,2,2,2], target = 8
<b>Output:</b> [[2,2,2,2]]
</code>
</pre>
 

#### Constraints:

- `1 <= nums.length <= 200`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
