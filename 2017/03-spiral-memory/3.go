package main

import (
	"fmt"
    "io/ioutil"
    "math"
    "sort"
    "strconv"
    "strings"
)

func GetChallenge() int {
    filename := "./input"
    b, _ := ioutil.ReadFile(filename)
    i, _ := strconv.ParseInt(strings.TrimSpace(string(b)), 0, 0)
    return int(i)
}

func Solve(i int) int {
    if i == 1 {
        return 1
    }
    rn := 0             // ring number
    previous := 0       // last number in the previous ring
    for {
        squares_in_ring := rn * 8
        ring_end_square := previous + squares_in_ring
        // are we on a ring that contains the puzzle input?
        if ring_end_square > i {
            // now we know how many rings away from the access square we are
            // (c) we next need to work out the offset from the input to one of
            // the north/south/east/west pathways
            south_square := ring_end_square - rn + 1
            west_square := south_square - (squares_in_ring / 4)
            north_square := west_square - (squares_in_ring / 4)
            east_square := north_square - (squares_in_ring / 4)
            // fmt.Println(squares_in_ring)
            // fmt.Println(south_square, west_square, north_square, east_square)
            offsets := []float64 {
                math.Abs(float64(i - south_square)),
                math.Abs(float64(i - west_square)),
                math.Abs(float64(i - north_square)),
                math.Abs(float64(i - east_square)),
            }
            // sort array and take the first item as the "min" value
            // doesn't work for negative numbers, but we don't expect them here
            sorted_offsets := sort.Float64Slice(offsets)
            sort.Sort(sorted_offsets)
            offset := int(sorted_offsets[0])
            // fmt.Println(rn, offset)
            return int(rn + offset)
        }
        rn++
        previous = ring_end_square
    }
}

func main() {
    fmt.Println(Solve(1) == 1)
    fmt.Println(Solve(12) == 3)
    fmt.Println(Solve(23) == 2)
    fmt.Println(Solve(1024) == 31)
    fmt.Println(Solve(GetChallenge()))
}
