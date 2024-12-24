/**
 ******************************************************************************
 * Advent of Code 2024 - Day 12 Part 2
 *
 * This one does not seem to bad, although I imagine some hiccups will happen.
 * You have to find each region, calculate its area and its perimeter and 
 * multiply those together to find the solution.
 *
 * This is only complicated by changing perimeter to number of sides. Seems 
 * annoying, really.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        24 Dec 2024
 * copyright:   2024. All rights reserved.
 *
 ******************************************************************************
 */

package main

// Imports.
import (
	"fmt"
	"log"
	"os"
	"strings"
)


// Types.
type Value struct {
    perimeter int
    area int
}


type Coord struct {
    y int
    x int
}


// Constants.
// None.


// Variables.
// None.


/**
 * Main function that finds solution to Advent of Code problem using the
 * data from the given input file.
 */
func main() {

    // Get filename.
    if len(os.Args) < 2 {
        fmt.Println("Add filename as argument")
        return
    }

    // Get filename.
    fileName := os.Args[1]

    file, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }

    // Process file.
    // Process file.
    strLines := strings.Split(string(file), "\n")
    strLines = strLines[:len(strLines) - 1]
    var bytesLines [][] byte

    for _, line := range strLines {
        bytesLines = append(bytesLines, [] byte(line))
    }

    solution := 0
    
    // Lets go through each letter and we need to determine if we have already 
    // done it yet somehow. I'm gonna try a marker.

    for i, line := range bytesLines {
        for j, pos := range line {
            
            // Check marker.
            if pos > 'Z' {
                continue
            }

            val := getValues(bytesLines, Coord{i, j})
            solution += val.area * val.perimeter
        }
    }

    // Print solution.
    fmt.Println("Day 12 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


func getValues(mapThing [][] byte, pos Coord) Value {

    chr := mapThing[pos.y][pos.x]
    area := 1
    perimeter := 0
    mapThing[pos.y][pos.x] = chr + 32

    // Count perimeter.
    // Check left.
    if pos.x == 0 || 
        (mapThing[pos.y][pos.x - 1] != chr && mapThing[pos.y][pos.x - 1] != chr + 32) {
        perimeter += 1
    }

    // Check down.
    if pos.y == len(mapThing) - 1 || 
        (mapThing[pos.y + 1][pos.x] != chr && mapThing[pos.y + 1][pos.x] != chr + 32) {
        perimeter += 1
    }

    // Check up.
    if pos.y == 0 || 
        (mapThing[pos.y - 1][pos.x] != chr && mapThing[pos.y - 1][pos.x] != chr + 32) {
        perimeter += 1
    }

    // Check right.
    if pos.x == len(mapThing[0]) - 1 || 
        (mapThing[pos.y][pos.x + 1] != chr && mapThing[pos.y][pos.x + 1] != chr + 32) {
        perimeter += 1
    }

    ret := Value{perimeter, area}

    // Spread.
    if pos.x > 0 && mapThing[pos.y][pos.x - 1] == chr {
        valLeft := getValues(mapThing, Coord{pos.y, pos.x - 1})
        ret.area += valLeft.area
        ret.perimeter += valLeft.perimeter
    }

    if pos.y < len(mapThing) - 1 && mapThing[pos.y + 1][pos.x] == chr {
        valDown := getValues(mapThing, Coord{pos.y + 1, pos.x})
        ret.area += valDown.area
        ret.perimeter += valDown.perimeter
    }

    if pos.y > 0 && mapThing[pos.y - 1][pos.x] == chr {
        valUp := getValues(mapThing, Coord{pos.y - 1, pos.x})
        ret.area += valUp.area
        ret.perimeter += valUp.perimeter
    }

    if pos.x < len(mapThing[0]) - 1 && mapThing[pos.y][pos.x + 1] == chr {
        valRight := getValues(mapThing, Coord{pos.y, pos.x + 1})
        ret.area += valRight.area
        ret.perimeter += valRight.perimeter
    }

    return ret

} // End func. 


