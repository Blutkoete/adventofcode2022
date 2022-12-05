import os


class CargoHold:
    """
    Represents the ship's cargo hold.
    """

    def __init__(self):
        """
        Create a new instance of a CargoHold object.
        """
        self._stacks1 = []
        self._stacks2 = []

    def add_new_crate_to_stack(self, idx, crate):
        """
        Adds a new crate to a stack. If no stack for the given index exists, a new stack is created and new empty stacks
        are added in between if necessary. New crates are added at the bottom of the stack.
        :param idx: The stack's index.
        :param crate: The crate marking.
        :return: None.
        """
        while len(self._stacks1) < idx:
            self._stacks1.append([])
            self._stacks2.append([])
        self._stacks1[idx - 1].insert(0, crate)
        self._stacks2[idx - 1].insert(0, crate)

    def move_crate(self, src_idx, dst_idx, count):
        """
        Moves crates from one stack to the other.
        :param src_idx: The source stack.
        :param dst_idx: The destination stack.
        :param count: The amount of crates to move.
        :return: None.
        """
        cargo_crate_9001 = []
        for _ in range(0, count):
            cargo_crate_9001.insert(0, self._stacks2[src_idx - 1].pop())
            self._stacks1[dst_idx - 1].append(self._stacks1[src_idx - 1].pop())
        self._stacks2[dst_idx - 1].extend(cargo_crate_9001)

    def print_stacks(self, solution_idx=1):
        """
        Prints the current stacks.
        :param solution_idx: Print the stacks for the given solution.
        :return: None.
        """
        stacks_to_print = None
        if solution_idx == 2:
            stacks_to_print = self._stacks2
        else:
            stacks_to_print = self._stacks1
        for idx_stack in range(0, len(stacks_to_print)):
            output = '{}: '.format(idx_stack + 1)
            for crate in stacks_to_print[idx_stack]:
                output = '{}{}'.format(output, crate)
            print(output)

    def __str__(self):
        """
        String representation of the CargoHold.
        :return: The solution required by Advent of Code day 5.
        """
        result_part1 = ''
        result_part2 = ''
        for idx in range(0, len(self._stacks1)):
            if len(self._stacks1[idx]) > 0:
                result_part1 = '{}{}'.format(result_part1, self._stacks1[idx][-1])
            if len(self._stacks1[idx]) > 0:
                result_part2 = '{}{}'.format(result_part2, self._stacks2[idx][-1])
        return '{}\n{}'.format(result_part1, result_part2)


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


def parse_and_apply_input(current_solution, line):
    """
    Parses an input line and applies it to the current solution.
    :param current_solution: The current solution.
    :param line: The input line.
    :return: The updated solution.
    """
    updated_solution = current_solution
    if line.startswith('move'):
        split_line = line.replace('move ', '').replace('from ', '').replace('to ', '').split(' ')
        count = int(split_line[0])
        src_idx = int(split_line[1])
        dst_idx = int(split_line[2])
        updated_solution.move_crate(src_idx, dst_idx, count)
    elif line.startswith(' 1') or line == '':
        pass
    else:
        for stack_idx, crate in zip(range(0, len(line)//4 + 1), range(1, len(line), 4)):
            if line[crate] != ' ':
                updated_solution.add_new_crate_to_stack(stack_idx + 1, line[crate])
    return updated_solution


def day05_main(input_path=None):
    """
    Solve day 04 of Advent of Code 2022.
    :param input_path: Path to the input data. If None, attempt to autodetect the path from this Python file name.
    :return: None.
    """
    solution = CargoHold()
    count = 0
    for line in get_aoc_input(input_path):
        solution = parse_and_apply_input(solution, line)
    print(solution)


if __name__ == '__main__':
    day05_main()
