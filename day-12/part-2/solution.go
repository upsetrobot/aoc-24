/**
 ******************************************************************************
 * Advent of Code 2024 - Day 12 Part 2
 *
 * This one does not seem to bad, although I imagine some hiccups will happen.
 * You have to find each region, calculate its area and its perimeter and 
 * multiply those together to find the solution.
 *
 * This is only complicated by changing perimeter to number of sides. Seems 
 * annoying, really. 
 *
 * My solution is really close. But it is susceptible to counting walls more 
 * than once. Which is very frustrating. I think a better solution is probably 
 * one that uses some sort of matrix math like pixel algorithms like
 * convolution. Alternatively, I may still be able to get this to work by going
 * around the outline, but then I still have to deal with middle stuff and one
 * wide things. I could also pass everything twice or multiple times to try to
 * pretend being on the line between fields. Anyway, back to the drawing
 * board... no pun intended.
 *
 * I finally fixed the issues I was having with double-counting walls at the 
 * end. All tests passed, but solution is incorrect. There must be a case 
 * were I am double-counting a wall other than the first position walls and 
 * not at the last position even though my algorithm is supposed to go around 
 * the outer edges before exploring the middle. My suspicion is that there is 
 * a wall in the middle that is being hit once, then another branch is hitting 
 * that same wall. My suspicion is that I need to account for double counting 
 * with every branch.... Ugh. I can test this theory I'm sure, but I will 
 * still need to move my double-counting logic into the recursive function 
 * which sucks. Oh, well, lemme make a test case for my theory and start 
 * moving the logic into the primary branching function.
 *
 * Okay, it works great... but it giving a low solution now. The tests all 
 * pass. So, now it is not counting a wall that it should be somewhere - 
 * somehow. 
 *
 * Finally, just a little error that I fixed before but forgot when I changed 
 * things. I could clean up the code to simplify it now that the approach 
 * makes sense, but I'm lazy. I also think a convolution solution might be 
 * possible... but, whatever.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        25 Dec 2024
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
	"strings"
)


// Types.
type Value struct {
    perimeter int
    area int
    lastPosition Coord
    dir int
    sideList [] int
}


type Coord struct {
    y int
    x int
}


// Constants.
// None.


// Variables.
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
    var bytesLines [][] byte
    sideList := [] int {}

    for _, line := range strLines {
        bytesLines = append(bytesLines, [] byte(line))
    }

    solution := 0
    
    // Lets go through each letter and we need to determine if we have already 
    // done it yet somehow. I'm gonna try a marker.

    for i, line := range bytesLines {
        for j, pos := range line {
            
            // Check marker.
            if pos > 'Z' {
                continue
            }

            val := getValues(bytesLines, Coord{i, j}, sideList, Coord{i, j}, 1)
            solution += val.area * val.perimeter

        } // End for.

    } // End for.

    // Print solution.
    fmt.Println("Day 12 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


func getValues(mapThing [][] byte, pos Coord, sideList [] int, start Coord, dir int) Value {

    chr := mapThing[pos.y][pos.x]
    area := 1
    perimeter := 0
    mapThing[pos.y][pos.x] = chr + 32

    // Count perimeter.
    // Check left.
    if pos.x == 0 || 
        (mapThing[pos.y][pos.x - 1] != chr && mapThing[pos.y][pos.x - 1] != chr + 32) {
        if !contains(sideList, 0) {
            sideList = append(sideList, 0)
            perimeter += 1
        }
    } else {
        sideList = remove(sideList, 0)
    }

    // Check down.
    if pos.y == len(mapThing) - 1 || 
        (mapThing[pos.y + 1][pos.x] != chr && mapThing[pos.y + 1][pos.x] != chr + 32) {
        if !contains(sideList, 1) {
            sideList = append(sideList, 1)
            perimeter += 1
        }
    } else {
        sideList = remove(sideList, 1)
    }

    // Check up.
    if pos.y == 0 || 
        (mapThing[pos.y - 1][pos.x] != chr && mapThing[pos.y - 1][pos.x] != chr + 32) {
        if !contains(sideList, 2) {
            sideList = append(sideList, 2)
            perimeter += 1
        }
    } else {
        sideList = remove(sideList, 2)
    }

    // Check right.
    if pos.x == len(mapThing[0]) - 1 || 
        (mapThing[pos.y][pos.x + 1] != chr && mapThing[pos.y][pos.x + 1] != chr + 32) {
        if !contains(sideList, 3) {
            sideList = append(sideList, 3)
            perimeter += 1
        }

    } else {
        sideList = remove(sideList, 3)
    }

    ret := Value{perimeter, area, pos, dir, sideList}

    // Okay, I am gonna try putting my double-counting wall checks here.
    lChr := mapThing[pos.y][pos.x]

    left := byte('.')
    leftCoords := Coord{pos.y, pos.x - 1}
    leftSides := [] int {}

    down := byte('.')
    downCoords := Coord{pos.y + 1, pos.x}
    downSides := [] int {}

    up := byte('.')
    upCoords := Coord{pos.y - 1, pos.x}
    upSides := [] int {}

    right := byte('.')
    rightCoords := Coord{pos.y, pos.x + 1}
    rightSides := [] int {}

    // Get chars.
    if leftCoords.x >= 0 {
        left = mapThing[leftCoords.y][leftCoords.x]
    }

    if downCoords.y < len(mapThing) {
        down = mapThing[downCoords.y][downCoords.x]
    }

    if upCoords.y >= 0 {
        up = mapThing[upCoords.y][upCoords.x]
    }
    
    if rightCoords.x < len(mapThing[0]) {
        right = mapThing[rightCoords.y][rightCoords.x]
    }

    // Get sides.
    if dir != 3 && left == lChr {
        if leftCoords.x == 0 || !(mapThing[leftCoords.y][leftCoords.x - 1] == left || mapThing[leftCoords.y][leftCoords.x - 1] == chr) {
            leftSides = append(leftSides, 0)
        }

        if leftCoords.y == len(mapThing) - 1 || !(mapThing[leftCoords.y + 1][leftCoords.x] == left || mapThing[leftCoords.y + 1][leftCoords.x] == chr) {
            leftSides = append(leftSides, 1)
        }

        if leftCoords.y == 0 || !(mapThing[leftCoords.y - 1][leftCoords.x] == left || mapThing[leftCoords.y - 1][leftCoords.x] == chr) {
            leftSides = append(leftSides, 2)
        }

        if leftCoords.x == len(mapThing[0]) - 1 || !(mapThing[leftCoords.y][leftCoords.x + 1] == left || mapThing[leftCoords.y][leftCoords.x + 1] == chr) {
            leftSides = append(leftSides, 3)
        }
    }

    if dir != 2 && down == lChr {
        if downCoords.x == 0 || !(mapThing[downCoords.y][downCoords.x - 1] == down || mapThing[downCoords.y][downCoords.x - 1] == chr) {
            downSides = append(downSides, 0)
        }

        if downCoords.y == len(mapThing) - 1 || !(mapThing[downCoords.y + 1][downCoords.x] == down || mapThing[downCoords.y + 1][downCoords.x] == chr) {
            downSides = append(downSides, 1)
        }

        if downCoords.y == 0 || !(mapThing[downCoords.y - 1][downCoords.x] == down || mapThing[downCoords.y - 1][downCoords.x] == chr) {
            downSides = append(downSides, 2)
        }

        if downCoords.x == len(mapThing[0]) - 1 || !(mapThing[downCoords.y][downCoords.x + 1] == down || mapThing[downCoords.y][downCoords.x + 1] == chr) {
            downSides = append(downSides, 3)
        }
    }

    if dir != 1 && up == lChr {
        if upCoords.x == 0 || !(mapThing[upCoords.y][upCoords.x - 1] == up || mapThing[upCoords.y][upCoords.x - 1] == chr) {
            upSides = append(upSides, 0)
        }

        if upCoords.y == len(mapThing) - 1 || !(mapThing[upCoords.y + 1][upCoords.x] == up || mapThing[upCoords.y + 1][upCoords.x] == chr) {
            upSides = append(upSides, 1)
        }

        if upCoords.y == 0 || !(mapThing[upCoords.y - 1][upCoords.x] == up || mapThing[upCoords.y - 1][upCoords.x] == chr) {
            upSides = append(upSides, 2)
        }

        if upCoords.x == len(mapThing[0]) - 1 || !(mapThing[upCoords.y][upCoords.x + 1] == up || mapThing[upCoords.y][upCoords.x + 1] == chr) {
            upSides = append(upSides, 3)
        }
    }

    if dir != 0 && right == lChr {
        if rightCoords.x == 0 || !(mapThing[rightCoords.y][rightCoords.x - 1] == right || mapThing[rightCoords.y][rightCoords.x - 1] == chr) {
            rightSides = append(rightSides, 0)
        }

        if rightCoords.y == len(mapThing) - 1 || !(mapThing[rightCoords.y + 1][rightCoords.x] == right || mapThing[rightCoords.y + 1][rightCoords.x] == chr) {
            rightSides = append(rightSides, 1)
        }

        if rightCoords.y == 0 || !(mapThing[rightCoords.y - 1][rightCoords.x] == right || mapThing[rightCoords.y - 1][rightCoords.x] == chr) {
            rightSides = append(rightSides, 2)
        }

        if rightCoords.x == len(mapThing[0]) - 1 || !(mapThing[rightCoords.y][rightCoords.x + 1] == right || mapThing[rightCoords.y][rightCoords.x + 1] == chr) {
            rightSides = append(rightSides, 3)
        }
    }

    // Account for double-count.
    for _, side := range leftSides {
        for _, curr := range sideList {
            if curr == side {
                ret.perimeter--
            }
        }
    }

    for _, side := range downSides {
        for _, curr := range sideList {
            if curr == side {
                ret.perimeter--
            }
        }
    }

    for _, side := range upSides {
        for _, curr := range sideList {
            if curr == side {
                ret.perimeter--
            }
        }
    }

    for _, side := range rightSides {
        for _, curr := range sideList {
            if curr == side {
                ret.perimeter--
            }
        }
    }

    // Spread.
    // Okay, what if we could always turn left (like a maze) - then spread 
    // to the middle. Our order of directions would depend on the direction 
    // we were going. We will still have the issue of counting one extra 
    // sometimes, but it should not be more than that I think. We could use 
    // a mode to determine if we reached the end of left turning. hmm....
    // Let's try that.

    // This solved one problem, but I'm still having problems.

    switch dir {

        // Going left.
        case 0:

            // Up.
            if pos.y > 0 && mapThing[pos.y - 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 2)
                copySideList = remove(copySideList, 1)
                valUp := getValues(mapThing, Coord{pos.y - 1, pos.x}, copySideList, start, 2)
                ret.area += valUp.area
                ret.perimeter += valUp.perimeter
                ret.lastPosition = valUp.lastPosition
                ret.dir = valUp.dir
                ret.sideList = valUp.sideList
            }

            // Left.
            if pos.x > 0 && mapThing[pos.y][pos.x - 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 0)
                copySideList = remove(copySideList, 3)
                valLeft := getValues(mapThing, Coord{pos.y, pos.x - 1}, copySideList, start, 0)
                ret.area += valLeft.area
                ret.perimeter += valLeft.perimeter
                ret.lastPosition = valLeft.lastPosition
                ret.dir = valLeft.dir
                ret.sideList = valLeft.sideList
            }

            // Down.
            if pos.y < len(mapThing) - 1 && mapThing[pos.y + 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 1)
                copySideList = remove(copySideList, 2)
                valDown := getValues(mapThing, Coord{pos.y + 1, pos.x}, copySideList, start, 1)
                ret.area += valDown.area
                ret.perimeter += valDown.perimeter
                ret.lastPosition = valDown.lastPosition
                ret.dir = valDown.dir
                ret.sideList = valDown.sideList
            }

            // Right.
            if pos.x < len(mapThing[0]) - 1 && mapThing[pos.y][pos.x + 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 3)
                copySideList = remove(copySideList, 0)
                valRight := getValues(mapThing, Coord{pos.y, pos.x + 1}, copySideList, start, 3)
                ret.area += valRight.area
                ret.perimeter += valRight.perimeter
                ret.lastPosition = valRight.lastPosition
                ret.dir = valRight.dir
                ret.sideList = valRight.sideList
            }

            break

        // Going Down.
        case 1:

            // Left.
            if pos.x > 0 && mapThing[pos.y][pos.x - 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 0)
                copySideList = remove(copySideList, 3)
                valLeft := getValues(mapThing, Coord{pos.y, pos.x - 1}, copySideList, start, 0)
                ret.area += valLeft.area
                ret.perimeter += valLeft.perimeter
                ret.lastPosition = valLeft.lastPosition
                ret.dir = valLeft.dir
                ret.sideList = valLeft.sideList
            }

            // Down.
            if pos.y < len(mapThing) - 1 && mapThing[pos.y + 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 1)
                copySideList = remove(copySideList, 2)
                valDown := getValues(mapThing, Coord{pos.y + 1, pos.x}, copySideList, start, 1)
                ret.area += valDown.area
                ret.perimeter += valDown.perimeter
                ret.lastPosition = valDown.lastPosition
                ret.dir = valDown.dir
                ret.sideList = valDown.sideList
            }

            // Right.
            if pos.x < len(mapThing[0]) - 1 && mapThing[pos.y][pos.x + 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 3)
                copySideList = remove(copySideList, 0)
                valRight := getValues(mapThing, Coord{pos.y, pos.x + 1}, copySideList, start, 3)
                ret.area += valRight.area
                ret.perimeter += valRight.perimeter
                ret.lastPosition = valRight.lastPosition
                ret.dir = valRight.dir
                ret.sideList = valRight.sideList
            }

            // Up.
            if pos.y > 0 && mapThing[pos.y - 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 2)
                copySideList = remove(copySideList, 1)
                valUp := getValues(mapThing, Coord{pos.y - 1, pos.x}, copySideList, start, 2)
                ret.area += valUp.area
                ret.perimeter += valUp.perimeter
                ret.lastPosition = valUp.lastPosition
                ret.dir = valUp.dir
                ret.sideList = valUp.sideList
            }

            break

        // Going up.
        case 2:

            // Right.
            if pos.x < len(mapThing[0]) - 1 && mapThing[pos.y][pos.x + 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 3)
                copySideList = remove(copySideList, 0)
                valRight := getValues(mapThing, Coord{pos.y, pos.x + 1}, copySideList, start, 3)
                ret.area += valRight.area
                ret.perimeter += valRight.perimeter
                ret.lastPosition = valRight.lastPosition
                ret.dir = valRight.dir
                ret.sideList = valRight.sideList
            }

            // Up.
            if pos.y > 0 && mapThing[pos.y - 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 2)
                copySideList = remove(copySideList, 1)
                valUp := getValues(mapThing, Coord{pos.y - 1, pos.x}, copySideList, start, 2)
                ret.area += valUp.area
                ret.perimeter += valUp.perimeter
                ret.lastPosition = valUp.lastPosition
                ret.dir = valUp.dir
                ret.sideList = valUp.sideList
            }

            // Left.
            if pos.x > 0 && mapThing[pos.y][pos.x - 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 0)
                copySideList = remove(copySideList, 3)
                valLeft := getValues(mapThing, Coord{pos.y, pos.x - 1}, copySideList, start, 0)
                ret.area += valLeft.area
                ret.perimeter += valLeft.perimeter
                ret.lastPosition = valLeft.lastPosition
                ret.dir = valLeft.dir
                ret.sideList = valLeft.sideList
            }

            // Down.
            if pos.y < len(mapThing) - 1 && mapThing[pos.y + 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 1)
                copySideList = remove(copySideList, 2)
                valDown := getValues(mapThing, Coord{pos.y + 1, pos.x}, copySideList, start, 1)
                ret.area += valDown.area
                ret.perimeter += valDown.perimeter
                ret.lastPosition = valDown.lastPosition
                ret.dir = valDown.dir
                ret.sideList = valDown.sideList
            }

            break

        // Going right.
        case 3:

            // Down.
            if pos.y < len(mapThing) - 1 && mapThing[pos.y + 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 1)
                copySideList = remove(copySideList, 2)
                valDown := getValues(mapThing, Coord{pos.y + 1, pos.x}, copySideList, start, 1)
                ret.area += valDown.area
                ret.perimeter += valDown.perimeter
                ret.lastPosition = valDown.lastPosition
                ret.dir = valDown.dir
                ret.sideList = valDown.sideList
            }

            // Right.
            if pos.x < len(mapThing[0]) - 1 && mapThing[pos.y][pos.x + 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 3)
                copySideList = remove(copySideList, 0)
                valRight := getValues(mapThing, Coord{pos.y, pos.x + 1}, copySideList, start, 3)
                ret.area += valRight.area
                ret.perimeter += valRight.perimeter
                ret.lastPosition = valRight.lastPosition
                ret.dir = valRight.dir
                ret.sideList = valRight.sideList
            }

            // Up.
            if pos.y > 0 && mapThing[pos.y - 1][pos.x] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 2)
                copySideList = remove(copySideList, 1)
                valUp := getValues(mapThing, Coord{pos.y - 1, pos.x}, copySideList, start, 2)
                ret.area += valUp.area
                ret.perimeter += valUp.perimeter
                ret.lastPosition = valUp.lastPosition
                ret.dir = valUp.dir
                ret.sideList = valUp.sideList
            }

            // Left.
            if pos.x > 0 && mapThing[pos.y][pos.x - 1] == chr {
                copySideList := make([] int, len(sideList))
                copy(copySideList, sideList)
                copySideList = remove(copySideList, 0)
                copySideList = remove(copySideList, 3)
                valLeft := getValues(mapThing, Coord{pos.y, pos.x - 1}, copySideList, start, 0)
                ret.area += valLeft.area
                ret.perimeter += valLeft.perimeter
                ret.lastPosition = valLeft.lastPosition
                ret.dir = valLeft.dir
                ret.sideList = valLeft.sideList
            }

            break

        default:
            log.Fatal("Invalid direction.")
            break

    } // End switch.

    return ret

} // End func. 


func contains(arr [] int, val int) bool {

    for _, v := range arr {
        if v == val {
            return true
        }
    }

    return false
}


func remove(arr [] int, val int) [] int {

    result := make([] int, 0, len(arr))

    for _, v := range arr {
        if v != val {
            result = append(result, v)
        }
    }

    return result
}


