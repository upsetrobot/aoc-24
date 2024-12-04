/**
 ******************************************************************************
 * Advent of Code 2024 - Day 3 Part 2 
 *
 * This appears to be a parsing problem. I wonder if regex may be a good 
 * library solution for this. This one adds `do` and `dont` instructions 
 * which enable and disable calculation. So, I will have to parse things 
 * one at a time instead of all at once.
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        03 Dec 2024
 * copyright:   2024. All rights reserved.
 * 
 ******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"os"
    "regexp"
    "strconv"
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

    patternArgs := `mul\((\d+),(\d+)\)`
    patternMul := `mul\(\d+,\d+\)`
    patternDo := `do\(\)`
    patternDont := `don't\(\)`

    reArgs, err := regexp.Compile(patternArgs)
    if err != nil {
        log.Fatal("Regex compilation failed:", err)
    }

    reMul, err := regexp.Compile(patternMul)
    if err != nil {
        log.Fatal("Regex compilation failed:", err)
    }

    reDo, err := regexp.Compile(patternDo)
    if err != nil {
        log.Fatal("Regex compilation failed:", err)
    }

    reDont, err := regexp.Compile(patternDont)
    if err != nil {
        log.Fatal("Regex compilation failed:", err)
    }

    reNext, err := regexp.Compile(patternMul + `|` + patternDo + `|` + patternDont)
    if err != nil {
        log.Fatal("Regex compilation failed:", err)
    }

    // Calculate solution.
    solution := 0
    words := string(file)
    start := 0
    enabled := true

    for {
        locNext := reNext.FindStringIndex(words[start:])
        locMul := reMul.FindStringIndex(words[start:])
        locDo := reDo.FindStringIndex(words[start:])
        locDont := reDont.FindStringIndex(words[start:])

        // Check for last of matches.
        if locNext == nil {
            break
        }

        // Determine which type match is.
        if enabled && locMul != nil && locNext[0] == locMul[0] {
            args := reArgs.FindStringSubmatch(words[start:])

            arg1, err := strconv.Atoi(args[1])
            if err != nil {
                log.Fatal("String conversion failed:", err)
            }

            arg2, err := strconv.Atoi(args[2])
            if err != nil {
                log.Fatal("String conversion failed:", err)
            }

            solution += arg1 * arg2

        } else if locDo != nil && locNext[0] == locDo[0] {
            enabled = true

        } else if locDont != nil && locNext[0] == locDont[0] {
            enabled = false
        }

        start += locNext[1]

    } // End for.

    // Print solution.
    fmt.Println("Day 3 Part 2")
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


