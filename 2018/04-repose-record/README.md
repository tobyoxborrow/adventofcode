# Day 4: Repose Record

https://adventofcode.com/2018/day/4

## Challenges
Covering the walls, someone has spent an hour starting every midnight for the past few months secretly observing this guard post! They've been writing down the ID of the one guard on duty that night.

* A: Find the guard that has the most minutes asleep. What minute does that guard spend asleep the most?
* B: Of all guards, which guard is most frequently asleep on the same minute?

## Sample Input
```
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
```

## Result
### python3
```
python3 ./guard_log.py
A: 19874
B: 22687

hyperfine --warmup 3 'python3 ./guard_log.py'
Benchmark #1: python3 ./guard_log.py
  Time (mean ± σ):     377.8 ms ± 158.8 ms    [User: 103.1 ms, System: 53.7 ms]
  Range (min … max):   324.8 ms … 829.8 ms
```

### golang
```
```

### rust
```
```
