/**
 ******************************************************************************
 * Advent of Code 2024 - Day X Part X 
 *
 * XXXXXXXX
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        XX Dec 2024
 * copyright:   2024. All rights reserved.
 * 
 ******************************************************************************
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }

    // Process file.
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

        // Parse line.
        line := scanner.Text()
        words := strings.Fields(line)

    }

    // Calculate solution.
    solution := 0


    // Print solution.
    fmt.Println("Day X Part X")
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


