/**
 ******************************************************************************
 * Advent of Code 2024 - Day 16 Part 2
 *
 * Had to clean up and start again....again. so close.
 * Still had wrong answer. It looks great, so I don't know.
 * Next thing is to try to make turns into nodes maybe or maybe a directed 
 * graph. I'm pretty sure there is a right hand from start path not accounted 
 * for. That's crazy. That's maybe that's the issue... maybe.
 *
 * Yep, kinda obvious. Yeah, that was stupid. Took me way to long. Oh well.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        06 Jan 2025
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
	"strings"
)


// Types.
type Node struct {
    x int
    y int
    scoreFromSource int
    scoreFromEnd int
    distFromSource int
    distFromEnd int
    visited bool
    visitedEnd bool
    edgeList [] Edge
    onMainPath bool
}

type Edge struct {
    a *Node
    b *Node
    scoreBetween int
    distBetween int
}


// Constants.
// None.


// Variables.

var count = 0


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

    // Make maze.
    var maze [][] byte
    
    for _, line := range strLines {
        maze = append(maze, []byte(line))
    }

    // Make node list.
    var end *Node
    var start *Node
    nodeList := make([] *Node, 0)

    for i, line := range maze {
        for j, pos := range line {
            switch pos {
                case 'S':
                    nodeList = append(nodeList, &Node{
                        j, 
                        i, 
                        0, 
                        math.MaxInt32, 
                        0, 
                        math.MaxInt32, 
                        false, 
                        false,
                        [] Edge{},
                        true,
                    })
                    start = nodeList[len(nodeList) - 1]
                    break

                case 'E':
                    nodeList = append(nodeList, &Node{
                        j, 
                        i, 
                        math.MaxInt32, 
                        0, 
                        math.MaxInt32, 
                        0, 
                        false, 
                        false,
                        [] Edge{},
                        true,
                    })
                    end = nodeList[len(nodeList) - 1]
                    break

                case '.':
                    left := maze[i][j - 1]
                    down := maze[i + 1][j]
                    up := maze[i - 1][j]
                    right := maze[i][j + 1]
                    leftB := false
                    downB := false
                    upB := false
                    rightB := false
                    num := 0

                    if left == '.' {
                        leftB = true
                        num++
                    }

                    if down == '.' {
                        downB = true
                        num++
                    }

                    if up == '.' {
                        upB = true
                        num++
                    }

                    if right == '.' {
                        rightB = true
                        num++
                    }

                    if num == 2 {
                        if !(leftB && rightB) && !(downB && upB) {
                            nodeList = append(nodeList, &Node{
                                j, 
                                i, 
                                math.MaxInt32, 
                                math.MaxInt32, 
                                math.MaxInt32, 
                                math.MaxInt32, 
                                false, 
                                false,
                                [] Edge{},
                                false,
                            })
                        }
                    } else if num > 2 {
                        nodeList = append(nodeList, &Node{
                            j, 
                            i, 
                            math.MaxInt32, 
                            math.MaxInt32, 
                            math.MaxInt32, 
                            math.MaxInt32, 
                            false, 
                            false,
                            [] Edge{},
                            false,
                        })
                    }
                    break

                // Does not matter.
                default:
                    break

            } // End switch.

        } // End for.

    } // End for.

    // Make edges.
    for i := 0; i < len(nodeList); i++ {
        srcX := nodeList[i].x
        srcY := nodeList[i].y

        for j := 0; j < len(nodeList); j++ {

            // How to make turns into nodes.
            if i == j {
                continue
            }

            dstX := nodeList[j].x
            dstY := nodeList[j].y
            curX := srcX
            curY := srcY
            count := 0
            score := 0
            mark := false

            if srcX == dstX {
                if srcY > dstY {
                    for curY != dstY {
                        curY--
                        count++
                        char := maze[curY][curX]
                        
                        if char == '#' {
                            count = 0
                            break
                        }

                    }

                } else {
                    for curY != dstY {
                        curY++
                        count++
                        char := maze[curY][curX]
                        
                        if char == '#' {
                            count = 0
                            break
                        }
                    }

                }

            } else if srcY == dstY {
                if srcX > dstX {
                    for curX != dstX {
                        curX--
                        count++
                        char := maze[curY][curX]
                        
                        if char == '#' {
                            count = 0
                            break
                        }
                    }

                } else {
                    for curX != dstX {
                        char := maze[curY][curX]

                        if char == 'S' {
                            mark = true
                        }

                        curX++
                        count++

                        if char == '#' {
                            count = 0
                            break
                        }
                    }
                }

            } // End if.

            if count > 0 {
                score = count + 1000

                // Added this for the start node.
                if mark {
                    score -= 1000
                }

                nodeList[i].edgeList = append(
                    nodeList[i].edgeList, 
                    Edge {nodeList[i], nodeList[j], score, count},
                )
            }

        } // End for.

    } // End for.

    // Now run Dykstras.
    var curr *Node
    cycles := 0

    for {
        cycles++
        flag := true
        min := math.MaxInt32

        for i := 0; i < len(nodeList); i++ {
            if !nodeList[i].visited && nodeList[i].scoreFromSource < min {
                min = nodeList[i].scoreFromSource
                curr = nodeList[i]
                flag = false
            }
        }

        if flag {
            break
        }

        curr.visited = true

        for i := 0; i < len(curr.edgeList); i++ {
            score := curr.scoreFromSource + curr.edgeList[i].scoreBetween
            dist := curr.distFromSource + curr.edgeList[i].distBetween

            if !curr.edgeList[i].b.visited && 
                score < curr.edgeList[i].b.scoreFromSource {

                curr.edgeList[i].b.scoreFromSource = score
                curr.edgeList[i].b.distFromSource = dist
            }
        }

    } // End for.

    // Now run Dykstras again (from end to every node instead of from start).
    for {
        cycles++
        flag := true
        min := math.MaxInt32

        for i := 0; i < len(nodeList); i++ {
            if !nodeList[i].visitedEnd && nodeList[i].scoreFromEnd < min {
                min = nodeList[i].scoreFromEnd
                curr = nodeList[i]
                flag = false
            }
        }

        if flag {
            break
        }

        curr.visitedEnd = true

        for i := 0; i < len(curr.edgeList); i++ {
            score := curr.scoreFromEnd + curr.edgeList[i].scoreBetween
            dist := curr.distFromEnd + curr.edgeList[i].distBetween

            if !curr.edgeList[i].b.visitedEnd && 
                score < curr.edgeList[i].b.scoreFromEnd {

                curr.edgeList[i].b.scoreFromEnd = score
                curr.edgeList[i].b.distFromEnd = dist
            }
        }

    } // End for.

    pathCount := 0
    nodeCount := 0
    counter := byte(0)
    pathCount++

    // Now draw all paths that are equal to the shortest path.
    for i := 0; i < len(nodeList); i++ {

        if nodeList[i].scoreFromSource + nodeList[i].scoreFromEnd ==
            end.scoreFromSource {

            nodeCount++

            curr = nodeList[i]
            next := curr

            // Draw path to source. Which one? Idk.
            for {

                if curr == start {
                    break
                }
                
                for _, edge := range curr.edgeList {
                    if edge.b.scoreFromEnd + edge.b.scoreFromSource == end.scoreFromSource {
                        if edge.b.scoreFromSource < curr.scoreFromSource {
                            next = edge.b
                        }
                    }
                }

                x := curr.x
                y := curr.y
                dx := next.x
                dy := next.y

                if byte('A' + counter) > 'Z' {
                    counter = 0
                }

                if x == dx {
                    if y > dy {
                        for y >= dy {
                            maze[y][x] = 'A' + counter
                            y--
                        }

                    } else {
                        for y <= dy {
                            maze[y][x] = 'A' + counter
                            y++
                        }
                    }

                } else {
                    if x > dx {
                        for x >= dx {
                            maze[y][x] = 'A' + counter
                            x--
                        }

                    } else {
                        for x <= dx {
                            maze[y][x] = 'A' + counter
                            x++
                        }
                    }
                }

                curr = next

            } // End for.

            // Draw path to end. Which one? Idk.
            curr = nodeList[i]
            next = curr

            for {

                if curr == end {
                    break
                }
                
                for _, edge := range curr.edgeList {
                    if edge.b.scoreFromEnd + edge.b.scoreFromSource == end.scoreFromSource {
                        if edge.b.scoreFromEnd < curr.scoreFromEnd {
                            next = edge.b
                        }
                    }
                }
                
                x := curr.x
                y := curr.y
                dx := next.x
                dy := next.y

                if x == dx {
                    if y > dy {
                        for y >= dy {
                            maze[y][x] = 'A' + counter
                            y--
                        }

                    } else {
                        for y <= dy {
                            maze[y][x] = 'A' + counter
                            y++
                        }
                    }

                } else {
                    if x > dx {
                        for x >= dx {
                            maze[y][x] = 'A' + counter
                            x--
                        }

                    } else {
                        for x <= dx {
                            maze[y][x] = 'A' + counter
                            x++
                        }
                    }
                }

                curr = next

            } // End for.

            counter++

        } // End if.

    } // End for.

    // Print maze and find solution.
    solution := 0

    for _, line := range maze {
        fmt.Println(string(line))

        for _, pos := range line {
            if pos > '.' {
                solution++
            }
        }
    }
    fmt.Println()
    
    // Print solution.
    fmt.Println("Day 16 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println("Cycles:", cycles)
    fmt.Println("Main Path Length:", end.distFromSource)
    fmt.Println("Main Path Score:", end.scoreFromSource)
    fmt.Println("Paths Drawn Count (Number of joins in all drawn paths):", nodeCount)
    fmt.Println()

} // End main.



