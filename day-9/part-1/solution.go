/**
 ******************************************************************************
 * Advent of Code 2024 - Day 9 Part 1
 *
 * This seems pretty simple. We have to use a "map" of alternating sizes of
 * files and sizes of freespace (single digit for both I think) to generate
 * another "map" representation of a disk and free space from which we can
 * move portions of files by ID to free contiguous space on the end. Then,
 * we have to iterate the disk and calculate sum of the products of block IDs
 * and file ID (I assume that there are only single digit file IDs).
 *
 * I think a good approach is to do it as it done in the example, but there is
 * probably a better mathematic solution. Maybe subtract available space
 * number from the last file ID... and remove it... and keep subtracting till
 * the file size is 0, then move right to the next file. But the files don't
 * have an explicit ID (rather they are indexed). ... I think I want to create
 * a struct to represent each file and put those in an array, then fill
 * everything up, and then calculate the checksum. ... Yeah, like we would
 * have to have the number of blocks so we can know the block ID of the next
 * block. Yeah, its basically the same, but without the ascii. Let's try that.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        17 Dec 2024
 * copyright:   2024. All rights reserved.
 *
 ******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"os"
)

// Object to represent file on disk.
type FileBlock struct {
    free bool
    id int
    sz int
    // I think that's all we really need. Block IDs and checksums can be 
    // calculated at the end from this info.
}


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
    solution := 0
    fileID := 0
    var fileList [] FileBlock

    for i, chr := range file {

        // Convert to decimal
        num := int(chr - 48)

        // Check for end of file.
        if num > 9 {
            break
        }

        // Flip-flop.
        if i % 2 == 0 {
            fileList = append(fileList, FileBlock{false, fileID, num})
            fileID++

        } else {
            fileList = append(fileList, FileBlock{true, 0, num})
        }

    } // End for.

    // Fill empty blocks.
    for i, j := 0, len(fileList) - 1; i < j; i, j = i + 1, j - 1 {
        for !fileList[i].free && i < len(fileList) {
            i++
        }

        for fileList[j].free && j >= 0 {
            j--
        }

        fileList[i].free = false
        fileList[i].id = fileList[j].id
        szDiff := fileList[i].sz - fileList[j].sz

        if szDiff > 0 {
            fileList[i].sz = fileList[j].sz
            fileList[j].free = true

            // Insert new free node.
            newSlice := make([] FileBlock, len(fileList), len(fileList) + 1)
            copy(newSlice, fileList[:i + 1])
            newSlice[i + 1] = FileBlock{true, 0, szDiff}
            copy(newSlice[i + 2:], fileList[i + 1:])
            fileList = newSlice
            
        } else if szDiff < 0 {
            fileList[j].sz -= fileList[i].sz
            j++

        } else {
            fileList[j].free = true
        }

    } // End for.

    // Calculate checksum.
    count := 0

    for _, blk := range fileList {
        if blk.free == true {
            break
        }

        for i := 0; i < blk.sz; i++ {
            solution += count * blk.id
            count++
        }
    }

    // Print solution.
    fmt.Println("Day 9 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


