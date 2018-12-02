use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("A: {}", solve_a());
    //println!("B: {}", solve_b());
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
