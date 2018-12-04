# Day 2: Inventory Management System

https://adventofcode.com/2018/day/2

## Challenges
You scan the likely candidate boxes again, counting the number that have an ID containing exactly two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary checksum and compare it to what your device predicts.

* A: What is the checksum for your list of box IDs?
* B: What letters are common between the two correct box IDs?

## Sample Input
```
asgwdcmbrkerohqoutfylvzpnx
asgwjcmbrkejihqoutfylvipne
asgwjcmbrkedihqoutvylizpnz
azgsjcmbrkedihqouifylvzpnx
asgwucmbrktddhqoutfylvzpnx
asgwocmbrkedihqoutfylvzivx
aqgwjcmbrkevihqvutfylvzpnx
tsgljcmbrkedihqourfylvzpnx
asgpjcmbrkedihqoutfnlvzsnx
astwjcmbrktdihqrutfylvzpnx
```

## Result
### python3
```
pypy3 ./checksum.py
A: 6175
B: asgwjcmzredihqoutcylvzinx

hyperfine --warmup 3 pypy3 ./checksum.py
Benchmark #1: pypy3 ./checksum.py
  Time (mean ± σ):      96.5 ms ±   5.8 ms    [User: 60.4 ms, System: 21.5 ms]
  Range (min … max):    91.2 ms … 115.5 ms
```

### golang
```
TODO
```

### rust
```
./rust/target/release/rust
A: 6175
B: asgwjcmzredihqoutcylvzinx

hyperfine --warmup 3 ./rust/target/release/rust
  Time (mean ± σ):      17.5 ms ±   2.2 ms    [User: 3.0 ms, System: 1.4 ms]
  Range (min … max):    15.3 ms …  30.2 ms
```
