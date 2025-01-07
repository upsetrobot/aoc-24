/**
 ******************************************************************************
 * Advent of Code 2024 - Day 16 Part 1
 *
 * Yeah, its still bad. Let's start over.
 *
 * Yeah, did it. Took way longer than it should of. I got get better with this 
 * kind of stuff, I guess.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        05 Jan 2025
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

type Dir int

const (
    Left    Dir = iota + 1
    Down
    Up
    Right
)

type Pos struct {
    x int
    y int
    dir Dir
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
    solution := 0
    var maze [][] byte
    
    for _, line := range strLines {
        maze = append(maze, []byte(line))
    }

    // Okay, let's try Dykstra's.
    type Node struct {
        x int
        y int
        char byte
        distFromSource int
        visited bool
    }

    type Edge struct {
        a *Node
        b *Node
        distBetween int
    }

    var end *Node
    nodeList := make([] *Node, 0)
    edgeList := make([] Edge, 0)

    // Make node list.
    for i, line := range maze {
        for j, pos := range line {
            switch pos {
                case 'S':
                    nodeList = append(nodeList, &Node{j, i, pos, 0, false})
                    break

                case 'E':
                    nodeList = append(nodeList, &Node{j, i, pos, math.MaxInt32, false})
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
                            nodeList = append(nodeList, &Node{j, i, pos, math.MaxInt32, false})
                        }
                    } else if num > 2 {
                        nodeList = append(nodeList, &Node{j, i, pos, math.MaxInt32, false})
                    }
                    break

                // Does not matter.
                default:
                    break

            } // End switch.

        } // End for.

    } // End for.

    // Make edge list.
    for i := 0; i < len(nodeList); i++ {
        count++

        srcX := nodeList[i].x
        srcY := nodeList[i].y

        for j := 0; j < len(nodeList); j++ {
            if i == j {
                continue
            }

            dstX := nodeList[j].x
            dstY := nodeList[j].y
            curX := srcX
            curY := srcY
            count := 0
            alreadyAdded := false

            for k := 0; k < len(edgeList); k++ {
                if edgeList[k].b.x == srcX && edgeList[k].b.y == srcY && edgeList[k].a.x == dstX && edgeList[k].a.y == dstY {
                    alreadyAdded = true
                    break
                }
            }

            if alreadyAdded {
                continue
            }
            
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
                        curX++
                        count++
                        char := maze[curY][curX]
                        
                        if char == '#' {
                            count = 0
                            break
                        }
                    }

                }

            } // End if.

            if count > 0 {
                count += 1000
                edgeList = append(edgeList, Edge{nodeList[i], nodeList[j], count})
            }

        } // End for.

    } // End for.

    // Now run Dykstras.
    var currNode *Node

    // Visit neighbors.
    for {
        count++ 

        flag := true
        min := math.MaxInt32

        for i := 0; i < len(nodeList); i++ {
            if !nodeList[i].visited && nodeList[i].distFromSource < min {
                min = nodeList[i].distFromSource
                currNode = nodeList[i]
                flag = false
                // break     <---- TOOK ME WAY TOO LONG TO FIGURE OUT THIS BAD.
            }
        }

        if flag {
            break
        }

        currNode.visited = true

        for i := 0; i < len(edgeList); i++ {
            edge := &edgeList[i]
            dist := currNode.distFromSource + edge.distBetween

            if edge.a == currNode {
                if !edge.b.visited && dist < edge.b.distFromSource {
                //if dist < edge.b.distFromSource {
                    edge.b.distFromSource = dist
                }

            } else if edge.b == currNode {
                if !edge.a.visited && dist < edge.a.distFromSource {
                //if dist < edge.a.distFromSource {
                    edge.a.distFromSource = dist
                }
            }
        }

    } // End for.

    for _, node := range nodeList {
        maze[node.y][node.x] = '+'
    }

    // Mark path.
    curr := end
    num := 0

    for {
        if curr.char == 'S' {
            break
        }

        num++
        fmt.Printf("Num: %d\n", num)
        fmt.Printf("Node x: %d\n", curr.x)
        fmt.Printf("Node y: %d\n", curr.y)
        fmt.Printf("Node dist: %d\n", curr.distFromSource)
        fmt.Println()

        maze[curr.y][curr.x] = 'o'

        min := curr.distFromSource
        next := curr

        for i := 0; i < len(edgeList); i++ {
            if edgeList[i].a == curr && edgeList[i].b.distFromSource < min {
                min = edgeList[i].b.distFromSource
                next = edgeList[i].b

            } else if edgeList[i].b == curr && edgeList[i].a.distFromSource < min {
                min = edgeList[i].a.distFromSource
                next = edgeList[i].a
            }
        }

        curr = next
    }

    // Print maze.
    for _, line := range maze {
        fmt.Println(string(line))
    }
    fmt.Println()
    
    // Find solution.
    solution = end.distFromSource

    // Print solution.
    fmt.Println("Day 16 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println("Count:", count)
    fmt.Println()

} // End main.


