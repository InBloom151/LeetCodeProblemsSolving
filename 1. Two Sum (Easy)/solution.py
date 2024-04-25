from typing import List

def twoSum(nums: List[int], target: int) -> List[int]:
    result = dict()
    
    for index, num in enumerate(nums):
        check = target - num
        if check in result:
            return [result[check], index]
        result[num] = index
        
    return []
    

examples = {
    9: [[2, 7, 11, 15], [0, 1]],
    6: [[3, 2, 4], [1, 2]],
    8: [[4, 4], [0, 1]]
}

for key, value in examples.items():
    print(twoSum(value[0], key) == value[1])
    