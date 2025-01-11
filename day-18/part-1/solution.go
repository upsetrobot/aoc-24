/**
 ******************************************************************************
 * Advent of Code 2024 - Day 18 Part 1
 *
 * Great, another shortest path problem. Fortunately, after the last debacle, I
 * probably know how to do this without a stupid optimized recursive brute
 * force algorithm (I say stupid, because that is the trap I have repeated
 * fallen into). My plan is plot the bytes and then just run Dykstra's.
 *
 * Worked out the box.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        10 Jan 2025
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
	"strconv"
	"strings"
)

// Types.

type Coord struct {
    x int
    y int
}


type Node struct {
    x int
    y int
    dist int
    edges [] *Edge
    visited bool
}


type Edge struct {
    a *Node
    b *Node
    weight int
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
    coords := [] Coord {}
    
    for _, line := range strLines {
        nums := strings.Split(line, ",")
        x, err := strconv.Atoi(nums[0])
        if err != nil {
            log.Fatal("Conversion failure.")
        }

        y, err := strconv.Atoi(nums[1])
        if err != nil {
            log.Fatal("Conversion failure.")
        }

        coords = append(coords, Coord {x, y})
    }

    // Make map.
    //width := 7          // Example size.
    //height := 7         // Example size.
    //steps := 12         // Example steps.
    width := 71
    height := 71
    steps := 1024
    level := make([][] byte, height)

    for i := 0; i < len(level); i++ {
        level[i] = make([] byte, width)

        for j := 0; j < len(level[i]); j++ {
            level[i][j] = '.'
        }
    }

    // Plot bytes.
    step := 0

    for _, coord := range coords {
        level[coord.y][coord.x] = '#'
        step++

        if step == steps {
            break
        }
    }

    // Make graph.
    nodes := [] *Node {}
    var start *Node
    var end *Node

    for i := 0; i < len(level); i++ {
        for j := 0; j < len(level[i]); j++ {
            if level[i][j] == '.' {
                node := Node {j, i, math.MaxInt64, [] *Edge {}, false}
                nodes = append(nodes, &node)

                if i == 0 && j == 0 {
                    start = &node
                }

                if i == height - 1 && j == width - 1 {
                    end = &node
                }
            }
        }
    }

    // Connect nodes.
    for _, node := range nodes {
        
        // Check left.
        if node.x > 0 {
            if level[node.y][node.x - 1] == '.' {
                for _, b := range nodes {
                    if b.y == node.y && b.x == node.x - 1 {
                        node.edges = append(node.edges, &Edge {node, b, 1})
                        break
                    }
                }
            }
        }

        // Check down.
        if node.y < height - 1 {
            if level[node.y + 1][node.x] == '.' {
                for _, b := range nodes {
                    if b.y == node.y + 1 && b.x == node.x {
                        node.edges = append(node.edges, &Edge {node, b, 1})
                        break
                    }
                }
            }
        }

        // Check up.
        if node.y > 0 {
            if level[node.y - 1][node.x] == '.' {
                for _, b := range nodes {
                    if b.y == node.y - 1 && b.x == node.x {
                        node.edges = append(node.edges, &Edge {node, b, 1})
                        break
                    }
                }
            }
        }

        // Check right.
        if node.x < width - 1 {
            if level[node.y][node.x + 1] == '.' {
                for _, b := range nodes {
                    if b.y == node.y && b.x == node.x + 1 {
                        node.edges = append(node.edges, &Edge {node, b, 1})
                        break
                    }
                }
            }
        }

    } // End for.

    // Run dykstra's.
    curr := start
    curr.dist = 0

    for {
        min := math.MaxInt64
        flag := true

        for _, node := range nodes {
            if !node.visited && node.dist < min {
                min = node.dist
                curr = node
                flag = false
            }
        }

        if flag {
            break
        }

        curr.visited = true

        for _, edge := range curr.edges {
            if !edge.b.visited && curr.dist + edge.weight < edge.b.dist {
                edge.b.dist = curr.dist + edge.weight
            }
        }
    }

    // Mark path.
    curr = end

    for {
        level[curr.y][curr.x] = 'O'

        if curr == start {
            break
        }

        min := curr.dist
        next := curr

        for _, edge := range curr.edges {
            if edge.b.dist < min {
                min = edge.b.dist
                next = edge.b
            }
        }

        if next == curr {
            break
        }

        curr = next
    }

    // Print map.
    for _, line := range level {
        for _, pos := range line {
            fmt.Print(string(pos))
        }

        fmt.Println()
    }

    fmt.Println()

    solution = end.dist

    // Print solution.
    fmt.Println("Day 18 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


