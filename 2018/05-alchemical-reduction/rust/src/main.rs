//const PUZZLE: &str = "dabAcCaCBAcCcaDA";
const PUZZLE: &str = include_str!("../../input");

fn main() {
    println!("A: {}", solve_a());
    //println!("B: {}", solve_b());
}

fn solve_a() -> usize {
    // The difference between a and A, b and B, and so on in ASCII
    let offset: u8 = 32;

    let mut polymer: Vec<u8> = PUZZLE.trim().bytes().collect();
    let mut reactions = 1;
    // loop until there are no more reactions within the polymer
    while reactions != 0 {
        let mut new_polymer = Vec::new();

        // braces to create scope for borrowing polymer
        // TODO: be better at rust ownerships
        {
            let mut iter = polymer.iter().peekable();
            let mut reacting = false;

            loop {
                // if we encountered reactive units on a previous iteration then here we will skip
                // so we consume both units
                if reacting {
                    let result = iter.next();
                    match result {
                        Some(_) => (),
                        None => break,
                    };
                    reacting = false;
                }

                // get the current unit
                let c: &u8;
                let result = iter.next();
                match result {
                    Some(v) => c = v,
                    None => break,
                }

                // peek at the next unit
                // if the peek fails, we are at the end
                let result = iter.peek();
                let p: &u8;
                match result {
                    Some(v) => p = *v,
                    None => {
                        new_polymer.push(*c);
                        break
                    },
                };

                // if the unit reacts with the next, skip this unit and the next
                if c == &(p + offset) || c == &(p - offset) {
                    reacting = true;
                    continue;
                }

                // if the unit doesn't react with the next, save it
                new_polymer.push(*c);
            }
        }

        reactions = polymer.len() - new_polymer.len();
        polymer = new_polymer;

        /*
        println!("Progress: {} {}", polymer.len(), changes);
        for c in polymer.iter().cloned() {
            print!("{}", c as char);
        }
        println!("");
        */
    }

    /*
    println!("Final:");
    for c in polymer.iter().cloned() {
        print!("{}", c as char);
    }
    println!("");
    */

    // count of how many units are in the polymer after reactions
    polymer.iter().cloned().count()
}
