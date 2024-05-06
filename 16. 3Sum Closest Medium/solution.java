import java.util.Arrays;

class solution {
    public static void main(String[] args) {
        System.out.println(threeSumClosest(new int[]{-1,2,1,-4}, 1));
        System.out.println(threeSumClosest(new int[]{0,0,0}, 1));
    }

    public static int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        double closestSum = Double.POSITIVE_INFINITY;
        double minDiff = Double.POSITIVE_INFINITY;

        for (int i = 0; i < nums.length - 2; i++) {
            int left = i + 1;
            int right = nums.length - 1;

            while (left < right) {
                double currentSum = nums[i] + nums[left] + nums[right];
                double diff = Math.abs(currentSum - target);

                if (diff < minDiff) {
                    minDiff = diff;
                    closestSum = currentSum;
                }

                if (currentSum < target) {
                    left++;
                } else if (currentSum > target) {
                    right--;
                } else {
                    return target;
                }
            }
        }

        return (int) closestSum;
    }
}
