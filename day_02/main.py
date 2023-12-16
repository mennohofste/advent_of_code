#!/usr/bin/env python

RED_CUBE_MAX = 12
GREEN_CUBE_MAX = 13
BLUE_CUBE_MAX = 14


def min_valid_set(color_set: str) -> tuple[int, int, int]:
    min_red = 0
    min_green = 0
    min_blue = 0
    for color in color_set.split(","):
        amount, color = color.strip().split()
        if color == "red" and int(amount) > min_red:
            min_red = int(amount)
        if color == "green" and int(amount) > min_green:
            min_green = int(amount)
        if color == "blue" and int(amount) > min_blue:
            min_blue = int(amount)
    return (min_red, min_green, min_blue)


def min_valid_game(game: str) -> tuple[int, int, int]:
    min_red = 0
    min_green = 0
    min_blue = 0
    for color_set in game.split(";"):
        color_set = color_set.strip()
        red, green, blue = min_valid_set(color_set)
        if red > min_red:
            min_red = red
        if green > min_green:
            min_green = green
        if blue > min_blue:
            min_blue = blue
    return (min_red, min_green, min_blue)


def part_1(lines: list[str]) -> int:
    valid_games = []
    for line in lines:
        game_info, game = line.strip().split(":")
        game_id = int(game_info.split()[-1])
        red, green, blue = min_valid_game(game)
        if red <= RED_CUBE_MAX and green <= GREEN_CUBE_MAX and blue <= BLUE_CUBE_MAX:
            valid_games.append(game_id)
    return sum(valid_games)


def part_2(lines: list[str]) -> int:
    game_powers = []
    for line in lines:
        _, game = line.strip().split(":")
        red, green, blue = min_valid_game(game)
        game_powers.append(red * green * blue)
    return sum(game_powers)


if __name__ == "__main__":
    with open("day_02/input.txt") as f:
        lines = f.readlines()

    print("Part 1:", part_1(lines))
    print("Part 2:", part_2(lines))
