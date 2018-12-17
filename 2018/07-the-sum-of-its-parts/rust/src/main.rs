use petgraph::{Graph, Incoming};
//use petgraph::algo::toposort;
use std::collections::HashMap;

/*
const PUZZLE: &str = "
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
";
*/
const PUZZLE: &str = include_str!("../../input");

fn main() {
    let graph = make_graph(PUZZLE);
    println!("A: {}", solve_a(&graph));
    //println!("B: {}", solve_b(&nodes));
}

// read in the nodes from the source input text
fn make_graph(input: &str) -> Graph<char, char> {
    let mut graph = Graph::<char, char>::new();
    let mut nodes = HashMap::new();

    // collect up all the instructions
    // fortunately the input is well formed: single letters at exact positions
    let instructions: Vec<(char, char)> = input.trim().lines()
        .filter(|x| x.contains("Step"))
        .map(|x| (x.chars().nth(5).unwrap(), x.chars().nth(36).unwrap()))
        .collect()
        ;

    // create the nodes
    for i in instructions.iter() {
        if let None = nodes.get(&i.0) {
            let index = graph.add_node(i.0);
            nodes.insert(&i.0, index);
        }
        if let None = nodes.get(&i.1) {
            let index = graph.add_node(i.1);
            nodes.insert(&i.1, index);
        }
    }

    // create the edges/lines
    for i in instructions.iter() {
        let n1 = nodes.get(&i.0).unwrap();
        let n2 = nodes.get(&i.1).unwrap();
        graph.add_edge(*n1, *n2, ' ');
    }

    graph
}

fn solve_a(graph: &Graph<char, char>) -> String {
    //let sorted = toposort(graph, None).unwrap();

    // start with all nodes without any incomming neighbors
    let mut available = Vec::new();
    for n in graph.node_indices() {
        if graph.neighbors_directed(n, Incoming).count() == 0 {
            //println!("{}", graph[n]);
            available.push((graph[n], n));
        }
    }

    let mut done = Vec::new();

    while available.len() > 0 {
        // sort the list into alphabetical order so when we are picking the
        // next edge to follow it is the lowest alphabetically
        available.sort_unstable_by(|a, b| a.cmp(b));

        /*
        print!("available[");
        for a in available.iter() {
            print!("{},", graph[a.1]);
        }
        print!("] ");
        */

        let mut chosen_index = None;

        // determine which next edge we can follow based on dependancies
        //print!("selection[");
        for (i, a) in available.iter().enumerate() {
            //print!("{},", graph[a.1]);
            let n_count = graph.neighbors_directed(a.1, Incoming).count();
            let d_count = graph.neighbors_directed(a.1, Incoming)
                .filter(|x| done.contains(x))
                .count();
            //print!("({}/{}),", n_count, d_count);

            if (n_count - d_count) == 0 {
                chosen_index = Some(i);
                //print!("!");
                break;
            }
        }
        //print!("] ");

        if chosen_index == None {
            //println!("*");
            break;
        }

        let node = available[chosen_index.unwrap()].1;

        available.remove(chosen_index.unwrap());

        // add this node to the list we have process
        done.push(node);

        //print!("chosen: {} ", graph[node]);

        // keep track of which edges are now available paths to take
        //print!("neighbors[");
        for n in graph.neighbors(node) {
            //print!("{},", graph[n]);

            let element = (graph[n], n);
            if ! available.contains(&element) {
                available.push(element);
            }
        }
        //print!("]");

        //println!();
    }

    let mut result = String::new();
    for d in done.iter() {
        result.push(graph[*d]);
    }
    result
}

/*
fn solve_b(points: &Vec<(isize,isize)>) -> usize {
    4
}
*/
