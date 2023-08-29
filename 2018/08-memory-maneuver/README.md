# Day 8: Memory Maneuver

https://adventofcode.com/2018/day/8

## Challenges
The navigation system's license file consists of a list of numbers (your puzzle input). The numbers define a data structure which, when processed, produces some kind of tree that can be used to calculate the license number.

The tree is made up of nodes; a single, outermost node forms the tree's root, and it contains all other nodes in the tree (or contains nodes that contain nodes, and so on).

* A: What is the sum of all metadata entries?
* B: What is the value of the root node?

## Sample Input
```
2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
```

Visually, these data represents the headers, children and metadata entries for nodes A, B, C and D.

```
2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
A----------------------------------
    B----------- C-----------
                     D-----
```

## Result
### golang
```
% ./main
A: 38780
B: 18232

% hyperfine --warmup 3 ./main
Benchmark 1: ./main
  Time (mean ± σ):      56.7 ms ±   4.3 ms    [User: 4.5 ms, System: 6.0 ms]
  Range (min … max):    47.4 ms …  66.9 ms    49 runs
```