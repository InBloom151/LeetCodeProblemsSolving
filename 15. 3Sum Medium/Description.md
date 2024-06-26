Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j, i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.
 

#### Example 1:

<pre>
<code>
<b>Input:</b> nums = [-1,0,1,2,-1,-4]
<b>Output:</b> [[-1,-1,2],[-1,0,1]]
<b>Explanation:</b> 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> nums = [0,1,1]
<b>Output:</b> []
<b>Explanation:</b> The only possible triplet does not sum up to 0.
</code>
</pre>

#### Example 3:

<pre>
<code>
<b>Input:</b> nums = [0,0,0]
<b>Output:</b> [[0,0,0]]
<b>Explanation:</b> The only possible triplet sums up to 0.
</code>
</pre>
 

#### Constraints:

- `3 <= nums.length <= 3000`
- `-105 <= nums[i] <= 105`
