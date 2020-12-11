# Day 3: Toboggan Trajectory

https://adventofcode.com/2020/day/3

## Challenges
* A: Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?
* B: What do you get if you multiply together the number of trees encountered on each of the listed slopes?

## Sample Input
```
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
```

## Result
```Shell
% time ./target/release/aoc202003
A: 272
B: 3898725600

real    0m0.034s
user    0m0.007s
sys     0m0.006s
```

Initial benchmark, compiling the regex on each use:
```Shell
% hyperfine ./target/release/aoc202003
Benchmark #1: ./target/release/aoc202003
  Time (mean ± σ):       3.8 ms ±   0.2 ms    [User: 2.3 ms, System: 0.7 ms]
  Range (min … max):     3.4 ms …   4.8 ms    398 runs
```
