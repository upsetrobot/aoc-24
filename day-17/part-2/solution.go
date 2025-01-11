/**
 ******************************************************************************
 * Advent of Code 2024 - Day 17 Part 2
 *
 * For history, see old files.
 *
 * I found a partial solution using math, but it goes wrong. I optimized
 * brute-forced from there (was off by like 100 billion or something). It was
 * close, but it should have been closer. I just let it run till I found the
 * answer which took a couple hours. The code is incomplete, needs cleaning,
 * and I need a better solution. I also don't think the solution is entirely
 * general even if it works on the example and input. More work is needed here,
 * but I am moving on for now.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        09 Jan 2025
 * copyright:   2025. All rights reserved.
 *
 ******************************************************************************
 */

package main

// Imports.
import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Types.

type State struct {
    a int
    b int
    c int
    ip int
    prog [] int
    out [] int
    len int
}


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
    words := [] string {}
    state := State {}
    
    for _, line := range strLines {
        words = strings.Split(line, " ")

        if len(words) != 0 {
            if words[0] == "Register" {
                switch words[1] {
                    case "A:":
                        state.a, err = strconv.Atoi(words[2])
                        if err != nil {
                            log.Fatal("Conversion error.")
                        }
                        break

                    case "B:":
                        state.b, err = strconv.Atoi(words[2])
                        if err != nil {
                            log.Fatal("Conversion error.")
                        }
                        break
                        
                    case "C:":
                        state.c, err = strconv.Atoi(words[2])
                        if err != nil {
                            log.Fatal("Conversion error.")
                        }
                        break

                    default:
                        log.Fatal("Parsing error.")
                }

            } else if words[0] == "Program:" {
                words = strings.Split(words[1], ",")

                for _, word := range words {
                    num, err := strconv.Atoi(word)
                    if err != nil {
                        log.Fatal("Program conversion error.")
                    }

                    state.prog = append(state.prog, num)
                }

            } // End if.

        } // End if.

    } // End for.

    // Run program.
    solution := 0

    // Part 2.
    // Okay, I need number of digits. That's len.
    state.len = len(state.prog)
    state.a = 0
    min := 0

    // Okay, I need to check each digit, starting from the end.
    for i := state.len - 1; i >= 0; i-- {
    // Maybe I can do it forward?
    //for i := 0; i < state.len - 1; i++ {

        // Check if last digit is the same.
        // Minimum number is 8**len.
        min = int(math.Pow(8, float64(i)))
        start := state
        startA := 0

        fmt.Println("i:", i)
        fmt.Println("min:", min)
    
        // Should happen within 8 steps.
        for j := 0; j < min * 8; j++ {
            start = state
            startA = solution + j * min
            start.a = startA
            runState(&start)

            fmt.Println("Current solution:", startA, "State: ", start.out)

            if len(start.out) - 1 >= i && start.out[i] == start.prog[i] {
                keepGoing := false

                for k := i; k < state.len; k++ {
                    if start.out[k] != start.prog[k] {
                        keepGoing = true
                    }
                }
                if !keepGoing {
                    break
                }
            }
        }

        solution = startA

        fmt.Println()
        fmt.Println("State selected:    ", start.out)
        fmt.Println("Program:           ", state.prog)
        fmt.Println("Current solution:  ", solution)
        fmt.Println()
        
    } // End for.

    // Brute rest.
    // Well, its close. It doesn't seem to be universal though.
    // If still off track, brute.
    solution = 202366924935110
    start := state

    for l := 0; l < 10000000; l++ {
        start = state
        startA := solution
        start.a = startA
        runState(&start)
        flag := false

        fmt.Println("Current solution:", startA, "State: ", start.out)

        solution -= 1

        for m := 0; m < state.len; m++ {
            if start.out[m] != start.prog[m] {
                flag = true
            }
        }

        if !flag {
            fmt.Println("*************************************************")
            break
        }
    }

    fmt.Println()
    fmt.Println("State selected:    ", start.out)
    fmt.Println("Program:           ", state.prog)
    fmt.Println("Current solution:  ", solution)
    fmt.Println()

    fmt.Println()

    // Print solution.
    fmt.Println("Day 17 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


func runState(state *State) {

    literal := state.prog[state.ip + 1]

    switch state.prog[state.ip] {
        case 0:
            switch literal {
                case 4:
                    literal = state.a
                    break

                case 5:
                    literal = state.b
                    break

                case 6:
                    literal = state.c
                    break

                default:
            }

            state.a = state.a / (1 << literal)
            state.ip += 2
            break;

        case 1:
            state.b = state.b ^ literal
            state.ip += 2
            break

        case 2:
            switch literal {
                case 4:
                    literal = state.a
                    break

                case 5:
                    literal = state.b
                    break

                case 6:
                    literal = state.c
                    break

                default:
            }

            state.b = literal % 8
            state.ip += 2
            break

        case 3:
            if state.a != 0 {
                state.ip = state.prog[state.ip + 1]

            } else {
                state.ip += 2
            }

            break

        case 4:
            state.b = state.b ^ state.c
            state.ip += 2
            break

        case 5:
            switch literal {
                case 4:
                    literal = state.a
                    break

                case 5:
                    literal = state.b
                    break

                case 6:
                    literal = state.c
                    break

                default:
            }

            literal = literal % 8
            state.out = append(state.out, literal)
            state.ip += 2
            break

        case 6:
            switch literal {
                case 4:
                    literal = state.a
                    break

                case 5:
                    literal = state.b
                    break

                case 6:
                    literal = state.c
                    break

                default:
            }

            state.b = state.a / (1 << literal)
            state.ip += 2
            break

        case 7:
            switch literal {
                case 4:
                    literal = state.a
                    break

                case 5:
                    literal = state.b
                    break

                case 6:
                    literal = state.c
                    break

                default:
            }

            state.c = state.a / (1 << literal)
            state.ip += 2
            break

        default:
            log.Fatal("Invalid opcode.")
        
    } // End switch.

    if state.ip >= state.len { 
        return
    }

    runState(state)

} // End runState.


