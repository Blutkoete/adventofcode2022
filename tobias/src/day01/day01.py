import os


def day01_main(input_path):
    with open(input_path) as input_file:
        input_lines = input_file.readlines()
        maximum_carried_calories = 0
        all_calories = []
        current_carried_calories = 0
        for input_line in input_lines:
            if input_line == '\n':
                if current_carried_calories > maximum_carried_calories:
                    maximum_carried_calories = current_carried_calories
                    print('New maximum: {}'.format(maximum_carried_calories))
                all_calories.append(current_carried_calories)
                current_carried_calories = 0
            else:
                current_carried_calories += int(input_line.replace('\n', ''))
        print('Global maximum: {}'.format(maximum_carried_calories))
        all_calories.sort()
        print('Top three sum: {}'.format(sum(all_calories[-3:])))


if __name__ == '__main__':
    day01_main(os.path.abspath('../../data/day01/input'))

