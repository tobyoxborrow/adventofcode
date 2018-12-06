# Day 5: Alchemical Reduction

https://adventofcode.com/2018/day/5

## Challenges
You scan the chemical composition of the suit's material and discover that it is formed by extremely long polymers (one of which is available as your puzzle input).

* A: How many units remain after fully reacting the polymer you scanned?
* B: What is the length of the shortest polymer you can produce by removing all units of exactly one type and fully reacting the result?

## Sample Input
```
hHsSmMHhhHwWlLojYCclLyJtPpTZzqdFfDYymMjJxXQOiIiSbBsGLROorMmlgvVkiIKRrGxXgZteETz
UunNbBAaWwplrRoOgGLlJvzZoNnEeqQOVjQWwqzZJBbjuUJfFSsjuKcCDdkgGvVZzmMWwUMLlodDOrU
```

## Result
### golang
```
TODO
```

### rust
```
% ./target/release/rust
A: 11814
B: 4282

% hyperfine --warmup 3 target/release/rust
Benchmark #1: target/release/rust
  Time (mean ± σ):     124.9 ms ±   5.5 ms    [User: 109.4 ms, System: 2.0 ms]
  Range (min … max):   117.8 ms … 139.2 ms
```
