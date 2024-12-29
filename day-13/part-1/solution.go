/**
 ******************************************************************************
 * Advent of Code 2024 - Day 13 Part 1
 *
 * This one is like an algebra problem. It is about figuring out the minimum 
 * of number of button presses for two different buttons to arrive at a 
 * destination. The word minimum reminds me of recursion. But, I also feel that 
 * this is just solving a system of equations. I will try to find a mathematic 
 * solution.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        29 Dec 2024
 * copyright:   2024. All rights reserved.
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
    ax := 0
    ay := 0
    bx := 0
    by := 0
    px := 0
    py := 0
    aPresses := 0
    bPresses := 0
    tokens := 0
    
    // Parse lines.
    for _, line := range strLines {
        if line == "" {
            continue
        }

        strs := strings.Split(line, " ")

        if strs[0] == "Button" {
            if strs[1] == "A:" {
                ax, err = strconv.Atoi(strs[2][2:len(strs[2]) - 1])
                if err != nil {
                    log.Fatal("Conversion failed.")
                }

                ay, err = strconv.Atoi(strs[3][2:])
                if err != nil {
                    log.Fatal("Conversion failed.")
                }

                continue

            } else if strs[1] == "B:" {
                bx, err = strconv.Atoi(strs[2][2:len(strs[2]) - 1])
                if err != nil {
                    log.Fatal("Conversion failed.")
                }

                by, err = strconv.Atoi(strs[3][2:])
                if err != nil {
                    log.Fatal("Conversion failed.")
                }

                continue

            } else {
                log.Fatal("Parsing error.")
            }

        } else if strs[0] == "Prize:" {
            px, err = strconv.Atoi(strs[1][2:len(strs[1]) - 1])
            if err != nil {
                log.Fatal("Conversion failed.")
            }

            py, err = strconv.Atoi(strs[2][2:])
            if err != nil {
                log.Fatal("Conversion failed.")
            }

            // Solve problem.
            // px = ax*i + bx*j
            // py = ay*i + by*j
            // i and j vary (like x and y with z as solution)
            // t = 3*i + j  
            // min_t = min(t) where px and py meet ax*i+bx*j and ay*i+by*j
            // Okay, lets solve.
            // px - bx*j = ax*i
            // (px - bx*j) / ax = i
            // Now sub.
            // py = ay*((px - bx*j) / ax) + by*j
            // Need to sep j.
            // py = ay*px/ax - ay*bx*j/ax + by*j
            // py = ay*px/ax + (-ay*bx/ax)*j + by*j
            // py = ay*px/ax + ((-ay*bx/ax) + by)*j
            // Now solve for j
            // py - ay*px/ax = ((-ay*bx/ax) + by)*j
            // (py - ay*px/ax) / ((-ay*bx/ax) + by) = j
            // Now if I solve j, then I can solve i. 
            // Probably should used wolfram or something. But I think this 
            // is right. There should be infinite solutions, so that's why 
            // I am concerned. I guess I just need to try it. Also, there are 
            // some simplifications I skipped... oh well.
            // Also, there are invalid solutions. I wonder how that will 
            // show in the math.

            // My solution was close, but under. I think it probably has to do 
            // with the division (integer division). So, I probably need to 
            // correct for that since it was off by one.
            bPresses = (py - ay * px / ax) / ((0 - ay) * bx / ax + by)
            aPresses = (px - bx * bPresses) / ax

            if aPresses < 0 || bPresses < 0 || aPresses > 100 || bPresses > 100 {
                continue
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses++
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses++
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses--
                aPresses--

                bPresses++
            }

            if px != ax * aPresses + bx * bPresses {
                bPresses++
            }

            if px != ax * aPresses + bx * bPresses {
                bPresses--
                bPresses--

                aPresses--
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses--
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses++
                aPresses++

                bPresses--
            }

            if px != ax * aPresses + bx * bPresses {
                bPresses--
            }

            if px != ax * aPresses + bx * bPresses {
                bPresses++
                bPresses++

                aPresses++
                bPresses--
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses++
                bPresses--
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses--
                aPresses--
                bPresses++
                bPresses++

                aPresses--
                bPresses++
            }

            if px != ax * aPresses + bx * bPresses {
                aPresses--
                bPresses++
            }

            if px != ax * aPresses + bx * bPresses {
                // aPresses++
                // aPresses++
                // bPresses--
                // bPresses--

                // log.Fatal("Inbetween.")
                // Maybe this is the problem.
                // Ok, I found two cases where it went wrong. Fixing.
                // Ok, I found another case (increment both). Fixing.
                // Had to check two up and two down.

                // fmt.Printf("ax: %d\n", ax)
                // fmt.Printf("ay: %d\n", ay)
                // fmt.Printf("bx: %d\n", bx)
                // fmt.Printf("by: %d\n", by)
                // fmt.Printf("px: %d\n", px)
                // fmt.Printf("py: %d\n", py)
                // fmt.Printf("aPresses: %d\n", aPresses)
                // fmt.Printf("bPresses: %d\n", bPresses)
                // fmt.Println()

                continue
            }

            if py != ay * aPresses + by * bPresses {
                // log.Fatal("Y not correct.")
                // Damn did not work here. 
                // Negative number.
                // But, I am not sure if this is the problem.

                // fmt.Printf("ax: %d\n", ax)
                // fmt.Printf("ay: %d\n", ay)
                // fmt.Printf("bx: %d\n", bx)
                // fmt.Printf("by: %d\n", by)
                // fmt.Printf("px: %d\n", px)
                // fmt.Printf("py: %d\n", py)
                // fmt.Printf("aPresses: %d\n", aPresses)
                // fmt.Printf("bPresses: %d\n", bPresses)
                // fmt.Println()

                // I don't think this is the problem.
                continue
            }

            tokens = aPresses * 3 + bPresses
            solution += tokens

            // fmt.Printf("ax: %d\n", ax)
            // fmt.Printf("ay: %d\n", ay)
            // fmt.Printf("bx: %d\n", bx)
            // fmt.Printf("by: %d\n", by)
            // fmt.Printf("px: %d\n", px)
            // fmt.Printf("py: %d\n", py)
            // fmt.Printf("aPresses: %d\n", aPresses)
            // fmt.Printf("bPresses: %d\n", bPresses)
            // fmt.Printf("tokens: %d\n", tokens)
            // fmt.Println()

            // Huh, it works. Funny, math works.

        } else {
            log.Fatal("Parsing error.")

        } // End if.

    } // End for.

    // Print solution.
    fmt.Println("Day 13 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


