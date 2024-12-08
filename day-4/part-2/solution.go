/**
 ******************************************************************************
 * Advent of Code 2024 - Day 4 Part 2 
 *
 * This is a word search, but this time, two words are required to be crossing.
 * This makes the challenge quite a bit more difficult. I think i will consider 
 * approaching this the same way as before but perhaps only looking right, so 
 * that I can avoid double-counting.
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        07 Dec 2024
 * copyright:   2024. All rights reserved.
 * 
 ******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)


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

    // Calculate solution.
    solution := 0
    lines := strings.Split(string(file), "\n")

    for i, line := range lines {

        for j, chr := range line {

            // I want to check diagonal down right and down diagonal up right for all 'M's and 'S's.
            if chr != 'M' && chr != 'S' {
                continue
            }

            if j > len(line) - 3 || i >= len(lines) - 3 {
                continue
            }

            // Make strings.
            str1 := string(chr) + string(lines[i + 1][j + 1]) + string(lines[i + 2][j + 2])
            str2 := string(lines[i + 2][j]) + string(lines[i + 1][j + 1]) + string(lines[i][j + 2])
            
            if (str1 == "MAS" || str1 == "SAM") && (str2 == "MAS" || str2 == "SAM") {
                solution += 1
            }

        } // End for.

    } // End for.

    // Print solution.
    fmt.Println("Day 4 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


/**
 * Function returns the absolute value of a given argument.
 * 
 * Arguments:
 *     x:int   Value to return absolute value of.
 * 
 * Returns:
 *     int     Absolute value of x.
 */
func absInt(x int) int {
    if x < 0 {
        return -x
    }

    return x
}


