import os


REAL_POSES = {'A': 'ROCK', 'B': 'PAPER', 'C': 'SCISSORS', 'X': 'ROCK', 'Y': 'PAPER', 'Z': 'SCISSORS'}
BASE_POINTS = {'ROCK': 1, 'PAPER': 2, 'SCISSORS': 3}
BASE_POINTS_PART_TWO = {'ROCK': {'X': 3, 'Y': 1, 'Z': 2},
                        'PAPER': {'X': 1, 'Y': 2, 'Z': 3},
                        'SCISSORS': {'X': 2, 'Y': 3, 'Z': 1}}
WIN_POINTS_PART_TWO = {'X': 0, 'Y': 3, 'Z': 6}


def play_round_part_one(combination):
    split_combination = combination.split(' ')
    choice_opponent = REAL_POSES[split_combination[0]]
    my_choice = REAL_POSES[split_combination[1]]
    score = BASE_POINTS[my_choice]
    if choice_opponent == 'ROCK':
        if my_choice == 'ROCK':
            score += 3
        elif my_choice == 'PAPER':
            score += 6
        else:
            score += 0
    elif choice_opponent == 'PAPER':
        if my_choice == 'ROCK':
            score += 0
        elif my_choice == 'PAPER':
            score += 3
        else:
            score += 6
    else:
        if my_choice == 'ROCK':
            score += 6
        elif my_choice == 'PAPER':
            score += 0
        else:
            score += 3
    return score


def play_round_part_two(combination):
    split_combination = combination.split(' ')
    choice_opponent = REAL_POSES[split_combination[0]]
    win_decision = split_combination[1]
    score = BASE_POINTS_PART_TWO[choice_opponent][win_decision] + WIN_POINTS_PART_TWO[win_decision]
    return score


def day01_main(input_path):
    with open(input_path) as input_file:
        score_part1 = 0
        score_part2 = 0
        for input_line in input_file.readlines():
            score_part1 += play_round_part_one(input_line.replace('\n', ''))
            score_part2 += play_round_part_two(input_line.replace('\n', ''))
        print('Score (part 1): {}'.format(score_part1))
        print('Score (part 2): {}'.format(score_part2))


if __name__ == '__main__':
    day01_main(os.path.abspath('../../data/day02/input'))

