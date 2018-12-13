use std::collections::HashSet;

/*
const PUZZLE: &str = "
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
";
*/
const PUZZLE: &str = include_str!("../../input");

fn main() {
    let points = parse_points(PUZZLE);
    println!("A: {}", solve_a(&points));
    //println!("B: {}", solve_b(&points));
}

// read in the points from the source input text
fn parse_points(input: &str) -> Vec<(isize,isize)> {
    let mut points = Vec::new();

    for line in input.trim().lines().filter(|x| x.find(',') != None) {
        // split x, y values from each line and parse into numbers
        let xy = line.splitn(2, ',')
            .filter_map(|s| s.trim().parse().ok())
            .collect::<Vec<isize>>();
        points.push((xy[0], xy[1]));
    }

    points
}

fn solve_a(points: &Vec<(isize,isize)>) -> usize {
    let grid_width = points.iter().map(|p| p.0).max().unwrap();
    let grid_height = points.iter().map(|p| p.1).max().unwrap();
    //println!("grid: {}w x {}h", grid_width, grid_height);

    let mut grid = Vec::new();
    let mut ignored = HashSet::new();

    // loop over each grid position
    // the grid is enlarged slightly so that we will calculate closest points
    // to grid positions outside our normal grid range. this is to identify
    // points that reach into the infinite that we can ignore
    for gy in -1..(grid_height+2) {
        for gx in -1..(grid_width+2) {
            // calculate distance of each point from this position on the grid
            // enumerate over each point, using its index as an id
            let distances = points.iter()
                .enumerate()
                .map(|p| (p.0, distance((gx, gy), ((p.1).0, (p.1).1))))
                .collect::<Vec<(usize,usize)>>();

            // find the shortest distance
            let s_distance = distances.iter().map(|p| p.1).min().unwrap();

            // the points that match the shortest distance
            let matches = distances.iter()
                .filter(|d| d.1 == s_distance)
                .map(|d| d.0)
                .collect::<Vec<usize>>();

            // save the point that is closest if there is just one
            if matches.len() == 1 {
                let closest = matches[0];
                grid.push(closest);
                //print!("{} ", closest);

                // remember which points extend into the infinite by checking
                // positions beyond the grid's expected normal size
                if gy == -1 || gy == grid_height+1 || gx == -1 || gx == grid_width+1 {
                    ignored.insert(closest);
                }
            } else {
                //print!(". ");
            }

        }
        //println!("");
    }

    /*
    for i in ignored.iter() {
        println!("{}", i);
    }
    */

    points.iter().enumerate()
        // skip point ids that reach into the infinite
        .filter(|p| ignored.contains(&p.0) == false)
        // change point ids to the count of times they appear in the grid
        .map(|p| grid.iter()
             .filter(|g| **g == p.0)
             .count()
             )
        // return the highest count
        .max()
        .unwrap()
}

// https://en.wikipedia.org/wiki/Taxicab_geometry
fn distance(a: (isize,isize), b: (isize,isize)) -> usize {
    (
        (a.0 - b.0).abs() +
        (a.1 - b.1).abs()
    ) as usize
}
