const PUZZLE: &str = include_str!("../../input");

fn main() {
    println!("A: {}", solve_a(PUZZLE));
}

fn solve_a(input: &str) -> usize {
    let mut iter = input.trim().bytes()
        .map(|x| x - 48)    // turn character code to ASCII number
        .peekable()         // so we can peek()
        ;

    let first = &input.as_bytes()[0];   // first character, for when we loop
    let mut result = 0;

    loop {
        let curr;
        match iter.next() {
            Some(v) => curr = v,
            None => break,          // actual end of input
        };

        let next;
        match iter.peek() {
            Some(v) => next = v,
            None => next = first,   // almost end of input, loop to first
        };

        if curr == *next {
            result += curr as usize
        }
    }

    result
}
