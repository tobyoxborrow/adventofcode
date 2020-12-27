use itertools::sorted;
use petgraph::graph::NodeIndex;
use petgraph::{Graph, Incoming};
use std::collections::HashMap;
use std::collections::HashSet;
#[macro_use]
extern crate lazy_static;
use regex::Regex;

const PUZZLE: &str = include_str!("../../input.txt");

#[test]
fn test_parse_line() {
    assert_eq!(
        parse_line("light red bags contain 1 bright white bag, 2 muted yellow bags."),
        (
            "light red".to_string(),
            vec![
                (1, "bright white".to_string()),
                (2, "muted yellow".to_string())
            ]
        )
    );
    assert_eq!(
        parse_line("bright white bags contain 1 shiny gold bag."),
        (
            "bright white".to_string(),
            vec![(1, "shiny gold".to_string()),]
        )
    );
    assert_eq!(
        parse_line("dotted black bags contain no other bags."),
        ("dotted black".to_string(), vec![])
    );
}
fn parse_line(line: &str) -> (String, Vec<(usize, String)>) {
    lazy_static! {
        static ref BAG_RE: Regex = Regex::new(r"^(.*?) bags contain (.*).$").unwrap();
        static ref SUBBAG_RE: Regex = Regex::new(r"^(\d+) (.*?) bags?").unwrap();
    }

    let cap = BAG_RE.captures(line).unwrap();
    let bag = cap[1].to_string();

    if line.contains("no other bags.") {
        return (bag, vec![]);
    }

    let mut subbags = Vec::<(usize, String)>::new();
    for subbag_str in cap[2].to_string().split(", ") {
        let cap2 = SUBBAG_RE.captures(subbag_str).unwrap();
        subbags.push((cap2[1].parse().unwrap(), cap2[2].to_string()));
    }

    (bag, subbags)
}

fn make_graph(puzzle: &str) -> Graph<String, usize> {
    let mut graph = Graph::<String, usize>::new();
    let mut nodes = HashMap::new();

    // create the nodes
    for line in puzzle.trim().lines() {
        let (bag, _) = parse_line(line);
        let index = graph.add_node(String::from(&bag));
        nodes.insert(String::from(&bag), index);
    }

    // join nodes
    for line in puzzle.trim().lines() {
        let (bag, subbags) = parse_line(line);
        let n1 = nodes.get(&bag).unwrap();
        for subbag in subbags.iter() {
            let n2 = nodes.get(&subbag.1).unwrap();
            graph.add_edge(*n1, *n2, subbag.0);
        }
    }

    graph
}

fn walk_graph(graph: &Graph<String, usize>, parent_index: NodeIndex) -> HashSet<NodeIndex> {
    let mut indexes = HashSet::new();
    let mut edges = graph.neighbors_directed(parent_index, Incoming).detach();
    while let Some(node_index) = edges.next_node(&graph) {
        indexes.insert(node_index);
        for child_index in walk_graph(&graph, node_index) {
            indexes.insert(child_index);
        }
    }

    indexes
}

#[test]
fn test_solve_a0() {
    let test_puzzle: &str = "
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
";
    assert_eq!(solve_a(test_puzzle), 4);
}
#[test]
fn test_solve_a1() {
    let test_puzzle: &str = "
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain no other bags.
vibrant plum bags contain no other bags.
";
    assert_eq!(solve_a(test_puzzle), 0);
}
#[test]
fn test_solve_a2() {
    let test_puzzle: &str = "
wicked olive bags contain 1 shiny gold bag.
shiny gold bags contain no other bags.
";
    assert_eq!(solve_a(test_puzzle), 1);
}
#[test]
fn test_solve_a3() {
    let test_puzzle: &str = "
wicked red bags contain 1 shiny gold bag.
wicked blue bags contain 2 shiny gold bags.
shiny gold bags contain no other bags.
";
    assert_eq!(solve_a(test_puzzle), 2);
}
#[test]
fn test_solve_a4() {
    let test_puzzle: &str = "
shiny gold bags contain no other bags.
wicked red bags contain 1 shiny gold bag.
awesome red bags contain 1 wicked red bag.
super red bags contain 1 awesome red bag.
";
    assert_eq!(solve_a(test_puzzle), 3);
}
#[test]
fn test_solve_a5() {
    let test_puzzle: &str = "
shiny gold bags contain no other bags.
wicked red bags contain 1 shiny gold bag.
awesome red bags contain 1 wicked red bag.
super red bags contain 1 awesome red bag.
mega red bags contain 1 super red bag.
mega yellow bags contain 1 super red bag.
uber red bags contain 1 mega red bag.
leet red bags contain 1 uber red bag.
";
    assert_eq!(solve_a(test_puzzle), 7);
}
fn solve_a(input: &str) -> usize {
    let graph = make_graph(input);
    let shiny_id = graph
        .node_indices()
        .find(|i| graph[*i] == "shiny gold")
        .unwrap();
    let all_incoming_ids = walk_graph(&graph, shiny_id);
    /*
    for incoming_id in sorted(all_incoming_ids.iter()) {
        println!(" -> {:?}", graph[*incoming_id]);
    }
    */
    all_incoming_ids.len()
}

fn main() {
    let result_a = solve_a(PUZZLE);

    println!("A: {}", result_a);
    //println!("B: {}", solve_b(&graph));
}
