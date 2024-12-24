/**
 ******************************************************************************
 * Advent of Code 2024 - Day 11 Part 1
 *
 * Okay, this one is a little weird, but simple. You have a list of numbers.
 * You just have to iterate a given number of times (25) and follow some rules 
 * for each value in the list. Seems really simple before starting it.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        24 Dec 2024
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
type Node struct {
    val int
    numDigits int
    str string
    next* Node
}

type LinkedList struct {
    head* Node
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
    numList := NewLinkedList()
    strNums := strings.Split(string(file), "\n")
    strNums = strings.Split(strNums[0], " ")
    solution := 0
    
    for _, strNum := range strNums {
        num, err := strconv.Atoi(strNum)
        if err != nil {
            log.Fatal("Failed conversion")
        }

        numList.Append(num, len(strNum), strNum)
    }

    // Process list.
    // In C, I would do this with a linked list. So, how do I do something 
    // like that in go?
    for i := 0; i < 25; i++ {
        node := numList.head
        skip := false

        for node != nil {

            // Skip if split.
            if skip {
                skip = false

            // 0 becomes 1.
            } else if node.val == 0 {
                node.val = 1
                node.str = "1"

            // Even makes two stones.
            } else if node.numDigits % 2 == 0 {
                right := node.str[node.numDigits / 2:]
                pos := 0

                // String leading zeros.
                for j, chr := range right {
                    if j == len(right) - 1 {
                        break

                    } else if chr == '0' {
                        pos += 1

                    } else {
                        break
                    }
                }

                if pos > 0 {
                    right = right[pos:]
                }

                num, err := strconv.Atoi(right)
                if err != nil {
                    log.Fatal("Conversion failed")
                }

                newNode := &Node{num, len(right), right, node.next}
                node.str = node.str[0:node.numDigits / 2]

                // String leading zeros.
                pos = 0

                for j, chr := range node.str {
                    if j == len(right) - 1 {
                        break

                    } else if chr == '0' {
                        pos += 1

                    } else {
                        break
                    }
                }

                if pos > 0 {
                    node.str = node.str[pos:]
                }

                node.val, err = strconv.Atoi(node.str)
                if err != nil {
                    log.Fatal("Conversion failed")
                }

                node.numDigits = len(node.str)
                node.next = newNode
                skip = true

            // Mulitply.
            } else {
                node.val *= 2024
                node.str = strconv.Itoa(node.val)
                node.numDigits = len(node.str)

            } // End if.

            node = node.next

        } // End for.

    }// End for.

    // Print solution.
    fmt.Println("Day 11 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    node := numList.head

    for node != nil {
        fmt.Println(node.val)
        solution += 1
        node = node.next
    }

    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


/**
 * Create a new list.
 */
func NewLinkedList() *LinkedList {
    return &LinkedList{head: nil}
}



/**
 * Append an int to the given linked list.
 */
// What kind of notation is this?
func (l *LinkedList) Append(val int, numDigits int, str string) {

    newNode := &Node{val: val, numDigits: numDigits, str: str, next: nil}

    if l.head == nil {
        l.head = newNode

    } else {
        current := l.head

        for current.next != nil {
            current = current.next
        }

        current.next = newNode
    }
}


