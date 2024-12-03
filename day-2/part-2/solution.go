/**
 ******************************************************************************
 * Advent of Code 2024 - Day 2 Part 2 
 *
 * In this one, we have to analyze a set of numbers and determine if they 
 * meet criteria. This one is much more complicated by allowing removal of 
 * one number.
 * 
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        02 Dec 2024
 * copyright:   2024. All rights reserved.
 * 
 ******************************************************************************
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }

    // Process file.
    defer file.Close()

    scanner := bufio.NewScanner(file)
    solution := 0

    for scanner.Scan() {

        // Parse line.
        line := scanner.Text()
        words := strings.Fields(line)
        var tmp []int

        for _, i := range words {
            num, err := strconv.Atoi(i)
            if err != nil {
                log.Fatal("Error converting string to number:", err)
            }

            tmp = append(tmp, num)
        }

        // Determine if report is safe.
        tmpNum := tmp[0]
        dir := 1

        if tmp[1] < tmpNum {
            dir = 0

        }         

        flag := 1

        for _, val := range tmp[1:] {

            // Unsafe?
            if !(dir == 1 && val > tmpNum && (val - tmpNum) >= 1 && (val - tmpNum) <= 3) && 
                !(dir == 0 && val < tmpNum && (tmpNum - val) >= 1 && (tmpNum - val) <= 3) {

                // Now go through each list and test each with a removed element.
                flag2 := 1

                for i := range tmp {

                    // Copy list and remove element.
                    var copyTmp []int

                    for j, val2 := range tmp {
                        if j == i {
                            continue
                        }

                        copyTmp = append(copyTmp, val2)
                    }

                    // Now test list.
                    tmpNum := copyTmp[0]
                    dir = 1
                    flag2 = 1

                    if copyTmp[1] < tmpNum {
                        dir = 0

                    } 

                    for _, val3 := range copyTmp[1:] {
                        if !(dir == 1 && val3 > tmpNum && (val3 - tmpNum) >= 1 && (val3 - tmpNum) <= 3) && 
                            !(dir == 0 && val3 < tmpNum && (tmpNum - val3) >= 1 && (tmpNum - val3) <= 3) {
                            flag2 = 0;
                            break;
                        }

                        tmpNum = val3
                    }

                    // Safe?
                    if flag2 == 1 {
                        break
                    }

                } // End for.

                if flag2 != 1 {
                    flag = 0
                }

                break

            } // End if.
            
            tmpNum = val

        } // End for.

        solution += flag

    } // End for.

    // Print solution.
    fmt.Println("Day 2 Part 2")
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


