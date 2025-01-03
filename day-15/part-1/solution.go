/**
 ******************************************************************************
 * Advent of Code 2024 - Day 15 Part 1
 *
 * This one seems pretty simple. We just have to move boxes around. I think 
 * part 2 is probably where this get hard. There a few approaches. I plan on 
 * taking the naive approach and just move characters around.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        02 Jan 2025
 * copyright:   2025. All rights reserved.
 *
 ******************************************************************************
 */

package main

// Imports.
import (
	"fmt"
	"log"
	"os"
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
    strLines := strings.Split(string(file), "\n")
    strLines = strLines[:len(strLines) - 1]
    solution := 0

    // Make map.
    rows := 0
    columns := len(strLines[0])
    robx := 0
    roby := 0

    // Count number of rows in map.
    for i, line := range strLines {
        if line == "" {
            break
        }

        rows++

        for j, pos := range line {
            if pos == '@' {
                robx = j
                roby = i
            }
        }
    }

    boxMap := make([][] byte, rows)

    for i := range boxMap {
        boxMap[i] = make([] byte, columns)
    }

    // Parse file.
    directionIndex := 0

    for i, line := range strLines {
        if line == "" {
            directionIndex = i + 1
            break
        } 

        for j, pos := range line {
            boxMap[i][j] = byte(pos)
        }
    }

    // Print.
    //for _, line := range boxMap {
    //    for _, pos := range line {
    //        fmt.Print(string(pos))
    //    }

    //    fmt.Print("\n")
    //}

    //fmt.Print("\n")

    // Process moves.
    // Probably should have put all positions in a list, but oh well.
    for i := directionIndex; len(strLines) - i > 0; i++ {
        for _, pos := range strLines[i] {
            switch pos {
                case '<':
                    switch boxMap[roby][robx - 1] {
                        case '#':
                           break

                        case '.':
                            boxMap[roby][robx] = '.'
                            boxMap[roby][robx - 1] = '@'
                            robx--

                        case 'O':

                            // Need to count boxes left.
                            count := 0
                            k := robx

                            for ; k > 0; k-- {
                                // I got a feeling, I'm gonna need this in part
                                // 2.
                                if boxMap[roby][k] == '#' || boxMap[roby][k] == '.' {
                                    break
                                }

                                if boxMap[roby][k] == 'O' {
                                    count++
                                }
                            }

                            if boxMap[roby][k] == '#' {
                                k++
                            }

                            // Make adjustments.
                            for count > 0 {
                                boxMap[roby][k] = 'O'
                                count--
                                k++
                            }

                            boxMap[roby][k] = '@'
                            newRobx := k
                            k++

                            for ; k <= robx; k++ {
                                boxMap[roby][k] = '.'
                            }

                            robx = newRobx

                            break

                        default:
                            log.Fatal("Invalid character.")

                    } // End switch.
                    
                    break

                case 'v':
                    switch boxMap[roby + 1][robx] {
                        case '#':
                           break

                        case '.':
                            boxMap[roby][robx] = '.'
                            boxMap[roby + 1][robx] = '@'
                            roby++

                        case 'O':

                            count := 0
                            k := roby

                            for ; k < len(boxMap); k++ {
                                if boxMap[k][robx] == '#' || boxMap[k][robx] == '.' {
                                    break
                                }

                                if boxMap[k][robx] == 'O' {
                                    count++
                                }
                            }

                            if boxMap[k][robx] == '#' {
                                k--
                            }

                            // Make adjustments.
                            for count > 0 {
                                boxMap[k][robx] = 'O'
                                count--
                                k--
                            }

                            boxMap[k][robx] = '@'
                            newRoby := k
                            k--

                            for ; k >= roby; k-- {
                                boxMap[k][robx] = '.'
                            }

                            roby = newRoby

                            break

                        default:
                            log.Fatal("Invalid character.")

                    } // End switch.

                    break

                case '^':
                    switch boxMap[roby - 1][robx] {
                        case '#':
                           break

                        case '.':
                            boxMap[roby][robx] = '.'
                            boxMap[roby - 1][robx] = '@'
                            roby--

                        case 'O':

                            count := 0
                            k := roby

                            for ; k > 0; k-- {
                                if boxMap[k][robx] == '#' || boxMap[k][robx] == '.' {
                                    break
                                }

                                if boxMap[k][robx] == 'O' {
                                    count++
                                }
                            }

                            if boxMap[k][robx] == '#' {
                                k++
                            }

                            // Make adjustments.
                            for count > 0 {
                                boxMap[k][robx] = 'O'
                                count--
                                k++
                            }

                            boxMap[k][robx] = '@'
                            newRoby := k
                            k++

                            for ; k <= roby; k++ {
                                boxMap[k][robx] = '.'
                            }

                            roby = newRoby

                            break

                        default:
                            log.Fatal("Invalid character.")

                    } // End switch.

                    break

                case '>':
                    switch boxMap[roby][robx + 1] {
                        case '#':
                           break

                        case '.':
                            boxMap[roby][robx] = '.'
                            boxMap[roby][robx + 1] = '@'
                            robx++

                        case 'O':

                            // Need to count boxes left.
                            count := 0
                            k := robx

                            for ; k < len(boxMap[0]); k++ {
                                if boxMap[roby][k] == '#' || boxMap[roby][k] == '.' {
                                    break
                                }

                                if boxMap[roby][k] == 'O' {
                                    count++
                                }
                            }

                            if boxMap[roby][k] == '#' {
                                k--
                            }

                            // Make adjustments.
                            for count > 0 {
                                boxMap[roby][k] = 'O'
                                count--
                                k--
                            }

                            boxMap[roby][k] = '@'
                            newRobx := k
                            k--

                            for ; k >= robx; k-- {
                                boxMap[roby][k] = '.'
                            }

                            robx = newRobx

                            break

                        default:
                            log.Fatal("Invalid character.")

                    } // End switch.
                    break

                default:
                    log.Fatal("Invalid direction.")

            } // End switch.

            // Print.
            //for _, line := range boxMap {
            //    for _, pos := range line {
            //        fmt.Print(string(pos))
            //    }

            //    fmt.Print("\n")
            //}

            //fmt.Print("\n")

        } // End for.

    } // End for.

    // Get solution.
    for i, line := range boxMap {
        for j, pos := range line {
            if pos == 'O' {
                solution += 100 * i 
                solution += j
            }
        }
    }

    // Print solution.
    fmt.Println("Day 15 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


