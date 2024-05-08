def generate_parenthesis(n: int) -> list[str]:
    result = []
    rec_generator("", 0, 0, n, result)
    return result

def rec_generator(current_string: str, opened: int, closed: int, n: int, result: list) -> None:
    if len(current_string) == 2 * n:
        result.append(current_string)
        return
    if opened < n:
        rec_generator(current_string + '(', opened + 1, closed, n, result)
    if closed < opened:
        rec_generator(current_string + ')', opened, closed + 1, n, result)


print(generate_parenthesis(3))
print(generate_parenthesis(1))