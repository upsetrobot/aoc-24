/**
 ******************************************************************************
 * Advent of Code 2024 - Day 1 Part 1 
 *
 * Basically, this one is sort two lists and add differences.
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        01 Dec 2024
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
    var list1 []int
    var list2 []int

    for scanner.Scan() {

        // Parse line.
        line := scanner.Text()
        words := strings.Fields(line)

        num1, err := strconv.Atoi(words[0])
        if err != nil {
            log.Fatal("Error converting string to number:", err)
        }

        num2, err := strconv.Atoi(words[1])
        if err != nil {
            log.Fatal("Error converting string to number:", err)
        }

        list1 = append(list1, num1)
        list2 = append(list2, num2)
    }

    sort.Ints(list1)
    sort.Ints(list2)

    if len(list1) != len(list2) {
        log.Fatalf("Lists do not have the same size: %d and %d", len(list1), len(list2))
    }

    // Calculate solution.
    solution := 0

    for i := range len(list1) {
        solution += absInt(list1[i] - list2[i])
    }

    // Print solution.
    fmt.Println("Day 1 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println("List 1:", list1)
    fmt.Println("List 2:", list2)
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


