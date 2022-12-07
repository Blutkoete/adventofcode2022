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


def day07_main(input_path=None):
    """
    Solve day 07 of Advent of Code 2022.
    :param input_path: Path to the input data. If None, attempt to autodetect the path from this Python file name.
    :return: None.
    """
    size_per_directory = {}
    size_per_directory.update({'/': 0})
    current_dirs = []
    for line in get_aoc_input(file_path=input_path):
        split_line = line.split(' ')
        if split_line[0] == '$':
            if split_line[1] == 'ls':
                pass
            else:
                if split_line[2] == '..':
                    current_dirs.pop()
                else:
                    current_dirs.append(split_line[2])
        elif split_line[0] == 'dir':
            full_path = ''
            for dir_ in current_dirs:
                if dir_ == '/':
                    full_path = '/'
                else:
                    full_path = '{}{}/'.format(full_path, dir_)
            full_path = '{}{}/'.format(full_path, split_line[1])
            size_per_directory.update({full_path: 0})
        else:
            full_path = ''
            for dir_ in current_dirs:
                if dir_ == '/':
                    full_path = '/'
                else:
                    full_path = '{}{}/'.format(full_path, dir_)
                current_size = size_per_directory[full_path]
                size_per_directory.update({full_path: current_size + int(split_line[0])})
    total_size_of_small_dirs = 0
    needed_space = 30000000 - (70000000 - size_per_directory['/'])
    possible_dirs = []
    for key in size_per_directory.keys():
        size_of_directory = size_per_directory[key]
        if size_of_directory <= 100000:
            total_size_of_small_dirs += size_of_directory
        if size_of_directory >= needed_space:
            possible_dirs.append(size_of_directory)
    possible_dirs.sort()
    print(total_size_of_small_dirs)
    print(possible_dirs[0])


if __name__ == '__main__':
    day07_main('../../data/day07/example')
    day07_main()
