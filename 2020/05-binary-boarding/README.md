# Day 5: Binary Boarding

https://adventofcode.com/2020/day/5

## Challenges
* A: What is the highest seat ID on a boarding pass?
* B: What is the ID of your seat?

## Sample Input
```
BFFFBBFRRR: row 70, column 7, seat ID 567.
FFFBBBFRRR: row 14, column 7, seat ID 119.
BBFFBBFRLL: row 102, column 4, seat ID 820.
```

## Result
```Shell
% time ./target/release/aoc202005
Boarding Passes: 901
Valid A: 908
Valid B: 619

real    0m0.013s
user    0m0.002s
sys     0m0.011s

% hyperfine ./target/release/aoc202005
Benchmark #1: ./target/release/aoc202005
  Time (mean ± σ):       1.2 ms ±   0.8 ms    [User: 1.2 ms, System: 0.2 ms]
  Range (min … max):     0.6 ms …  10.4 ms    931 runs
```
