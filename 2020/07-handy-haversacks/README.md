# Day 7: Handy Haversacks

https://adventofcode.com/2020/day/7

## Challenges
* A: How many bag colors can eventually contain at least one shiny gold bag?
* B: How many individual bags are required inside your single shiny gold bag?

## Sample Input
```
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
```

## Result
```Shell
time ./target/release/aoc202007
A: 177
B: 34988

real    0m0.008s
user    0m0.008s
sys     0m0.000s
```
