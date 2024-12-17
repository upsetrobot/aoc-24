/**
 ******************************************************************************
 * Advent of Code 2024 - Day 7 Part 1
 *
 * This one seems simple. You just have to parse each line and determine if
 * the numbers could be added or multiplied (or any combination thereof) to
 * result in a test value.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        15 Dec 2024
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

    file, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }

    // Process file.
    strLines := strings.Split(string(file), "\n")
    strLines = strLines[:len(strLines) - 1]

    // Calculate solution.
    solution := 0

    for _, str := range strLines {

        // Need to parse numbers.
        line := strings.ReplaceAll(str, ":", "")
        numsLine := strings.Split(line, " ")
        var nums []int

        for _, i := range numsLine {
            num, err := strconv.Atoi(i)
            if err != nil {
                log.Fatal("Error converting string to number:", err)
            }

            nums = append(nums, num)
        }

        // Now, need to iterate 2^len to try each combination.
        // Could probably use some math magic to reduce the complexity, but 
        // who would want to do that?
        // I feel like a recursive function would have been better, but I am 
        // going to try an iterative solution.
        for i := range int(math.Pow(2, float64(len(nums) - 2))) {
            testSumProduct := nums[0]
            sumProduct := nums[1]

            for j, val := range nums[2:] {
                iterVal := int(math.Pow(2, float64(j)))

                if i & iterVal == 0 {
                    sumProduct += val

                } else {
                    sumProduct *= val
                }
            }

            // Check. 
            if sumProduct == testSumProduct {
                solution += testSumProduct
                break
            }
        }
        
    } // End for.

    // Print solution.
    fmt.Println("Day 7 Part 1")
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


