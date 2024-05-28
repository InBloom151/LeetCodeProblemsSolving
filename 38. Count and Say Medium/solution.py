def count_and_say(n: int) -> str:
    if n == 1:
        return "1"
    
    def next_sequence(sequence: str) -> str:
        result = []
        count = 1
        char = sequence[0]

        for i in range(1, len(sequence)):
            if sequence[i] == char:
                count += 1
            else:
                result.append(str(count))
                result.append(char)
                char = sequence[i]
                count = 1
        result.append(str(count))
        result.append(char)

        return "".join(result)
    
    current_sequence = "1"
    for _ in range(1, n):
        current_sequence = next_sequence(current_sequence)

    return current_sequence

print(count_and_say(4))
print(count_and_say(1))