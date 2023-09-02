#!/usr/bin/env python3


def main():
    """Main"""
    lines = get_challenge()
    print("A: {}".format(solve_a(lines)))
    print("B: {}".format(solve_b(lines)))


def get_challenge():
    """Read lines from the challenge input file"""
    lines = list()
    with open("./input", 'r') as hdl:
        for line in hdl.readlines():
            lines.append(line.strip())
    return lines


def solve_a(box_ids):
    """Solve Part A"""
    boxes_with_two = 0
    boxes_with_three = 0
    for box_id in box_ids:
        box_id_alphabet = dict()
        box_id_has_two = False
        box_id_has_three = False
        for letter in box_id:
            if letter not in box_id_alphabet:
                box_id_alphabet[letter] = 1
            else:
                box_id_alphabet[letter] += 1
        for box_id_letter in box_id_alphabet:
            if box_id_alphabet[box_id_letter] == 2:
                box_id_has_two = True
            elif box_id_alphabet[box_id_letter] == 3:
                box_id_has_three = True
        if box_id_has_two:
            boxes_with_two += 1
        if box_id_has_three:
            boxes_with_three += 1
    return boxes_with_two * boxes_with_three


def solve_b(box_ids):
    """Solve Part B"""
    for box_id in box_ids:
        for other_box_id in box_ids:
            if box_id == other_box_id:
                continue
            # if more than one character is different we can skip it
            # we are only interested in boxes that have exactly one difference
            differences = 0
            difference_index = -1
            for index, letter in enumerate(box_id):
                other_letter = other_box_id[index]
                if letter != other_letter:
                    difference_index = index
                    differences += 1
                if differences > 1:
                    # no good
                    break
            if differences != 1:
                # no good
                continue
            # return the common letters only
            return "%s%s" % (box_id[0:difference_index], box_id[difference_index+1:])
    raise RuntimeError("Failed to solve part B")


if __name__ == '__main__':
    main()
