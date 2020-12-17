use bitvec::prelude::*;
use itertools::Itertools;
use std::collections::HashMap;

const PUZZLE: &str = include_str!("../../input.txt");

fn generate_seat_id(row: u8, col: u8) -> usize {
    (row as usize * 8) + col as usize
}

#[test]
fn test_parse_boarding_pass() {
    assert_eq!(parse_boarding_pass("BFFFBBFRRR"), (70, 7, 567));
    assert_eq!(parse_boarding_pass("FFFBBBFRRR"), (14, 7, 119));
    assert_eq!(parse_boarding_pass("BBFFBBFRLL"), (102, 4, 820));
}
fn parse_boarding_pass(str_pass: &str) -> (u8, u8, usize) {
    // sample line
    // BBFBBBBRRL
    // rrrrrrrsss  <- row/col bits of boarding pass

    // The first 7 bits are the row bits either Front(0) or Back(1)
    // Read (and write) the bit pattern in reverse, skip the 3 col bits
    let mut row_bits = BitVec::<Lsb0, u8>::new();
    for c in str_pass
        .chars()
        .rev()
        .skip(3)
        .take(7)
        .map(|x| if x == 'F' { false } else { true })
    {
        row_bits.push(c);
    }

    // The last 3 bits are the col bits either Left(0) or Right(1)
    // Read (and write) the bit pattern in reverse, only take the last 3 bits
    let mut col_bits = BitVec::<Lsb0, u8>::new();
    for c in str_pass
        .chars()
        .rev()
        .take(3)
        .map(|x| if x == 'L' { false } else { true })
    {
        col_bits.push(c);
    }

    let row: u8 = unsafe { *row_bits.as_ptr() };
    let col: u8 = unsafe { *col_bits.as_ptr() };
    let seat_id = generate_seat_id(row, col);

    (row, col, seat_id)
}

fn main() {
    let mut total_passes = 0;
    let mut highest_seat_id = 0;
    let mut seats = HashMap::<u8, Vec<u8>>::new();

    for str_pass in PUZZLE.lines() {
        total_passes += 1;
        let (row, col, seat_id) = parse_boarding_pass(str_pass);
        // println!("row: {:?}, col: {:?}, seat: {:?}", row, col, seat);
        if seat_id > highest_seat_id {
            highest_seat_id = seat_id;
        }

        // for part B
        seats
            .entry(row)
            .and_modify(|e| e.push(col))
            .or_insert(vec![col]);
    }

    // to find our seat we know from the instructions:
    // it can't be the first or last (so skip(1) to at least ignore the start)
    // it should have a row in front and behind it
    /*
    let mut our_row: u8 = 0;
    let mut our_cols: &Vec<u8> = &Vec::new();
    for window in seats.iter().sorted().skip(1).tuple_windows::<(_, _, _)>() {
        // not mentioned in the instructions, but let's assume the rows before/behind are equal size
        if window.0 .1.len() != window.2 .1.len() {
            continue;
        }
        // assume the first row we find with differing length is the one we are looking for
        if window.1 .1.len() == window.0 .1.len() {
            continue;
        }
        our_row = *window.1 .0;
        our_cols = window.1 .1;
        break;
    }
    assert_ne!(our_row, 0);
    */
    let our_window = seats
        .into_iter()
        .sorted()
        .skip(1)  // assume not first row (cockpit?)
        .tuple_windows::<(_, _, _)>()  // get 3 rows at a time
        .filter(|x| x.0 .1.len() == x.2 .1.len())  // ignore mismatched before/after rows
        .filter(|y| y.1 .1.len() != y.0 .1.len())  // find mismatched before/middle rows
        .nth(0)
        .unwrap();
    // println!("{:?}", our_window);
    let our_row = our_window.1 .0;
    let our_cols = our_window.1 .1;

    // find the empty column in our row
    /*
    let mut our_col: u8 = 0;
    for expected_col in 1..7 {
        if our_cols.contains(&expected_col) {
            continue;
        }
        our_col = expected_col;
        break;
    }
    assert_ne!(our_row, 0);
    */
    let our_col: u8 = *[1, 2, 3, 4, 5, 6, 7]
        .iter()
        .filter(|x| !our_cols.contains(x))
        .nth(0)
        .unwrap();

    let our_seat_id = generate_seat_id(our_row, our_col);

    println!("Boarding Passes: {}", total_passes);
    println!("Valid A: {}", highest_seat_id);
    println!("Valid B: {}", our_seat_id);
}
