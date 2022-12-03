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


def get_item_priority(item):
    """
    Returns the item priority (1..52) for a given item (a..Z).
    :param item: The item to get the priority for.
    :return: The priority as an integer.
    """
    if 96 < ord(item) <= 122:
        return ord(item) - 96
    elif 64 < ord(item) <= 90:
        return ord(item) - 64 + 26
    else:
        raise ValueError('Invalid item "{}"'.format(item))


def day03_main(input_path=None):
    """
    Solve day 03 of Advent of Code 2022.
    :param input_path: Path to the input data. If None, attempt to autodetect the path from this Python file name.
    :return:
    """
    sum_of_priorities_misplaced_items = 0
    sum_of_priorities_groups = 0
    group_lines = []
    for line in get_aoc_input():
        # Day 1
        first_compartment = line[:(len(line)//2)]
        second_compartment = line[(len(line)//2):]
        for item in first_compartment:
            if second_compartment.find(item) >= 0:
                sum_of_priorities_misplaced_items += get_item_priority(item)
                break
        group_lines.append(line)
        if len(group_lines) == 3:
            for item in group_lines[0]:
                if group_lines[1].find(item) >= 0 and group_lines[2].find(item) >= 0:
                    sum_of_priorities_groups += get_item_priority(item)
                    break
            group_lines = []
    print('Sum of priorities of misplaced items: {}'.format(sum_of_priorities_misplaced_items))
    print('Sum of group priorities: {}'.format(sum_of_priorities_groups))


if __name__ == '__main__':
    day03_main()
