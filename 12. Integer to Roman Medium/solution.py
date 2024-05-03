def int_to_roman(num: int) -> str:
    ASSOCIATIONS = {
        "1": "I",
        "5": "V",
        "10": "X",
        "50": "L",
        "100": "C",
        "500": "D",
        "1000": "M",
        "4": "IV",
        "9": "IX",
        "40": "XL",
        "90": "XC",
        "400": "CD",
        "900": "CM",
    }

    result = []
    num = str(num)

    for i in range(len(num) - 1, -1, -1):

        str_num = num[i] + "0" * (len(num) - i - 1)

        if str_num.replace("0", "") == "":
            continue

        roman = ASSOCIATIONS.get(str_num)

        if roman:
            result.insert(0, roman)
        else:
            f, d = num_arr(str_num)
            result.insert(0, ASSOCIATIONS.get("1" + f[1:]) * d)
            result.insert(0, ASSOCIATIONS.get(f))

    return "".join(result)


def num_arr(str_num: str):
    arr_num = str_num.split("0")
    num = int(arr_num[0])
    x = 0
    y = ""
    if len(arr_num) > 1:
        y = arr_num[1] + "0"
    while num != 5:
        num -= 1
        x += 1
    return str(num) + y, x


print(int_to_roman(3749))
print(int_to_roman(58))
print(int_to_roman(1994))