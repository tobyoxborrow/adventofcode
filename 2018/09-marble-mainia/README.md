# Day 9: Marble Mania

https://adventofcode.com/2018/day/9

## Challenges
The Elves play this game by taking turns arranging the marbles in a circle according to very particular rules. The marbles are numbered starting with 0 and increasing by 1 until every marble has a number.

* A: What is the winning Elf's score?
* B: What would the new winning Elf's score be if the number of the last marble were 100 times larger?

## Sample Input
```
10 players; last marble is worth 1618 points: high score is 8317
```

## Result
### golang
```
% ./main
A: 386018
B: 3085518618

% hyperfine --warmup 3 ./main
Benchmark 1: ./main
  Time (mean ± σ):     840.7 ms ±  75.1 ms    [User: 998.8 ms, System: 271.3 ms]
  Range (min … max):   765.7 ms … 985.6 ms    10 runs
```

The 100x increase for Part 2 is meant to be especially slow. If just Part 1 the hyperfine results are:
```
% hyperfine --warmup 3 ./mainA
Benchmark 1: ./mainA
  Time (mean ± σ):      61.3 ms ±   4.3 ms    [User: 5.4 ms, System: 9.5 ms]
  Range (min … max):    51.0 ms …  73.0 ms    44 runs
```