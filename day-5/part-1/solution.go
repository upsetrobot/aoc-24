/**
 ******************************************************************************
 * Advent of Code 2024 - Day 5 Part 1
 *
 * This one gives rules which you have check each pair for. I am not sure the
 * best way. I feel like a function would help, then iterate each list and
 * run the check. The function would need the rule state which is parsed
 * before the checks.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        08 Dec 2024
 * copyright:   2024. All rights reserved.
 *
 ******************************************************************************
 */

package main

import (
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

    file, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }

    // Divide the input into two lists.
    lines := strings.Split(string(file), "\n")
    var ruleList [][]int
    var pagesList [][]int

    for _, line := range lines {
        if line == "" {
            continue
        }

        // Check for comma I guess.
        if !strings.ContainsRune(line, ',') {

            // Get two numbers.
            nums := strings.Split(line, "|")

            num1, err := strconv.Atoi(nums[0])
            if err != nil {
                log.Fatal("String conversion failed:", err)
            }

            num2, err := strconv.Atoi(nums[1])
            if err != nil {
                log.Fatal("String conversion failed:", err)
            }

            ints := [] int {num1, num2}

            ruleList = append(ruleList, ints)

        } else {
            // Get two numbers.
            nums := strings.Split(line, ",")
            var ints []int

            for _, val := range nums {
                if val == "" {
                    continue
                }

                num, err := strconv.Atoi(val)
                if err != nil {
                    log.Fatal("String conversion failed:", err)
                }

                ints = append(ints, num)
            }
            pagesList = append(pagesList, ints)

        } // End if.

    } // End for.

    // Calculate solution.
    solution := 0

    // Now check each list.
    for _, pages := range pagesList {
        flag := true

        for j, page := range pages {
            
            // Look up if value is in left side. If true, find out if second value is in rest.
            for _, rule := range ruleList {
                if page == rule[0] {

                    for _, val := range pages[:j] {
                        if val == rule[1] {
                            flag = false
                            break
                        }
                    }

                    if !flag {
                        break
                    }

                } // End if.

            } // End for.

            if !flag {
                break
            }

        } // End for.

        if flag {
            solution += pages[len(pages) / 2]
        }

    } // End for.

    // Print solution.
    fmt.Println("Day 5 Part 1")
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


