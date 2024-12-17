/**
 ******************************************************************************
 * Advent of Code 2024 - Day 8 Part 1
 *
 * In this one, we have to calculate the number of antinodes formed from pairs 
 * of nodes. The antinodes are at a distance away from the nodes equal to the 
 * distance between two nodes in a line and in the direction of the line.
 * So, I guess we have to check each antenna, check each other antenna, 
 * save coordinates to a list or mark them (probably make a copy), then add 
 * them up.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        16 Dec 2024
 * copyright:   2024. All rights reserved.
 *
 ******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"math"
	"os"
    "unicode"
	"strconv"
	"strings"
)

type Thing struct {
    He 
}


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
    strLines := strings.Split(string(file), "\n")
    strLines = strLines[:len(strLines) - 1]

    // Calculate solution.
    solution := 0

    // Copy grid for marking.
    strLinesCopy := strings.Split(string(file), "\n")
    strLinesCopy = strLinesCopy[:len(strLinesCopy) - 1]
    list [] int

    for i, str := range strLines {
        for j, pos : = str {
            if unicode.isAlphanumeric(pos) {
                list = append(list, pos)


            }

        } // End for.
        
    } // End for.

    // Print solution.
    fmt.Println("Day 8 Part 1")
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


