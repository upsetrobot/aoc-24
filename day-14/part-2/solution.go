/**
 ******************************************************************************
 * Advent of Code 2024 - Day 14 Part 2
 *
 * This one is a little more abstract. We receive a list of locations and
 * velocities of robots. We also receive the size of a map and some rules that
 * robots can circle the map. We have to calculate the number of robots in each
 * quadrant after 100 seconds.
 *
 * Part 1 was easy. Part 2 sounds ridiculous! You have to find the lowest
 * number of seconds when the robots' positions form a Christmas Tree. It
 * probably like 1 or something stupid. I think I will start will a key press
 * solution. Otherwise, I have to define "Christmas tree".
 *
 * I got lucky on this one. I saw that some of the robots were gathered
 * together horizontally or vertically sometimes, so I decided to run the
 * simulation and print when the first row is blank (first I did two blank
 * lines, but that was running to long). A better solution might have been to
 * count the number of bots in each row and determine is so many bots are just
 * a few rows. But my visual identification solution worked in just less than
 * 250 key presses (which could have been better), but I am lazy.
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
	"bufio"
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

    // Define reader.
    reader := bufio.NewReader(os.Stdin)

    // Print positions.
    buffer := make([]string, 103)

    for i := range buffer {
        buffer[i] = strings.Repeat(".", 101)
    }

    makeMap(robotList, buffer)

    for _, line := range buffer {
        fmt.Println(line)
    }

    fmt.Println()
    fmt.Println("Current seconds: ", solution)

    fmt.Print("Press enter...")

    // Wait for key press
    _, err = reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Simulate 100 seconds.
    for j := 0; j < 100000000; j++ {
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
        }

        solution++
        makeMap(robotList, buffer)

        // If the first row is blank, print.
        if buffer[0] == strings.Repeat(".", 101) {

            for _, line := range buffer {
                fmt.Println(line)
            }

            fmt.Println()
            fmt.Println("Current seconds: ", solution)
            fmt.Println()
            fmt.Print("Press enter...")

            // Wait for key press
            _, err = reader.ReadString('\n')
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
        }

    } // End for.

    // Print solution.
    fmt.Println("Day 14 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


func makeMap(robotList [] Robot, buffer []string) {
    for i := range buffer {
        buffer[i] = strings.Repeat(".", 101)
    }

    for _, robot := range robotList {
        buffer[robot.py] = buffer[robot.py][0:robot.px] + "0" + buffer[robot.py][robot.px + 1:]
    }
}


