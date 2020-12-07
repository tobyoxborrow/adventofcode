use itertools::Itertools;

fn parse_input() -> Vec<i32> {
    const PUZZLE: &str = include_str!("../../input.txt");
    let mut input_numbers = Vec::new();
    for line in PUZZLE.lines() {
        let input_number: i32 = line.parse().expect("Not a number!");
        input_numbers.push(input_number);
    }
    input_numbers
}

fn main() {
    let input_numbers = parse_input();
    println!("A: {}", solve_a(&input_numbers));
    println!("B: {}", solve_b(&input_numbers));
}

fn solve_a(input_numbers: &Vec<i32>) -> i32 {
    let result = input_numbers.into_iter().combinations(2).filter(|x| x[0] + x[1] == 2020).nth(0).unwrap();
    println!("N: {:?}", result);
    result[0] * result[1]
}

fn solve_b(input_numbers: &Vec<i32>) -> i32 {
    let result = input_numbers.into_iter().combinations(3).filter(|x| x[0] + x[1] + x[2] == 2020).nth(0).unwrap();
    println!("N: {:?}", result);
    result[0] * result[1] * result[2]
}
