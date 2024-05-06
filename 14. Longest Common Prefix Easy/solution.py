def longest_common_prefix(strs: list[str]) -> str:
    if not strs:
            return ""

    strs.sort()
    first = strs[0]
    last = strs[-1]

    for i, char in enumerate(first):
        if char != last[i]:
            return first[:i]
    return first

print(longest_common_prefix(["flower","flow","flight"]))
print(longest_common_prefix(["dog","racecar","car"]))