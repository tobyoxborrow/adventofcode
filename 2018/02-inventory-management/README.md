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
time pypy3 checksum.py
A: 6175
B: asgwjcmzredihqoutcylvzinx

real	0m0.086s
user	0m0.058s
sys	0m0.023s
```

### golang
```
TODO
```

### rust
```
time ./target/release/rust
A: 6175

real	0m0.013s
user	0m0.002s
sys	0m0.003s
```
