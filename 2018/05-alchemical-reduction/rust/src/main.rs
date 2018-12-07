//const PUZZLE: &str = "dabAcCaCBAcCcaDA";
const PUZZLE: &str = include_str!("../../input");

fn main() {
    println!("A: {}", solve_a());
    println!("B: {}", solve_b());
}

fn solve_a() -> usize {
    trigger_polymer(PUZZLE)
}

fn solve_b() -> usize {
    // improved based on solutions from Reddit:
    // use min() to select the smallest value

    const OFFSET: u8 = 32;

    let mut shortest_polymer = std::usize::MAX;

    for unit in 65u8..90 {    // A through Z in ASCII
        let polymer: &str = &PUZZLE.replace(unit as char, "")
                                   .replace((unit + OFFSET) as char, "");
        let polymer_length = trigger_polymer(polymer);
        shortest_polymer = std::cmp::min(polymer_length, shortest_polymer);
    }

    shortest_polymer
}

fn trigger_polymer(polymer: &str) -> usize {
    // improved based on solutions from Reddit
    // using stack to pop/push for efficiency
    // using 'a XOR 32' to compare units as this works for ASCII
    // e.g.: A XOR 32 == a and a XOR 32 == A

    let mut result = Vec::new();

    for c in polymer.trim().bytes() {
        if (! result.is_empty()) && &(c ^32) == result.last().unwrap() {
            result.pop();
        } else {
            result.push(c);
        }
    }

    result.len()
}
