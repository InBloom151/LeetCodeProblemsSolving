import java.util.HashMap;
import java.util.Map;

public class solution {

    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> result = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int check = target - nums[i];
            if (result.containsKey(check)) {
                return new int[]{result.get(check), i};
            }
            result.put(nums[i], i);
        }

        return new int[]{};
    }

    public static void main(String[] args) {
        Map<Integer, int[][]> examples = Map.of(
                9, new int[][]{{2, 7, 11, 15}, {0, 1}},
                6, new int[][]{{3, 2, 4}, {1, 2}},
                8, new int[][]{{4, 4}, {0, 1}}
        );

        for (Map.Entry<Integer, int[][]> entry : examples.entrySet()) {
            System.out.println(equalArrays(twoSum(entry.getValue()[0], entry.getKey()), entry.getValue()[1]));
        }
    }

    public static boolean equalArrays(int[] arr1, int[] arr2) {
        if (arr1.length != arr2.length) {
            return false;
        }
        for (int i = 0; i < arr1.length; i++) {
            if (arr1[i] != arr2[i]) {
                return false;
            }
        }
        return true;
    }
} 
