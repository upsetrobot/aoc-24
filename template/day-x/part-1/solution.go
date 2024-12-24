/**
 ******************************************************************************
 * Advent of Code 2024 - Day 11 Part 1
 *
 * fff
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
	"strconv"
	"strings"
)


// Types.
// None.


// Constants.
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
    solution := 0
    
    for i, line := range strLines {
        for j, pos := range line {
        }
    }

    // Print solution.
    fmt.Println("Day 11 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


