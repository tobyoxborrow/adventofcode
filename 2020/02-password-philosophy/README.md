# Day 2: Password Philosophy

https://adventofcode.com/2020/day/2

## Challenges
* A: How many passwords are valid according to their policies?
* B: How many passwords are valid according to the new interpretation of the policies?

## Sample Input
```
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
```

## Result
```Shell
% ./target/release/aoc202002
A: 628
B: 705
```

Initial benchmark, compiling the regex on each use:
```Shell
% hyperfine ./target/release/aoc202002
Benchmark #1: ./target/release/aoc202002
  Time (mean ± σ):      80.9 ms ±   1.7 ms    [User: 77.5 ms, System: 1.6 ms]
  Range (min … max):    79.0 ms …  85.9 ms    34 runs
```

Speed-up after using lazy_static to compile the regex once:
```Shell
% hyperfine ./target/release/aoc202002
Benchmark #1: ./target/release/aoc202002
  Time (mean ± σ):       3.6 ms ±   0.4 ms    [User: 1.9 ms, System: 0.9 ms]
  Range (min … max):     2.5 ms …   5.6 ms    446 runs
```
