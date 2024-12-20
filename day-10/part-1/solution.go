/**
 ******************************************************************************
 * Advent of Code 2024 - Day 10 Part 1
 *
 * I like this one. You just have to find all the paths from 0 to 9 and count
 * only the nines. I really want to do a recursive solution for this one. I
 * thought it would also be cool to do one where we just maintain a list of
 * all current positions and move them all at once (but the time complexity
 * saving really would not be that much because they each still need to move,
 * so I don't think I will do that). Anyway, easy. Because the paths aren't
 * important, I think I will just add the nines to a list and count only the
 * unique ones at the end.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        19 Dec 2024
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

// Types.
type Coord struct {
    x int
    y int
    chr rune
}


type Dir int


// Constants.
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

    // Process file.
    // Process file.
    strLines := strings.Split(string(file), "\n")
    strLines = strLines[:len(strLines) - 1]
    solution := 0
    
    for i, line := range strLines {
        for j, pos := range line {
            var coordList [] Coord

            if pos == '0' {
                coordList = getPath(strLines, Coord{j, i, pos}, coordList)

                // Remove duplicates
                uniqueList := make(map[string] Coord)

                for _, pos := range coordList {
                    key := fmt.Sprintf("%d,%d", pos.x, pos.y)
                    uniqueList[key] = pos
                }

                solution += len(uniqueList)
            }
        }
    }

    // Print solution.
    fmt.Println("Day 10 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


/**
 * Function that records positions of nines in a list.
 */
func getPath(mapThing [] string, pos Coord, list [] Coord) [] Coord {
        
    prev, err := strconv.Atoi(string(pos.chr))
    if err != nil {
        log.Fatal("Issue with non-numerical character")
    }

    left := pos
    down := pos
    up := pos
    right := pos

    left.x--
    down.y++
    up.y--
    right.x++
        
    if left.x >= 0 && 
        left.x < len(mapThing[0]) && 
        left.y >= 0 && 
        left.y < len(mapThing) {

        left.chr = rune(mapThing[left.y][left.x])

        curr, err := strconv.Atoi(string(left.chr))
        if err != nil {
            log.Fatal("Issue with non-numerical character")
        }
        
        if curr == prev + 1 {
            if curr == 9 {
                list = append(list, left)

            } else {
                list = getPath(mapThing, left, list)
            }
        }
    }

    if down.x >= 0 && 
        down.x < len(mapThing[0]) && 
        down.y >= 0 && 
        down.y < len(mapThing) {

        down.chr = rune(mapThing[down.y][down.x])

        curr, err := strconv.Atoi(string(down.chr))
        if err != nil {
            log.Fatal("Issue with non-numerical character")
        }
        
        if curr == prev + 1 {
            if curr == 9 {
                list = append(list, down)

            } else {
                list = getPath(mapThing, down, list)
            }
        }
    }

    if up.x >= 0 && 
        up.x < len(mapThing[0]) && 
        up.y >= 0 && 
        up.y < len(mapThing) {

        up.chr = rune(mapThing[up.y][up.x])

        curr, err := strconv.Atoi(string(up.chr))
        if err != nil {
            log.Fatal("Issue with non-numerical character")
        }
        
        if curr == prev + 1 {
            if curr == 9 {
                list = append(list, up)

            } else {
                list = getPath(mapThing, up, list)
            }
        }
    }

    if right.x >= 0 && 
        right.x < len(mapThing[0]) && 
        right.y >= 0 && 
        right.y < len(mapThing) {

        right.chr = rune(mapThing[right.y][right.x])

        curr, err := strconv.Atoi(string(right.chr))
        if err != nil {
            log.Fatal("Issue with non-numerical character")
        }
        
        if curr == prev + 1 {
            if curr == 9 {
                list = append(list, right)

            } else {
                list = getPath(mapThing, right, list)
            }
        }
    }

    return list

} // End getPath






