def longest_substring_length(s: str) -> int:
    last_occurrence = [-1] * 256
    start = 0
    max_length = 0

    for i, char in enumerate(s):
        char_code = ord(char)
        if last_occurrence[char_code] >= start:
            start = last_occurrence[char_code] + 1
        last_occurrence[char_code] = i
        max_length = max(max_length, i - start + 1)

    return max_length
        

print(longest_substring_length("abcabcbb"))
print(longest_substring_length("bbbbb"))
print(longest_substring_length("pwwkew"))
