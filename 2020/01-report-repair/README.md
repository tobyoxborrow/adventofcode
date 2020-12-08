# Day 1: Report Repair

https://adventofcode.com/2020/day/1

## Challenges
* A: Find the two entries that sum to 2020; what do you get if you multiply them together?
* B: In your expense report, what is the product of the three entries that sum to 2020?

## Sample Input
```
1721
979
366
299
675
1456
```

## Result
```
% time ./target/release/aoc202001
N: [528, 1492]
A: 787776
N: [447, 611, 962]
B: 262738554

real    0m0.060s
user    0m0.031s
sys     0m0.007s


% hyperfine ./target/release/aoc202001
Benchmark #1: ./target/release/aoc202001
  Time (mean ± σ):      24.9 ms ±   0.6 ms    [User: 23.2 ms, System: 0.8 ms]
  Range (min … max):    24.2 ms …  27.4 ms    95 runs
```
