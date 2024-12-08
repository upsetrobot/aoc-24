/**
 ******************************************************************************
 * Advent of Code 2024 - Day 6 Part 2 
 *
 * This should be pretty simple. You have to count the number of spaces of a 
 * path of a guard that always turns right when encountering a '#'.
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
	"strings"
)


type Dir int


const (
    Left Dir = iota
    Down
    Up
    Right
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

    // Well, we could just trace the path of the guard.
    strLines := strings.Split(string(file), "\n")
    strLines = strLines[:len(strLines) - 1]

    // Convert to bytes.
    lines := make([][]byte, len(strLines))

    for i, str := range strLines {
        lines[i] = []byte(str)
    }

    // First, find the position of the guard.
    coords := []int {0, 0}

    for i, line := range lines {
        for j, chr := range line {
            if chr == '^' {
                coords = []int {i, j}
                break
            }
        }

        if coords[0] != 0 && coords[1] != 0 {
            break
        }

    } // End for.

    // Now we have to count the path.
    // Calculate solution.
    solution := 0
    var dir Dir = Up

    // The guard cannot be on the outside ring, right?
    for coords[0] != 0 && 
        coords[0] != len(lines) - 1 &&
        coords[1] != 0 &&
        coords[1] != len(lines[0]) - 1 {

        // Need to move.
        // Get direction.
        switch dir {
            case Left:
                if lines[coords[0]][coords[1] - 1] == '#' {
                    dir = Up
                }
                break
            case Down:
                if lines[coords[0] + 1][coords[1]] == '#' {
                    dir = Left
                }
                break
            case Up:
                if lines[coords[0] - 1][coords[1]] == '#' {
                    dir = Right
                }
                break
            case Right:
                if lines[coords[0]][coords[1] + 1] == '#' {
                    dir = Down
                }
                break
            default:
                log.Fatal("Direction not defined correctly")

        } // End switch.

        // Count and mark.
        if lines[coords[0]][coords[1]] != 'X' {
            solution += 1
        }

        lines[coords[0]][coords[1]] = 'X'

        // Move direction.
        switch dir {
            case Left:
                coords[1] -= 1
                break
            case Down:
                coords[0] += 1
                break
            case Up:
                coords[0] -= 1
                break
            case Right:
                coords[1] += 1
                break
            default:
                log.Fatal("Direction not defined correctly")
        }


    } // End for.

    solution += 1       // For last step.

    // Print solution.
    fmt.Println("Day 6 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


