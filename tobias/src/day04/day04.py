import os


def get_aoc_input(file_path=None, remove_line_ending=True):
    """
    Gets the individual lines of an Advent of Code input file.
    :param file_path: Path to the input file. If None, try to autodetect the path from the name of this Python file
                      name.
    :param remove_line_ending: If True, remove the \n at the end of each line.
    :return: Yields every line of the file as a string.
    """
    if file_path is None:
        own_file_name = os.path.splitext(os.path.basename(__file__))[0]
        file_path = '../../data/{}/input'.format(own_file_name)
    with open(file_path) as input_file:
        for input_line in input_file.readlines():
            if remove_line_ending:
                yield input_line.replace('\n', '')
            else:
                yield input_line


def line_to_ranges(line):
    """
    Converts an input line to two range tuples.
    :param line: The line to parse.
    :return: Two tuples (a,b) and (c,d) with the ranges.
    """
    split_line = line.split(',')
    split_range1 = split_line[0].split('-')
    split_range2 = split_line[1].split('-')
    return (int(split_range1[0]), int(split_range1[1])), (int(split_range2[0]), int(split_range2[1]))


def do_ranges_contain_each_other(range1, range2):
    """
    Checks if range1 lies fully within range2 or the other way round.
    :param range1: The first range as a tuple (a,b).
    :param range2: The second range as a tuple (c,d).
    :return: True if one of the changes fully contains the other.
    """
    return (range1[0] <= range2[0] and range2[1] <= range1[1]) or (range2[0] <= range1[0] and range1[1] <= range2[1])


def do_ranges_contain_overlap(range1, range2):
    """
    Checks if range1 and range2 overlap.
    :param range1: The first range as a tuple (a,b).
    :param range2: The second range as a tuple (c,d).
    :return: True if the two ranges overlap.
    """
    range1[1] >= range2[0] and range1[0] <= range2[1]
    return (range1[1] >= range2[0] and range1[0] <= range2[1]) or (range2[1] >= range1[0] and range2[0] <= range1[1])


def day04_main(input_path=None):
    """
    Solve day 04 of Advent of Code 2022.
    :param input_path: Path to the input data. If None, attempt to autodetect the path from this Python file name.
    :return: None.
    """
    count_contain = 0
    count_overlap = 0
    for line in get_aoc_input(file_path=input_path):
        range1, range2 = line_to_ranges(line)
        if do_ranges_contain_each_other(range1, range2):
            count_contain += 1
        if do_ranges_contain_overlap(range1, range2):
            count_overlap += 1
    print('Count "Contains": {}'.format(count_contain))
    print('Count "Overlap": {}'.format(count_overlap))


if __name__ == '__main__':
    day04_main()
