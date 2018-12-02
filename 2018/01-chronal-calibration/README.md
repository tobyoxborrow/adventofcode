# Day 1: Chronal Calibration

https://adventofcode.com/2018/day/1

## Challenges
* A: The device shows a sequence of changes in frequency (your puzzle
input). A value like +6 means the current frequency increases by 6; a
value like -3 means the current frequency decreases by 3.
* B: You notice that the device repeats the same frequency change list
over and over. To calibrate the device, you need to find the first
frequency it reaches twice.

## Sample Input
```
-17
-20
-15
-2
-7
-4
-18
-7
-5
-6
```

## Result
### golang
```
time ./01-chronal-calibration
A: 580
B: 81972

real   0m0.045s
user   0m0.028s
sys    0m0.007s
```

### rust
```
time ./target/release/rust
A: 580
B: 81972

real	0m0.055s
user	0m0.038s
sys	0m0.010s
```
