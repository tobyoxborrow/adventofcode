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
A: 38780
```