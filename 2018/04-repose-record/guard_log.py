#!/usr/bin/env python3
"""
# Day 4: Repose Record

https://adventofcode.com/2018/day/4

## Challenges
Covering the walls, someone has spent an hour starting every midnight for the
past few months secretly observing this guard post! They've been writing down
the ID of the one guard on duty that night.

* A: Find the guard that has the most minutes asleep. What minute does that
     guard spend asleep the most?
* B: Of all guards, which guard is most frequently asleep on the same minute?
"""

from collections import Counter
from datetime import datetime
import re


def main():
    """Main"""
    lines = get_challenge()
    guards = parse_guard_shifts(lines)
    print("A: {}".format(solve_a(guards)))
    print("B: {}".format(solve_b(guards)))


def get_challenge():
    """Read lines from the challenge input file"""
    lines = list()
    with open("./input.txt", 'r') as hdl:
        for line in hdl.readlines():
            lines.append(line.strip())
    return lines


def parse_guard_shifts(log_lines):
    """
    Parse a line of input

    Example shift entry:
    [1518-11-02 23:56] Guard #3463 begins shift
    """

    log_entries = dict()

    # parse date from log entries
    for log_line in log_lines:
        matches = re.search(r'^\[([^]]+)\] (.*)', log_line)
        timestamp = datetime.strptime(matches.group(1), "%Y-%m-%d %H:%M")
        log_entries[timestamp] = matches.group(2)

    # parse the shifts
    shifts = list()
    shift_tmp = dict()
    nap_tmp = dict()
    for timestamp in sorted(log_entries.keys()):    # sorted by date
        if 'begins shift' in log_entries[timestamp]:
            # start a new shift, save previous naps and shifts if present
            if nap_tmp:
                shift_tmp['naps'].append(nap_tmp)
            if shift_tmp:
                shifts.append(shift_tmp)
            matches = re.search(r'Guard #(\d+)', log_entries[timestamp])
            nap_tmp = dict()
            shift_tmp = {
                'id': matches.group(1),
                'start': timestamp,
                'end': 59,
                'naps': list(),
                }
        elif 'falls asleep' in log_entries[timestamp]:
            nap_tmp = {
                'start': timestamp,
                'end': None,
                }
        elif 'wakes up' in log_entries[timestamp]:
            nap_tmp['end'] = timestamp
            shift_tmp['naps'].append(nap_tmp)
    # save previous naps and shifts if present
    if shift_tmp:
        shifts.append(shift_tmp)

    # transform shifts collection into collection of guards and their shifts
    guards = dict()
    for shift in shifts:
        # add guard entry if not present
        if shift['id'] not in guards:
            guards[shift['id']] = list()

        # create a minute element for the nap length
        # this makes it easy to count activity per minute in the solutions
        for nap in shift['naps']:
            # store which minutes they were sleeping
            min_count = nap['end'] - nap['start']
            for offset in range(int(min_count.seconds / 60)):
                guards[shift['id']].append(nap['start'].minute + offset)

    return guards


def solve_a(guards):
    # sort guards by minutes slept
    sorted_keys = sorted(guards.keys(), key=lambda guard: len(guards[guard]))
    sleepiest_guard = sorted_keys[-1]

    # find the most common minute slept
    minutes = Counter(guards[sleepiest_guard])
    sleepiest_minute = minutes.most_common(1)[0][0]

    return int(sleepiest_guard) * int(sleepiest_minute)


def solve_b(guards):
    sleepiest_log = list()
    for minute in range(0, 59):
        # go through the shifts and find who was asleep at this minute
        minute_log = (minute, None, None)
        for guard in guards:
            minutes_slept = guards[guard].count(minute)
            if not minute_log[1] or minutes_slept > minute_log[1]:
                minute_log = (minute, int(minutes_slept), int(guard))
        sleepiest_log.append(minute_log)

    solution = (None, None, None)
    for entry in sleepiest_log:
        if not solution[1] or entry[1] > solution[1]:
            solution = entry

    return solution[0] * solution[2]


if __name__ == '__main__':
    main()
