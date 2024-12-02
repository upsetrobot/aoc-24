/**
 ******************************************************************************
 * Advent of Code 2024 - Day 2 Part 2 
 *
 * In this one, we have to analyze a set of numbers and determine if they 
 * meet criteria. Part 2 is made complex by adding an extra condition - a 
 * condition failure can happen once.
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        02 Dec 2024
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
    solution := 0

    for scanner.Scan() {

        // Parse line.
        line := scanner.Text()
        words := strings.Fields(line)
        var tmp []int

        for _, i := range words {
            num, err := strconv.Atoi(i)
            if err != nil {
                log.Fatal("Error converting string to number:", err)
            }

            tmp = append(tmp, num)
        }

        // Determine if report is safe.
        tmpNum := tmp[0]
        dir := 1

        if tmp[1] < tmpNum {
            dir = 0

        } else if tmp[1] == tmpNum {
            continue
        }

        flag := 1
        count := 0

        for i := range tmp {
            if i == 0 {
                continue
            }
            if !(dir == 1 && tmp[i] > tmpNum && (tmp[i] - tmpNum) >= 1 && (tmp[i] - tmpNum) <= 3) && 
                !(dir == 0 && tmp[i] < tmpNum && (tmpNum - tmp[i]) >= 1 && (tmpNum - tmp[i]) <= 3) {
                count += 1

                if count > 1 {
                    flag = 0
                    break
                }

            } else {
                tmpNum = tmp[i]
            }
        }

        solution += flag
    }

    // Print solution.
    fmt.Println("Day 2 Part 1")
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


