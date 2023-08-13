# Day 3: No Matter How You Slice It

https://adventofcode.com/2018/day/3

## Challenges
Nobody can even agree on how to cut the fabric. The whole piece of fabric they're working on is a very large square - at least 1000 inches on each side. Each Elf has made a claim about which area of fabric would be ideal for Santa's suit. All claims have an ID and consist of a single rectangle with edges parallel to the edges of the fabric.

* A: How many square inches of fabric are within two or more claims?
* B: What is the ID of the only claim that doesn't overlap?

## Sample Input
```
#1 @ 861,330: 20x10
#2 @ 491,428: 28x23
#3 @ 64,746: 20x27
#4 @ 406,769: 25x28
#5 @ 853,621: 17x26
#6 @ 311,802: 27x28
#7 @ 947,977: 14x13
#8 @ 786,5: 18x23
#9 @ 420,429: 14x24
#10 @ 138,206: 29x28
```

## Result
### python3
```
% hyperfine --warmup 3 "/usr/bin/pypy3 ./sliceit.py"
Benchmark 1: /usr/bin/pypy3 ./sliceit.py
  Time (mean ± σ):      1.128 s ±  0.168 s    [User: 1.002 s, System: 0.100 s]
  Range (min … max):    0.936 s …  1.426 s    10 runs
```

### golang
```
% hyperfine --warmup 3 ./main
Benchmark 1: ./main
  Time (mean ± σ):     681.5 ms ±  20.5 ms    [User: 524.3 ms, System: 190.2 ms]
  Range (min … max):   659.1 ms … 710.9 ms    10 runs
```
