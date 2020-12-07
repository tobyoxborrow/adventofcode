use petgraph::{Graph, Incoming};
use petgraph::graph::NodeIndex;
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
const PUZZLE: &str = include_str!("../../input");
*/
const PUZZLE: &str = "
Step A must be finished before step B can begin.
Step B must be finished before step C can begin.
Step C must be finished before step D can begin.
Step D must be finished before step E can begin.
Step E must be finished before step F can begin.
Step F must be finished before step G can begin.
Step G must be finished before step H can begin.
";

fn main() {
    let graph = make_graph(PUZZLE);
    println!("A: {}", solve_a(&graph));
    println!("B: {}", solve_b(&graph));
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
    // start with all nodes without any incomming neighbors
    let mut available = Vec::new();
    for n in graph.node_indices() {
        if graph.neighbors_directed(n, Incoming).count() == 0 {
            available.push((graph[n], n));
        }
    }

    let mut path = Vec::new();

    while available.len() > 0 {
        // sort the list into alphabetical order so when we are picking the
        // next edge to follow it is the lowest alphabetically
        available.sort_unstable_by(|a, b| a.cmp(b));

        let mut chosen_index = None;

        // determine which next edge we can follow based on dependancies
        for (i, a) in available.iter().enumerate() {
            let n_count = graph.neighbors_directed(a.1, Incoming).count();
            let d_count = graph.neighbors_directed(a.1, Incoming)
                .filter(|x| path.contains(x))
                .count();

            if (n_count - d_count) == 0 {
                chosen_index = Some(i);
                break;
            }
        }

        if chosen_index == None {
            break;
        }

        let node = available[chosen_index.unwrap()].1;

        available.remove(chosen_index.unwrap());

        // add this node to the list we have process
        path.push(node);

        // keep track of which edges are now available paths to take
        for n in graph.neighbors(node) {
            let element = (graph[n], n);
            if ! available.contains(&element) {
                available.push(element);
            }
        }
    }

    path.iter().map(|x| graph[*x]).collect::<String>()
}

struct Worker {
    id: u8,
    job: Option<NodeIndex>,
    time_remaining: usize,
}

impl Worker {
    fn new(id: u8) -> Worker {
        Worker {
            id: id,
            job: None,
            time_remaining: 0,
        }
    }

    fn is_finished(&self) -> bool {
        self.job != None && self.time_remaining == 0
    }

    fn is_working(&self) -> bool {
        self.job != None && self.time_remaining > 0
    }

    fn tick(&mut self) {
        self.time_remaining -= 1;
    }

    fn work_on(&mut self, job: NodeIndex, cost: usize) {
        self.job = Some(job);
        self.time_remaining = cost;
    }

    // reset the worker to an initial state, returns previously worked on job
    // will panic if reset when not working on anything
    fn reset(&mut self) -> NodeIndex {
        let completed_unit = self.job.unwrap();
        self.job = None;
        self.time_remaining = 0;
        completed_unit
    }
}

fn solve_b(graph: &Graph<char,char>) -> usize {
    //const ALPHABET_SECONDS: usize = 1_560; // fixed time we must add, 60 * 26

    // our 5 workers ready to build
    let mut workers = Vec::new();
    for id in 0u8..5 {
        workers.push(Worker::new(id));
    }

    // start with all nodes without any incomming neighbors
    // there are more than one in the puzzle input so a few workers will start
    // building at the same time
    let mut available = Vec::new();
    for n in graph.node_indices() {
        if graph.neighbors_directed(n, Incoming).count() == 0 {
            available.push((graph[n], n));
        }
    }

    // though we know the path from solve_a we can use it again to help us
    // track dependancies
    let mut path = Vec::new();

    // how many seconds have pass, part of the result
    let mut ticks = 0;

    while path.len() < graph.node_count() {
        println!("tick: {} path: {} nodes: {}", ticks, path.len(), graph.node_count());
        // sort the list into alphabetical order so when we are picking the
        // next edge to follow it is the lowest alphabetically
        available.sort_unstable_by(|a, b| a.cmp(b));

        // handle each worker
        for worker in &mut workers {
            print!("worker[{}]", worker.id);
            if worker.is_working() {
                println!(" is working on {} for {} more ticks", graph[worker.job.unwrap()], worker.time_remaining);
                worker.tick();
                continue;
            } else if worker.is_finished() {
                print!(" is finsihed with {}", graph[worker.job.unwrap()]);
                path.push(worker.reset());
                // this worker can work on something new this tick
            }
            print!(" is available for work");

            // if we get here then the worker is available to work

            let mut chosen_index = None;

            // determine which next edge we can follow based on dependancies
            for (i, a) in available.iter().enumerate() {
                let n_count = graph.neighbors_directed(a.1, Incoming).count();
                let d_count = graph.neighbors_directed(a.1, Incoming)
                    .filter(|x| path.contains(x))
                    .count();

                if (n_count - d_count) == 0 {
                    chosen_index = Some(i);
                    break;
                }
            }

            if chosen_index == None {
                println!(" but no work is available right now!");
                continue;
            }

            let node = available[chosen_index.unwrap()].1;

            available.remove(chosen_index.unwrap());

            let node_cost = ((graph[node] as u8 - 65) as usize) + 60;

            println!(" and will work on {} for {} ticks", graph[node], node_cost);

            worker.work_on(node, node_cost);

            // keep track of which edges are now available paths to take
            for n in graph.neighbors(node) {
                let element = (graph[n], n);
                if ! available.contains(&element) {
                    available.push(element);
                }
            }
        }

        ticks += 1;
    }

    //ticks + ALPHABET_SECONDS
    ticks - 1
}
