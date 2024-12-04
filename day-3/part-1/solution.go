/**
 ******************************************************************************
 * Advent of Code 2024 - Day 3 Part 1 
 *
 * This appears to be a parsing problem. I wonder if regex may be a good 
 * library solution for this.
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        03 Dec 2024
 * copyright:   2024. All rights reserved.
 * 
 ******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"os"
    "regexp"
    "strconv"
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

    pattern := `mul\((\d+),(\d+)\)`

    re, err := regexp.Compile(pattern)
    if err != nil {
        log.Fatal("Regex compilation failed:", err)
    }

    matches := re.FindAllStringSubmatch(string(file), -1)
    if err != nil {
        log.Fatal("Regex search failed:", err)
    }

    // Calculate solution.
    solution := 0

    for _, i := range matches {
        arg1, err := strconv.Atoi(i[1])
        if err != nil {
            log.Fatal("String conversion failed:", err)
        }

        arg2, err := strconv.Atoi(i[2])
        if err != nil {
            log.Fatal("String conversion failed:", err)
        }

        solution += arg1 * arg2
    }

    // Print solution.
    fmt.Println("Day 3 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println("File:", file)
    fmt.Println("Matches:", matches)
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


