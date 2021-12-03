from io import open_code


def bianry_converter(x):
    return int(x, base=2)

def halv(x):
    return x >> 1

def most_common_right_bit(numbers):
    counter = 0
    for number in numbers:
        counter += (number & 1)
    if counter > len(numbers) // 2:
        return "1"
    return "0"

def solve_first_problem(diagnostics):
    binary_length = len(diagnostics[0])
    diagnostics = list(map(bianry_converter, diagnostics))
    gamma_rate = ""

    for _ in range(binary_length):
        gamma_rate += most_common_right_bit(diagnostics)
        diagnostics = list(map(halv, diagnostics))

    gamma_rate = bianry_converter(gamma_rate[::-1])
    epsilon_rate = ((2 ** binary_length) - 1) ^ gamma_rate

    print(gamma_rate, epsilon_rate, gamma_rate *epsilon_rate)


def get_oxygen_rating(diagnostics):
    binary_length = len(diagnostics[0])
    diagnostics = list(map(bianry_converter, diagnostics))

    index = 1
    while len(diagnostics) > 1:
        zero_candidates = []
        one_candidates = []
        for number in diagnostics:
            if ((number >> binary_length-index) & 1) == 1:
                one_candidates.append(number)
            else:
                zero_candidates.append(number)
        if len(one_candidates) >= len(zero_candidates):
            diagnostics = one_candidates
        else:
            diagnostics = zero_candidates
        index += 1
    return diagnostics[0]

def get_co2_rating(diagnostics):
    binary_length = len(diagnostics[0])
    diagnostics = list(map(bianry_converter, diagnostics))

    index = 1
    while len(diagnostics) > 1:
        zero_candidates = []
        one_candidates = []
        for number in diagnostics:
            if ((number >> binary_length-index) & 1) == 1:
                one_candidates.append(number)
            else:
                zero_candidates.append(number)
        if len(zero_candidates) <= len(one_candidates):
            diagnostics = zero_candidates
        else:
            diagnostics = one_candidates
        index += 1
    return diagnostics[0]

def solve_second_problem(diagnostics):
    oxygen = get_oxygen_rating(diagnostics)
    co2 = get_co2_rating(diagnostics)
    print(oxygen, co2, oxygen * co2)


if __name__ == '__main__':
    with open("puzzle-3input") as puzzle:
        diagnostics = puzzle.read().split("\n")

    print("First soltion:")
    solve_first_problem(diagnostics)
    print("Second soltion:")
    solve_second_problem(diagnostics)