# Day 4: Passport Processing

https://adventofcode.com/2020/day/4

## Challenges
* A: In your batch file, how many passports are valid?
* B: Count the number of valid passports - those that have all required fields and valid values. Continue to treat cid as optional. In your batch file, how many passports are valid?

## Sample Input
```
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
```

## Result
```Shell
% time ./target/release/aoc202004
Total Passports: 259
Valid A: 196
Valid B: 114

real    0m0.047s
user    0m0.006s
sys     0m0.006s

% hyperfine ./target/release/aoc202004
Benchmark #1: ./target/release/aoc202004
  Time (mean ± σ):       4.5 ms ±   0.3 ms    [User: 2.5 ms, System: 1.1 ms]
  Range (min … max):     3.5 ms …   5.7 ms    346 runs
```
