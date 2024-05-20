from collections import defaultdict

def find_substring(s: str, words: list[str]) -> list[int]:
    if not words[0]:
        return []
    
    word_length = len(words[0])
    word_count = len(words)
    word_frequency = defaultdict(int)
    
    for word in words:
        word_frequency[word] += 1
    
    result = []
    
    for i in range(word_length):
        left = i
        right = i
        current_count = 0
        current_frequency = defaultdict(int)
        
        while right + word_length <= len(s):
            word = s[right:right + word_length]
            right += word_length
            
            if word in word_frequency:
                current_frequency[word] += 1
                
                current_count += 1
                
                while current_frequency[word] > word_frequency[word]:
                    left_word = s[left:left + word_length]
                    current_frequency[left_word] -= 1
                    left += word_length
                    current_count -= 1
                
                if current_count == word_count:
                    result.append(left)
            else:
                current_frequency.clear()
                current_count = 0
                left = right
    
    return result


print(find_substring("barfoothefoobarman", ["foo","bar"]))
print(find_substring("wordgoodgoodgoodbestword", ["word","good","best","word"]))
print(find_substring("barfoofoobarthefoobarman", ["bar","foo","the"]))
print(find_substring("wordgoodgoodgoodbestword", ["word","good","best","good"]))