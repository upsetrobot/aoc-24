/**
 ******************************************************************************
 * Advent of Code 2024 - Day 8 Part 2
 *
 * In this one, we have to calculate the number of antinodes formed from pairs
 * of nodes. The antinodes are at a distance away from the nodes equal to the
 * distance between two nodes in a line and in the direction of the line.
 * So, I guess we have to check each antenna, check each other antenna,
 * save coordinates to a list or mark them (probably make a copy), then add
 * them up.
 *
 * Part II just seems to expand antinodes to the ends of the map and include 
 * the active antennas.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        17 Dec 2024
 * copyright:   2024. All rights reserved.
 *
 ******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)


type Thing struct {
    data rune
    row int
    col int
}


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
    var list [] Thing

    for i, str := range strLines {
        for j, pos := range str {
            if unicode.IsLetter(pos) || unicode.IsDigit(pos) {
                list = append(list, Thing {pos, i, j})
            }

        } // End for.
        
    } // End for.

    sort.Slice(list, func(i, j int) bool {
        return list[i].data < list[j].data 
    })

    chr := '.'
    prevChr := '.'
    row := 0
    col := 0

    for i, preRec := range list {
        chr = preRec.data
        row = preRec.row
        col = preRec.col
        
        for _, rec := range list[i + 1:] { 
            if chr != rec.data {
                if chr == prevChr {
                    newStr, err := replaceAtIndex(strLines[row], col, '#')
                    if err != nil {
                        fmt.Println("Error:", err)
                        return
                    }

                    strLines[row] = newStr
                }
                break
            }

            newStr, err := replaceAtIndex(strLines[row], col, '#')
            if err != nil {
                fmt.Println("Error:", err)
                return
            }

            strLines[row] = newStr

            rowDiff := rec.row - row
            colDiff := rec.col - col
            tmpRow := row
            tmpCol := col
            flag1 := true
            flag2 := true

            for mult1, mult2 := 2, 1; 
                flag1 || flag2; 
                mult1, mult2 = mult1 + 1, mult2 + 1 {

                mark1Row := tmpRow + rowDiff * mult1
                mark1Col := tmpCol + colDiff * mult1
                mark2Row := tmpRow - rowDiff * mult2
                mark2Col := tmpCol - colDiff * mult2

                if mark1Row >= 0 && mark1Row < len(strLines) && 
                    mark1Col >= 0 && mark1Col < len(strLines[0]) {

                    newStr, err := replaceAtIndex(strLines[mark1Row], mark1Col, '#')
                    if err != nil {
                        fmt.Println("Error:", err)
                        return
                    }

                    strLines[mark1Row] = newStr

                } else {
                    flag1 = false
                }

                if mark2Row >= 0 && mark2Row < len(strLines) && 
                    mark2Col >= 0 && mark2Col < len(strLines[0]) {

                    newStr, err := replaceAtIndex(strLines[mark2Row], mark2Col, '#')
                    if err != nil {
                        fmt.Println("Error:", err)
                        return
                    }

                    strLines[mark2Row] = newStr

                } else {
                    flag2 = false
                }

            } // End for.

        } // End for.

        prevChr = chr

    } // End for.

    // Solution.
    for _, line := range strLines {
        for _, pos := range line {
            if pos == '#' {
                solution += 1
            }
        }
    }

    solution += 1

    // Print solution.
    fmt.Println("Day 8 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println("Data:")
    
    for _, line := range strLines {
        fmt.Println(line)
    }

    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


/**
 * Function replaces a character at a given index with a new given character.
 * 
 * Arguments:
 *     s:string     String to replace character in.
 *     index:int    Index to replace character for.
 *     newChar:rune New character.
 * 
 * Returns:
 *     string  New string with character at the given index replaced with the
 *             given new character.
 *     err     Error if index is out of bounds.
 */
func replaceAtIndex(s string, index int, newChar rune) (string, error) {

	// Convert the string to a slice of runes
	runes := []rune(s)

	// Check if the index is valid
	if index < 0 || index >= len(runes) {
        return "", fmt.Errorf("index out of bounds")
	}

	// Replace the character at the specified index
	runes[index] = newChar

	// Convert the slice of runes back to a string
	return string(runes), nil
}


