def preprocess(s: str) -> str:
    return '#' + '#'.join(s) + '#'

def palindrom_sub(s: str) -> str:
    if len(s) < 1:
        return ""

    processed_s = preprocess(s)
    n = len(processed_s)
    P = [0] * n
    C, R = 0, 0

    for i in range(1, n - 1):
        mirror = 2 * C - i
        
        if i < R:
            P[i] = min(R - i, P[mirror])
        
        while i + (1 + P[i]) < n and i - (1 + P[i]) >= 0 and processed_s[i + (1 + P[i])] == processed_s[i - (1 + P[i])]:
            P[i] += 1
        
        if i + P[i] > R:
            C = i
            R = i + P[i]
    
    max_len = max(P)
    center_index = P.index(max_len)
    
    start_index = (center_index - max_len) // 2
    return s[start_index:start_index + max_len]


print(palindrom_sub("babad"))
print(palindrom_sub("cbbd"))