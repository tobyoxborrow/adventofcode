#[macro_use] extern crate lazy_static;

use regex::Regex;

fn parse_line(line: &str) -> (usize, usize, char, String) {
    // Example line:
    // 1-3 a: abcde
    lazy_static! {
        static ref RE: Regex = Regex::new(r"^(\d+)-(\d+) ([a-z]): (.*)$").unwrap();
    }
    let cap = RE.captures(line).unwrap();
    let n1: usize = cap[1].parse().unwrap();
    let n2: usize = cap[2].parse().unwrap();
    let c: char = cap[3].chars().nth(0).unwrap();
    (n1, n2, c, cap[4].to_string())
}

fn main() {
    const PUZZLE: &str = include_str!("../../input.txt");

    let mut count_a = 0;
    let mut count_b = 0;

    for line in PUZZLE.lines() {
        let (n1, n2, c, password) = parse_line(line);
        // println!("{}, {}, {}, {}", n1, n2, c, password);

        // in part A the first two numbers are min/max
        let c_count = password.chars().filter(|x| x == &c).count();
        if c_count >= n1 && c_count <= n2 {
            count_a += 1;
        }

        // in part B the first two numbers are character positions
        let p1 = password.chars().nth(n1 - 1).unwrap();
        let p2 = password.chars().nth(n2 - 1).unwrap();
        if (p1 == c && p2 != c) || (p1 != c && p2 == c)  {
            count_b += 1;
        }
    }

    println!("A: {}", count_a);
    println!("B: {}", count_b);
}
