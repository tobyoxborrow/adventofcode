#[macro_use]
extern crate lazy_static;
use regex::Regex;
use std::collections::HashMap;

const PUZZLE: &str = include_str!("../../input.txt");

fn scan_passport(passport: &str) -> HashMap<&str, &str> {
    let mut scanned_passport = HashMap::new();
    for attribute in passport.split(|c| c == '\n' || c == ' ') {
        let tokens: Vec<&str> = attribute.split(":").collect();
        // println!("{:?} {:?}", tokens[0], tokens[1]);
        scanned_passport.insert(tokens[0], tokens[1]);
    }
    scanned_passport
}

fn is_valid_date(value: &str, min: usize, max: usize) -> bool {
    if value.len() != 4 {
        return false;
    }
    let v: usize = value.parse().unwrap();
    if v < min || v > max {
        return false;
    }
    true
}

fn is_valid_attribute(key: &str, value: &str) -> bool {
    lazy_static! {
        static ref RE_HGT: Regex = Regex::new(r"^(\d+)(cm|in)$").unwrap();
    }
    lazy_static! {
        static ref RE_HCL: Regex = Regex::new(r"^#[0-9a-z]{6}$").unwrap();
    }
    lazy_static! {
        static ref RE_PID: Regex = Regex::new(r"^\d{9}$").unwrap();
    }
    if key == "byr" {
        return is_valid_date(value, 1920, 2002);
    } else if key == "iyr" {
        return is_valid_date(value, 2010, 2020);
    } else if key == "eyr" {
        return is_valid_date(value, 2020, 2030);
    } else if key == "hgt" {
        if !RE_HGT.is_match(value) {
            return false;
        }
        let cap = RE_HGT.captures(value).unwrap();
        let num: usize = cap[1].parse().unwrap();
        let unit = &cap[2];
        if unit == "cm" {
            if num < 150 || num > 193 {
                return false;
            }
        } else if unit == "in" {
            if num < 59 || num > 76 {
                return false;
            }
        } else {
            return false;
        }
    } else if key == "hcl" {
        if !RE_HCL.is_match(value) {
            return false;
        }
    } else if key == "ecl" {
        if !["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
            .iter()
            .any(|&i| i == value)
        {
            return false;
        }
    } else if key == "pid" {
        if !RE_PID.is_match(value) {
            return false;
        }
    }
    return true;
}

fn is_valid_password(passport: &HashMap<&str, &str>) -> usize {
    let required_keys = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];
    for rkey in required_keys.iter() {
        if !passport.contains_key(rkey) {
            return 0;
        }
    }
    1
}

fn is_valid_password_b(passport: &HashMap<&str, &str>) -> usize {
    let required_keys = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];
    for rkey in required_keys.iter() {
        if !passport.contains_key(rkey) {
            // println!("{:?}", passport);
            // println!("bad passport, missing attribute {}", rkey);
            return 0;
        }
        if !is_valid_attribute(rkey, passport.get(rkey).unwrap()) {
            // println!("{:?}", passport);
            // println!("bad passport, bad attribute {}", rkey);
            return 0;
        }
    }
    1
}

fn main() {
    let mut total_passports: usize = 0;
    let mut count_a: usize = 0;
    let mut count_b: usize = 0;

    for passport in PUZZLE.split("\n\n") {
        total_passports += 1;
        let scanned_passport = scan_passport(passport.trim());
        count_a += is_valid_password(&scanned_passport);
        count_b += is_valid_password_b(&scanned_passport);
    }

    println!("Total Passports: {}", total_passports);
    println!("Valid A: {}", count_a);
    println!("Valid B: {}", count_b);
}
