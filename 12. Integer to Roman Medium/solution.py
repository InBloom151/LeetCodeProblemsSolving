def int_to_roman(num: int) -> str:
    ASSOCIATIONS = {
        1000: "M",
        900: "CM",
        500: "D",
        400: "CD",
        100: "C",
        90: "XC",
        50: "L",
        40: "XL",
        10: "X",
        9: "IX",
        5: "V",
        4: "IV",
        1: "I",
    }

    result = ""

    for k, v in ASSOCIATIONS.items():
        while num >= k:
            result += v
            num -= k

    return result


print(int_to_roman(3749))
print(int_to_roman(58))
print(int_to_roman(1994))