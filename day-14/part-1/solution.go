/**
 ******************************************************************************
 * Advent of Code 2024 - Day 14 Part 1
 *
 * This one is a little more abstract. We receive a list of locations and
 * velocities of robots. We also receive the size of a map and some rules that
 * robots can circle the map. We have to calculate the number of robots in each
 * quadrant after 100 seconds.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        01 Jan 2025
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
	"strconv"
	"strings"
)


// Types.

type Robot struct {
    px int
    py int
    vx int
    vy int
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
    solution := 0
    var robotList [] Robot
    
    for _, line := range strLines {
        words := strings.Split(line, " ")
        positions := strings.Split(words[0], ",")
        velocities := strings.Split(words[1], ",")
        
        px, err := strconv.Atoi(positions[0][2:])
        if err != nil {
            log.Fatal("Conversion failed.")
        }

        py, err := strconv.Atoi(positions[1])
        if err != nil {
            log.Fatal("Conversion failed.")
        }

        vx, err := strconv.Atoi(velocities[0][2:])
        if err != nil {
            log.Fatal("Conversion failed.")
        }

        vy, err := strconv.Atoi(velocities[1])
        if err != nil {
            log.Fatal("Conversion failed.")
        }

        robotList = append(robotList, Robot{px, py, vx, vy})

    } // End for.

    // Simulate 100 seconds.
    for j := 0; j < 100; j++ {
        for i, robot := range robotList {
            robotList[i].px += robot.vx
            robotList[i].px = robotList[i].px % 101

            if robotList[i].px < 0 {
                robotList[i].px = 101 + robotList[i].px
            }

            robotList[i].py += robot.vy
            robotList[i].py = robotList[i].py % 103

            if robotList[i].py < 0 {
                robotList[i].py = 103 + robotList[i].py
            }

            fmt.Println("x", robot.px)
            fmt.Println("y", robot.py)
            fmt.Println()
        }

        fmt.Println("----")
    }

    // Count quadrants.
    q1 := 0
    q2 := 0
    q3 := 0
    q4 := 0

    for _, robot := range robotList {
        if robot.px < 101 / 2 {
            if robot.py < 103 /2 {
                q1++
                
            } else if robot.py > 103 / 2 {
                q2++
            }

        } else if robot.px > 101 / 2 {
            if robot.py < 103 /2 {
                q3++
                
            } else if robot.py > 103 / 2 {
                q4++
            }
        }
    }

    // Calculate solution.
    solution = q1 * q2 * q3 * q4

    // Print solution.
    fmt.Println("Day 14 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


