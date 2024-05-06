def letter_combinations(digits: str) -> list[str] | list:
    PHONE_NUMS = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }
    if digits == "":
        return []
    
    result = list()
    get_combinations('', 0, digits, result, PHONE_NUMS)
    return result


def get_combinations(combination: str, index: int, digits: str, result: list, associations: dict) -> None:
    if index == len(digits):
        result.append(combination)
        return
    current_digit = digits[index]
    for letter in associations[current_digit]:
        get_combinations(combination + letter, index + 1, digits, result, associations)
    

print(letter_combinations("23"))
print(letter_combinations(""))
print(letter_combinations("2"))

    