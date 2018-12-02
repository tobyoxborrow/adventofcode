use std::fs::File;
use std::io::{BufRead, BufReader, Read};
use std::collections::HashMap;

fn main() {
    println!("A: {}", solve_a());
    println!("B: {}", solve_b());
}

fn solve_a() -> i32 {
    let file = File::open("../input").expect("file not found");
    let reader = BufReader::new(file);
    let mut frequency: i32 = 0;
    for line in reader.lines() {
        let change: i32 = line.unwrap().parse().expect("Not a number!");
        frequency += change;
    }

    frequency
}

fn solve_b() -> i32 {
    let mut frequency: i32 = 0;
    let mut history = HashMap::new();
    loop {
        let file = File::open("../input").expect("file not found");
        let mut reader = BufReader::new(file);
        for (_, line) in reader.by_ref().lines().enumerate() {
            history.insert(frequency, true);
            let change: i32 = line.unwrap().parse().expect("Not a number!");
            frequency += change;
            if let Option::Some(_i) = history.get(&frequency) {
                return frequency;
            };
        };
    }
}
