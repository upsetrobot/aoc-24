/**
 ******************************************************************************
 * Advent of Code 2024 - Day 6 Part 2
 *
 * This should be pretty simple. You have to count the number of spaces of a
 * path of a guard that always turns right when encountering a '#'.
 *
 * For part 2, it much more difficult by asking for us to add an obstacle and
 * calculate the number of times it results in the guard going in a loop.
 * 
 * I did a brute force solution, but it was wrong, so I decided to only test 
 * cases where an obstacle was adjacent on or the path of the guard would 
 * traverse... still incorrect solution.
 *
 * Had a problem. I think you have to account for right right where the guard 
 * does not take a step forward... maybe.
 *
 * Yeah I had to account for tiny boxes and strips.
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
    initCoords := []int {0, 0}

    for i, line := range lines {
        for j, chr := range line {
            if chr == '^' {
                initCoords = []int {i, j}
                break
            }
        }

        if initCoords[0] != 0 && initCoords[1] != 0 {
            break
        }

    } // End for.

    // Now we have to count the path.
    // Calculate solution.
    coords := [] int {initCoords[0], initCoords[1]}
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

        // Mark.
        lines[coords[0]][coords[1]] = 'X'

        // Move direction.
        switch dir {
            case Left:
                if lines[coords[0]][coords[1] - 1] != '#' {
                    coords[1] -= 1
                }
                break
            case Down:
                if lines[coords[0] + 1][coords[1]] != '#' {
                    coords[0] += 1
                }
                break
            case Up:
                if lines[coords[0] - 1][coords[1]] != '#' {
                    coords[0] -= 1
                }
                break
            case Right:
                if lines[coords[0]][coords[1] + 1] != '#' {
                    coords[1] += 1
                }
                break
            default:
                log.Fatal("Direction not defined correctly")
        }

        lines[coords[0]][coords[1]] = 'X'

    } // End for.

    // Need to save a copy of array.
    savedArr := make([][]byte, len(lines))

    for i := range lines {
        savedArr[i] = make([]byte, len(lines[i]))
        copy(savedArr[i], lines[i])
    }

    
    coords = [] int {initCoords[0], initCoords[1]}

    //for _, line := range lines {
    //    fmt.Println(string(line))
    //}

    //fmt.Println()

    // Now only test obstructions adjacent to or on 'X's.
    for i, line := range lines {
        for j := range line {

            // Need to copy values.
            for k := range savedArr {
                copy(lines[k], savedArr[k])
            }

            if i == initCoords[0] && j == initCoords[1] {
                continue
            }

            if lines[i][j] == '#' {
                continue
            }

            if lines[i][j] != 'X' {
                if !(i > 0 && lines[i - 1][j] == 'X') &&
                    !(i < len(lines) - 1 && lines[i + 1][j] == 'X') &&
                    !(j > 0 && lines[i][j - 1] == 'X') &&
                    !(j < len(line) - 1 && lines[i][j + 1] == 'X') {
                    continue
                }
            }

            // Run scenario.
            lines[i][j] = '#'
            coords := [] int {initCoords[0], initCoords[1]}
            var dir Dir = Up
            found := false

            // Account for tiny box.
            count := 0

            // The guard cannot be on the outside ring, right?
            for coords[0] != 0 && 
                coords[0] != len(lines) - 1 &&
                coords[1] != 0 &&
                coords[1] != len(lines[0]) - 1 {
                
                // Account for tiny box.
                pos1 := coords[0]
                pos2 := coords[1]

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

                // Move direction.
                switch dir {
                    case Left:

                        // Test if we have been here before.
                        if lines[coords[0]][coords[1]] == 'L' {
                            found = true
                        }

                        lines[coords[0]][coords[1]] = 'L'

                        // Check for block.
                        if lines[coords[0]][coords[1] - 1] != '#' {
                            coords[1] -= 1
                        }
                        break
                    case Down:
                        if lines[coords[0]][coords[1]] == 'D' {
                            found = true
                        }

                        lines[coords[0]][coords[1]] = 'D'

                        if lines[coords[0] + 1][coords[1]] != '#' {
                            coords[0] += 1
                        }
                        break
                    case Up:
                        if lines[coords[0]][coords[1]] == 'U' {
                            found = true
                        }

                        lines[coords[0]][coords[1]] = 'U'

                        if lines[coords[0] - 1][coords[1]] != '#' {
                            coords[0] -= 1
                        }
                        break
                    case Right:
                        if lines[coords[0]][coords[1]] == 'R' {
                            found = true
                        }

                        lines[coords[0]][coords[1]] = 'R'

                        if lines[coords[0]][coords[1] + 1] != '#' {
                            coords[1] += 1
                        }
                        break
                    default:
                        log.Fatal("Direction not defined correctly")
                }

                // Account for tiny box.
                if pos1 == coords[0] && pos2 == coords[1] {
                    count += 1
                }
                
                if found  || count > 100 {
                    solution += 1

                    //for _, pLine := range lines {
                    //    fmt.Println(string(pLine))
                    //}

                    //fmt.Println()
                    break
                }

            } // End for.

        } // End for.

    } // End for.


    // Print solution.
    fmt.Println("Day 6 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


