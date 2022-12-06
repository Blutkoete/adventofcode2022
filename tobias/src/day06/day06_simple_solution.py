def day06_main_simple(input_path, code_length):
    """
    Solve day 06 of Advent of Code 2022 in a simpler way.
    :param input_path: Path to the input data.
    :param code_length: The length of the code sequence to look for.
    :return: None.
    """
    with open(input_path) as input_file:
        code_line = input_file.readline()
        for idx in range(0, len(code_line)):
            if len(set(code_line[idx:idx+code_length])) == code_length:
                print(idx+code_length)
                break


if __name__ == '__main__':
    day06_main_simple('../../data/day06/input', 4)
    day06_main_simple('../../data/day06/input', 14)
