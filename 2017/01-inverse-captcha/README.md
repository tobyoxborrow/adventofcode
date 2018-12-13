# Day 1: Inverse Captcha

https://adventofcode.com/2017/day/1

## Challenges
You may only leave by solving a captcha to prove you're not a human.

The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.

* A: What is the solution to your captcha?
* B: Instead of considering the next digit, it wants you to consider the digit halfway around the circular list. What is the solution to your new captcha?

## Sample Input
```
91212129
```

## Result
### python
```
% pypy3 ./1.py
True
True
1177

% pypy3 ./1b.py
True
True
1060

hyperfine --warmup 3 'pypy3 ./1.py'
  Time (mean ± σ):      78.4 ms ±   1.8 ms    [User: 45.2 ms, System: 19.6 ms]
  Range (min … max):    74.7 ms …  82.5 ms
```

### golang
```
% ./01-inverse-captcha
1177

hyperfine --warmup 3 './01-inverse-captcha'
  Time (mean ± σ):      20.3 ms ±   1.6 ms    [User: 1.7 ms, System: 1.7 ms]
  Range (min … max):    17.9 ms …  29.2 ms
```

### rust
```
% ./rust/target/release/aoc201701
A: 1177

hyperfine --warmup 3 './rust/target/release/aoc201701'
  Time (mean ± σ):      16.8 ms ±  14.2 ms    [User: 1.3 ms, System: 1.2 ms]
  Range (min … max):    13.1 ms … 152.2 ms
```
