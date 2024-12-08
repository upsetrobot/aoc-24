/**
 ******************************************************************************
 * Advent of Code 2024 - Day 5 Part 1 
 *
 * blblbl
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

            // Elimination method will not work. Maybe try starting at each 
            // 'X' and just counting strings that match.
            if chr != 'X' {
                continue
            }

            left := "X"
            down := "X"
            up := "X"
            right := "X"
            leftUp := "X"
            rightUp := "X"
            leftDown := "X"
            rightDown := "X"

            // Test each direction.
            if j >= 3 {
                left += string(line[j - 1])
                left += string(line[j - 2])
                left += string(line[j - 3])
            }

            if i < len(lines) - 4 {
                down += string(lines[i + 1][j])
                down += string(lines[i + 2][j])
                down += string(lines[i + 3][j])
                
                if j >= 3 {
                    leftDown += string(lines[i + 1][j - 1])
                    leftDown += string(lines[i + 2][j - 2])
                    leftDown += string(lines[i + 3][j - 3])
                }

                if j <= len(line) - 4 {
                    rightDown += string(lines[i + 1][j + 1])
                    rightDown += string(lines[i + 2][j + 2])
                    rightDown += string(lines[i + 3][j + 3])
                }
            }

            if i >= 3 {
                up += string(lines[i - 1][j])
                up += string(lines[i - 2][j])
                up += string(lines[i - 3][j])
                
                if j >= 3 {
                    leftUp += string(lines[i - 1][j - 1])
                    leftUp += string(lines[i - 2][j - 2])
                    leftUp += string(lines[i - 3][j - 3])
                }

                if j <= len(line) - 4 {
                    rightUp += string(lines[i - 1][j + 1])
                    rightUp += string(lines[i - 2][j + 2])
                    rightUp += string(lines[i - 3][j + 3])
                }
            }

            if j <= len(line) - 4 {
                right = line[j:j + 4]
            }

            // Count XMAS.
            if left == "XMAS" {
                solution += 1
            }

            if down == "XMAS" {
                solution += 1
            }

            if up == "XMAS" {
                solution += 1
            }

            if right == "XMAS" {
                solution += 1
            }

            if leftDown == "XMAS" {
                solution += 1
            }

            if leftUp == "XMAS" {
                solution += 1
            }

            if rightDown == "XMAS" {
                solution += 1
            }

            if rightUp == "XMAS" {
                solution += 1
            }

        } // End for.

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


