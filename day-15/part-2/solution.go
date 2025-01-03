/**
 ******************************************************************************
 * Advent of Code 2024 - Day 15 Part 2
 *
 * This one seems pretty simple. We just have to move boxes around. I think 
 * part 2 is probably where this get hard. There a few approaches. I plan on 
 * taking the naive approach and just move characters around.
 *
 * Ugh, now the boxes are twice as wide, but the robot is still one wide. I 
 * want to just use the same algorithm but I have to tie the boxes together 
 * somehow. Yeah, this is just stupid! The boxes now have different alignments 
 * and can push other boxes they are only half-aligned to. The instructions 
 * want you to generate a new map from the first.
 * 
 * Okay, how to do this... man, I knew I should have gone for a mathematic 
 * solution looking at it. Maybe each point can have two halves, but then how 
 * do I track the state on two different positions? I really do not want to 
 * re-write my solution or extend it to account for boxes that can span
 * multiple rows and columns... Well, I guess it doesn't matter what side of
 * the box is where... wait, it does because they are connected... ugh.
 * Honestly, this one pisses me off and I have not even attempted it yet. I
 * could make two types of Os (like a right O and a left O) then run the same
 * algorithm with exception that left and right Os would move. But I think that
 * would be the same problem of actually doubling the map just without the
 * doubling. In which case, I might as well double the map, I think. I don't
 * know, I think I can probably do it with doubling the map and just have each
 * point have two characters as a state (which is basically doubling the map
 * but without actually doing it). It might make debugging harder, but it would
 * save significant parts of my code... Yeah, I guess I have no choice. Its
 * either that, actually expand the map and adjust the box movements, or come
 * up with some sort of graphical/mathematic solution which I do not want to
 * do. It shouldn't be too bad. I guess I will give it a shot. Actually, I will
 * probably just do it the stupid way and double the map and then account for
 * every stupid box (I really want to move a box with one movement, not
 * depending on which side of the box). I think its probably the best way.
 *
 * Yeah, I definitely overthought this one. I made much harder than it was by 
 * continually having the mindset of moving to the end instead of one step -- I 
 * made the same the mistake in part 1... So stupid. Anyway, it's not that 
 * hard if you figure out recording the positions of affected boxes. My
 * solution was definitely overly complex. Oh, well.. at least it's over.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        03 Jan 2025
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
        boxMap[i] = make([] byte, columns * 2)
    }

    robx *= 2

    // Parse file.
    directionIndex := 0

    for i, line := range strLines {
        if line == "" {
            directionIndex = i + 1
            break
        } 

        for j, pos := range line {
            switch pos {
                case '@':
                    boxMap[i][j * 2] = byte(pos)
                    boxMap[i][j * 2 + 1] = '.'
                    break

                case 'O':
                    boxMap[i][j * 2] = '['
                    boxMap[i][j * 2 + 1] = ']' //stupid.
                    break

                default:
                    boxMap[i][j * 2] = byte(pos)
                    boxMap[i][j * 2 + 1] = byte(pos)
            }
        }
    }

    // Print.
    //for _, line := range boxMap {
    //    for _, pos := range line {
    //        fmt.Print(string(pos))
    //    }

    //    fmt.Print("\n")
    //}

    fmt.Print("\n")

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

                        case '[':
                            fallthrough

                        case ']':

                            // Need to count boxes left.
                            count := 0
                            k := robx

                            for ; k > 0; k-- {
                                // I got a feeling, I'm gonna need this in part
                                // 2.
                                if boxMap[roby][k] == '#' || boxMap[roby][k] == '.' {
                                    break
                                }

                                if boxMap[roby][k] == '[' || boxMap[roby][k] == ']' {
                                    count++
                                }
                            }

                            if boxMap[roby][k] == '#' {
                                k++
                            }

                            // Make adjustments.
                            flag := true

                            for count > 0 {
                                if flag {
                                    boxMap[roby][k] = '['

                                } else {
                                    boxMap[roby][k] = ']'
                                }

                                flag = !flag
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

                // Up and down is where it gets stupid.
                // We may be able to use some sort of recursive function.
                // Or add positions to a list.
                case 'v':
                    switch boxMap[roby + 1][robx] {
                        case '#':
                           break

                        case '.':
                            boxMap[roby][robx] = '.'
                            boxMap[roby + 1][robx] = '@'
                            roby++

                        case '[':
                            fallthrough     // Took forever to figure out this 
                                            // was a problem.

                        case ']':

                            // Okay, I need 2 pointers to go up. Then more if 
                            // they push other boxes. So, if we have a list 
                            // of positions to move up, we can keep sweeping.
                            // But we have to note the configuration somehow.
                            // So... sweeper list for the check. Actually, 
                            // that's not pefect either (cuz what about 
                            // splits). What about a matrix of positions...
                            // like it grows in rows and columns based on 
                            // checks, then we can iterate on row to do the 
                            // stopping checks...yeah, that will probably work.

                            // Okay, now we need a position reference and to 
                            // start detection.
                            // Actually, I think we need a sweep and a matrix.

                            // Actually, no. I laid down and determined that 
                            // the way to do this is to record all positions 
                            // of congruent boxes first (and bracket character 
                            // probably); save that; then move one forward; 
                            // then add or subtract one from the position list 
                            // ys. then check if those positions have a block;
                            // then if they do, undo the add/subtract; then 
                            // replace original positions with dots and 
                            // replace the new positions with brackets.
                            k := roby

                            type Box struct {
                                x int
                                y int
                                char byte
                            }
                            
                            // Make list of boxes.
                            boxList := [] Box {{robx, roby, '@'}}
                            xList := [] int {robx}

                            // Scan for all boxes.
                            for ; k < len(boxMap); k++ {
                                var newXList [] int
                            
                                // Scan next row for more boxes.
                                for _, x := range xList {
                                    if boxMap[k + 1][x] == '[' {
                                        newXList = append(newXList, x)
                                        newXList = append(newXList, x + 1)
                                        boxList = append(boxList, Box{x, k + 1, '['})
                                        boxList = append(boxList, Box{x + 1, k + 1, ']'})

                                    } else if  boxMap[k + 1][x] == ']' {
                                        newXList = append(newXList, x)
                                        newXList = append(newXList, x - 1)
                                        boxList = append(boxList, Box{x - 1, k + 1, '['})
                                        boxList = append(boxList, Box{x, k + 1, ']'})
                                    } 
                                }

                                if len(newXList) == 0 {
                                    break
                                }

                                xList = newXList[:]
                                
                            } // End for.

                            // Okay, now start moving forward.
                            // So, stupid. I accident wrote an algorith that 
                            // slides the end... again - instead of going one 
                            // step. Let's correct that.
                            found := false

                            for _, box := range boxList {
                                if boxMap[box.y + 1][box.x] == '#' {
                                    found = true
                                    break
                                } 
                            }

                            if found {
                                break
                            }

                            // Now we found how far. Replace original boxes.
                            for _, box := range boxList {
                                boxMap[box.y][box.x] = '.'
                            }

                            // Now move boxes.
                            for _, box := range boxList {
                                boxMap[box.y + 1][box.x] = box.char
                            }

                            roby++

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

                        case '[':
                            fallthrough

                        case ']':
                            k := roby

                            type Box struct {
                                x int
                                y int
                                char byte
                            }
                            
                            // Make list of boxes.
                            boxList := [] Box {{robx, roby, '@'}}
                            xList := [] int {robx}

                            // Scan for all boxes.
                            for ; k > 0; k-- {
                                var newXList [] int
                            
                                // Scan next row for more boxes.
                                for _, x := range xList {
                                    if boxMap[k - 1][x] == '[' {
                                        newXList = append(newXList, x)
                                        newXList = append(newXList, x + 1)
                                        boxList = append(boxList, Box{x, k - 1, '['})
                                        boxList = append(boxList, Box{x + 1, k - 1, ']'})

                                    } else if  boxMap[k - 1][x] == ']' {
                                        newXList = append(newXList, x)
                                        newXList = append(newXList, x - 1)
                                        boxList = append(boxList, Box{x - 1, k - 1, '['})
                                        boxList = append(boxList, Box{x, k - 1, ']'})
                                    } 
                                }

                                if len(newXList) == 0 {
                                    break
                                }

                                xList = newXList[:]
                                
                            } // End for.

                            found := false

                            for _, box := range boxList {
                                if boxMap[box.y - 1][box.x] == '#' {
                                    found = true
                                    break
                                } 
                            }

                            if found {
                                break
                            }

                            // Now we found how far. Replace original boxes.
                            for _, box := range boxList {
                                boxMap[box.y][box.x] = '.'
                            }

                            // Now move boxes.
                            for _, box := range boxList {
                                boxMap[box.y - 1][box.x] = box.char
                            }

                            roby--

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

                        case '[':
                            fallthrough

                        case ']':

                            // Need to count boxes left.
                            count := 0
                            k := robx

                            for ; k < len(boxMap[0]); k++ {
                                if boxMap[roby][k] == '#' || boxMap[roby][k] == '.' {
                                    break
                                }

                                if boxMap[roby][k] == '[' || boxMap[roby][k] == ']' {
                                    count++
                                }
                            }

                            if boxMap[roby][k] == '#' {
                                k--
                            }

                            // Make adjustments.
                            flag := true

                            for count > 0 {
                                if flag {
                                    boxMap[roby][k] = ']'

                                } else {
                                    boxMap[roby][k] = '['
                                }

                                flag = !flag
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
            if pos == '[' {
                solution += 100 * i 
                solution += j
            }
        }
    }

    // Print solution.
    fmt.Println("Day 15 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


