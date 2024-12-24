/**
 ******************************************************************************
 * Advent of Code 2024 - Day 11 Part 2
 *
 * Okay, this one is a little weird, but simple. You have a list of numbers.
 * You just have to iterate a given number of times (25) and follow some rules 
 * for each value in the list. Seems really simple before starting it.
 *
 * Part 2 is easy. They just want to do it 75 times. Let's see if that affects 
 * execution time.
 *
 * Yeah, it was the hardest for me so far. Had to use dynamic programming.
 * The table is what it did it.
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
type Item struct {
    iter int
    id int
    val int
}


// Constants.
// None.


// Variables.
var gItemArr [] Item


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
    var numList [] int
    strNums := strings.Split(string(file), "\n")
    strNums = strings.Split(strNums[0], " ")
    solution := 0
    
    // Changed numList to int array.
    for _, strNum := range strNums {
        num, err := strconv.Atoi(strNum)
        if err != nil {
            log.Fatal("Failed conversion")
        }

        numList = append(numList, num)
    }

    // Process list.
    // In C, I would do this with a linked list. So, how do I do something 
    // like that in go?

    // WOW. That takes a long time and I cut it off when my system hit 54 GB 
    // of memory. Definitely need a better solution.

    // I'm wondering if I can count without storing info. 
    // Or if I can just save like one byte or something. But I don't think 
    // that's possible. 

    // Or maybe there is a mathematical solution... maybe.
    // Wow, nice, this is hard. The num of digits isn't easily predictable 
    // without the number.

    // Changing one of the roots by one affects the result a lot. I feel like 
    // using some kind of mathematic solution is possible. It is like a 
    // geometric progression or something or somekind of expansion, 
    // intuitively, at least. 

    // Maybe each number contributes a solution individually. So like 
    // sol(125, 17) = sol(125) + sol(17). Yeah, that might help, because 
    // then we don't have to store everything in memory, just enough to get 
    // one solution at a time (with with 75 iteration can still be alot). 
    // ... Okay, this seems to be true. So, let's try that.

    // Okay, it makes sense, but even one number expansion takes 54 GB.
    // 25 iterations is fine. There must be a crazy exponential gain in 
    // numbers during expansion obviously. So if that is the root of the 
    // problem, we have to address that. ... Maybe we can use a recursive 
    // solution that is returning the count of each number without 
    // maintaining a list. This makes sense and may solve the memory 
    // problem, so let's try that.

    // I did not finish this idea, because it seems the same. I could 
    // precalculate solutions and store them in a list and then use those 
    // solutions. But I think I will still have problems with a memoization 
    // solution. It would help of course, but I still have to calculate it 
    // once... Actually, I think this would probably help. So, how do I do 
    // that... Lemme think, we need to go through each num, if num is in 
    // table, add that to solution, if not, do a round of expansion, then 
    // return count of that list, ... the expansion part is the key... 
    // During expansion, we generate a list, then to get the next expansion, 
    // we have to look up each value again and expand. So back to the recurse 
    // solution. vvvv

    for _, num := range numList {
        solution += recurseCount(num, 75)

    } // End for.

    // Print solution.
    fmt.Println("Day 11 Part 2")
    fmt.Println("Filename:", fileName)
    fmt.Println()

    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


/**
 * Recursively calculate count of stones based on given value.
 */
func recurseCount(val int, iterations int) int {
    
    // Need to figure out exit criteria. It will have something to do with 
    // the number of iterations which changes... uh... 
    // It's almost like we need to return a list... Or need to accept a list.
    // Which is bad... I have a feeling like this is logically equivalent to 
    // the iterative solution... including memory. Because I need the list to 
    // find the solution for the number.
    // So idea right now is: sol(x) = listAfterOneIter[], for i in list: 
    // sol += sol(i) which is not right. Or: sol(x, i) = y ; 25 expansions 
    // length which is the same as the iterative solution. What is sol(0)

    // Okay, back here. If value is in table (which is another problem), then
    // we return that. If not, we generate an expansion, during expansion, 
    // we intend to do all iterations, but the iterations change, so we need 
    // to account for the number. We do one expansion and return count + 
    // recurse (iter - 1). When the iteration gets to one (or 0)... Oh, 0 
    // would be one (no expansion is one number). And one would be count of 
    // one expansion. Ugh... Okay, we take in val and iterations.
    
    if iterations == 0 {

        return 1
    }

    // 2. Now we check table.

    // Not in table?
    // 3. Now we do one expansion.

    // Now we have an expansion list. So we have a count. 
    // 4. If iter == 1 , we can return count.
    // And we have to save val:count in table.
    
    // If not, we have go through the list and return 
    // recurse iter - 1 for each val so we need a sum.

    // Then we return the sum. Wow what a recursive-iterative monster. 
    // Let's try this.

    // Ok, first I have to figure out a good hashtable solution in go. 
    // Also, I don't think we will ever hit iter 0... to do that, we would 
    // remove step 4. So, I guess we can skip that.
    // For the table, I think I'm just gonna use an array. Screw a hashtable.
    // So, I will use a global struct array (bad choice?).

    // Check table.
    for _, item := range gItemArr {
        if iterations == item.iter && val == item.id {
            
            return item.val
        }
    }

    // Make expansion list.
    var numList [] int

    if val == 0 {

        numList = append(numList, 1)

    } else {
        strNum := strconv.Itoa(val)

        if len(strNum) % 2 == 0 {
            left := strNum[0:len(strNum) / 2]
            right := strNum[len(strNum) / 2:]

            leftNum, err := strconv.Atoi(left)
            if err != nil {
                log.Fatal("Conversion failed")
            }

            rightNum, err := strconv.Atoi(right)
            if err != nil {
                log.Fatal("Conversion failed")
            }

            numList = append(numList, leftNum)
            numList = append(numList, rightNum)

        } else {
            numList = append(numList, val * 2024)
        
        } // End if.

    } // End if.

    // Now I have expansion list.
    sum := 0

    for _, num := range numList {
        sum += recurseCount(num, iterations - 1)
    }

    // Add sum to table.
    gItemArr = append(gItemArr, Item{iterations, val, sum})

    return sum

} // End func.


