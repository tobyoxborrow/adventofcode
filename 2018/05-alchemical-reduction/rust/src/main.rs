//const PUZZLE: &str = "dabAcCaCBAcCcaDA";
const PUZZLE: &str = include_str!("../../input");

fn main() {
    println!("A: {}", solve_a());
    println!("B: {}", solve_b());
}

fn solve_a() -> usize {
    let polymer: Vec<u8> = PUZZLE.trim().bytes().collect();
    trigger_polymer(polymer)
}

fn solve_b() -> usize {
    // The difference between A and a, B and b, and so on in ASCII
    const OFFSET: u8 = 32;

    let mut shortest_polymer = std::usize::MAX;

    for unit in 65..90 {    // A through Z in ASCII
        let polymer: Vec<u8> = PUZZLE.trim().bytes()
            .filter(|x| x != &(unit))
            .filter(|x| x != &(unit + OFFSET))
            .collect();
        let polymer_length = trigger_polymer(polymer);
        //println!("{} {}", unit as char, polymer_length);
        if polymer_length < shortest_polymer {
            shortest_polymer = polymer_length;
        }
    }

    shortest_polymer
}

fn trigger_polymer(mut polymer: Vec<u8>) -> usize {
    // The difference between A and a, B and b, and so on in ASCII
    const OFFSET: u8 = 32;

    let mut reactions = std::usize::MAX;

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
                if c == &(p + OFFSET) || c == &(p - OFFSET) {
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
