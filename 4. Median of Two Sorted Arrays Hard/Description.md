Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

 

#### Example 1:

<pre>
<code>
<b>Input:</b> nums1 = [1,3], nums2 = [2]
<b>Output:</b> 2.00000
<b>Explanation:</b> merged array = [1,2,3] and median is 2.
</code>
</pre>

#### Example 2:

<pre>
<code>
<b>Input:</b> nums1 = [1,2], nums2 = [3,4]
<b>Output:</b> 2.50000
<b>Explanation:</b> merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
</code>
</pre>
 

#### Constraints:

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-106 <= nums1[i], nums2[i] <= 106`
