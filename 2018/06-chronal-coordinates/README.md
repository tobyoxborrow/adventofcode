# Day 6: Chronal Coordinates

https://adventofcode.com/2018/day/6

## Challenges
The device produces a list of coordinates (your puzzle input). Using only the Manhattan distance, determine the area around each coordinate by counting the number of integer X,Y locations that are closest to that coordinate (and aren't tied in distance to any other coordinate).

* A: What is the size of the largest area that isn't infinite?
* B: What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?

## Sample Input
```
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
```

## Result
### rust
```
% target/release/aoc201806
A: 3293
B: 45176

Benchmark #1: target/release/aoc201806
  Time (mean ± σ):      53.1 ms ±   1.1 ms    [User: 37.0 ms, System: 1.8 ms]
  Range (min … max):    51.2 ms …  56.2 ms
```

### golang
```
% ./main
A: 3293
B: 45176

% hyperfine --warmup 3 ./main
Benchmark 1: ./main
  Time (mean ± σ):      71.0 ms ±   6.5 ms    [User: 16.4 ms, System: 15.7 ms]
  Range (min … max):    62.0 ms …  90.3 ms    36 runs
```

Part A danger regions:

![](grid_a.png)

Part B safe region:

![](grid_b.png)