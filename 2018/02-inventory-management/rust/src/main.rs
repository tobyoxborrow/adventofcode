use std::collections::HashMap;

const PUZZLE: &str = include_str!("input");

fn main() {
    println!("A: {}", solve_a());
    println!("B: {}", solve_b());
}

fn solve_a() -> i32 {
    let mut count_2 = 0;
    let mut count_3 = 0;

    for box_id in PUZZLE.lines() {
        // build map of letters and how many times they occur
        let mut id_letters = HashMap::new();
        for id_letter in box_id.chars() {
            let count = id_letters.entry(id_letter).or_insert(0);
            *count += 1;
        }

        let mut has_2: bool = false;
        let mut has_3: bool = false;

        for id_letter in id_letters {
            if id_letter.1 == 2 { has_2 = true; }
            else if id_letter.1 == 3 { has_3 = true; }
        }

        if has_2 { count_2 += 1; }
        if has_3 { count_3 += 1; }
    }

    count_2 * count_3
}

fn solve_b() -> String {
    for box_id in PUZZLE.lines() {
        for other_box_id in PUZZLE.lines() {
            if box_id == other_box_id { continue; }

            // create pair of bytes from both strings
            let difference: u32 = box_id.bytes().zip(other_box_id.bytes())
                .filter(|x| x.0 != x.1) // filter just non-matching bytes
                .fold(0, |acc, _val| acc + 1);  // count how many

            if difference != 1 { continue; }

            // create pair of bytes from both strings
            let answer: Vec<u8> = box_id.bytes().zip(other_box_id.bytes())
                .filter(|x| x.0 == x.1) // filter just matching bytes
                .map(|x| x.0)   // return just one of the pairs, doesn't matter which
                .collect();     // transform into collection

            return String::from_utf8(answer).unwrap()
        }
    }
    panic!("Could not find solution");
}
