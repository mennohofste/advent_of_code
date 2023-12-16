def part_1(lines: list[str]) -> int:
    total = 0
    for line in lines:
        numbers = [c for c in line if c.isdigit()]
        total += int(numbers[0] + numbers[-1])
    return total


def part_2(lines: list[str]) -> int:
    total = 0
    for line in lines:
        left_line = ""
        left_number = None
        for char in line:
            left_line += char
            left_line = (
                left_line.replace("one", "1")
                .replace("two", "2")
                .replace("three", "3")
                .replace("four", "4")
                .replace("five", "5")
                .replace("six", "6")
                .replace("seven", "7")
                .replace("eight", "8")
                .replace("nine", "9")
            )
            if left_line[-1].isdigit():
                left_number = left_line[-1]
                break

        right_line = ""
        right_number = None
        for char in reversed(line):
            right_line = char + right_line
            right_line = (
                right_line.replace("one", "1")
                .replace("two", "2")
                .replace("three", "3")
                .replace("four", "4")
                .replace("five", "5")
                .replace("six", "6")
                .replace("seven", "7")
                .replace("eight", "8")
                .replace("nine", "9")
            )
            if right_line[0].isdigit():
                right_number = right_line[0]
                break

        if left_number is None or right_number is None:
            raise ValueError("Could not find digits in line")
        total += int(left_number + right_number)
    return total


if __name__ == "__main__":
    with open("day_1/input.txt") as f:
        lines = f.readlines()

    print("Part 1:", part_1(lines))
    print("Part 2:", part_2(lines))
