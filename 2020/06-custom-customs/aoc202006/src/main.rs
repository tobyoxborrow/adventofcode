use std::collections::HashMap;

const PUZZLE: &str = include_str!("../../input.txt");

#[test]
fn test_parse_group() {
    assert_eq!(parse_group("abc"), (3, 3));
    assert_eq!(parse_group("a\nb\nc"), (3, 0));
    assert_eq!(parse_group("ab\nac"), (3, 1));
    assert_eq!(parse_group("a\na\na\na\na"), (1, 1));
    assert_eq!(parse_group("b"), (1, 1));
}
fn parse_group(group: &str) -> (usize, usize) {
    // part A
    let mut by_answer = HashMap::new();
    for c in group.chars().filter(|x| x != &'\n') {
        by_answer.entry(c).and_modify(|y| *y += 1).or_insert(1);
    }

    // part B
    let num_members = group.split('\n').count();
    let all_answers = by_answer.iter().filter(|x| x.1 == &num_members).count();

    (by_answer.len(), all_answers)
}

fn main() {
    let mut total_groups: usize = 0;
    let mut count_a: usize = 0;
    let mut count_b: usize = 0;

    for group in PUZZLE.split("\n\n") {
        total_groups += 1;
        // println!("{:?}", group);
        let (group_count, is_all_group) = parse_group(group.trim());
        // println!("{:?}", group_count);
        count_a += group_count;
        count_b += is_all_group;
    }

    println!("Total Groups: {}", total_groups);
    println!("Count A: {}", count_a);
    println!("Count B: {}", count_b);
}
