const PUZZLE: &str = include_str!("../../input.txt");

fn is_tree(some_char: char) -> usize {
    if some_char == '#' {
        return 1;
    }
    return 0;
}

fn slide(move_x: usize, move_y: usize) -> usize {
    let mut tree_count: usize = 0;

    let mut pos_x: usize = 0;

    for line in PUZZLE.lines().step_by(move_y) {
        let pos_char = line.chars().cycle().nth(pos_x);
        match pos_char {
            Some(v) => tree_count += is_tree(v),
            None    => break,
        }
        pos_x += move_x;
    }
    tree_count
}

fn main() {
    println!("A: {}", slide(3, 1));
    println!("B: {}", slide(1, 1) * slide(3, 1) * slide(5, 1) * slide(7, 1) * slide(1, 2));
}
