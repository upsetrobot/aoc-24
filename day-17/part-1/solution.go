/**
 ******************************************************************************
 * Advent of Code 2024 - Day 17 Part 1
 *
 * Seems like a little computing emulation going on here. We can just create a
 * state and functions and run the program. Sounds easy.
 *
 *      Literal operands:   values 0 to 7 (3-bit) (I have no idea how to do
 *                          3-bit in go, but I can just probably use ints)
 *      Combo operands:     0 - 3 are literal, 4 rA, 5 rB, 6 rC.
 *      `adv` (0):          rA = rA / 2**(combo) (integral division).
 *      `bxl` (1):          rB = b ^ (literal).
 *      `bst` (2):          rB = (combo) % 8.
 *      `jnz` (3):          if rA == 0, nop; else jmp literal.
 *      `bxc` (4):          rB = rB ^ rC; operand is ignored.
 *      `out` (5):          print combo % 8. Output separated by commas.
 *      `bdv` (6):          rB = rA / 2^(combo) (integral division).
 *      `cdv` (7):          rC = rA / 2^(combo) (integral division).
 *
 * Worked right away. Nice.
 *
 * file:        solution.go
 * brief:       Solution for Advent of Code challenge in GoLang.
 * author:      upsetrobot
 * date:        09 Jan 2025
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

type State struct {
    a int
    b int
    c int
    ip int
    prog [] int
    out string
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
    words := [] string {}
    state := State {}
    
    for _, line := range strLines {
        words = strings.Split(line, " ")

        if len(words) != 0 {
            if words[0] == "Register" {
                switch words[1] {
                    case "A:":
                        state.a, err = strconv.Atoi(words[2])
                        if err != nil {
                            log.Fatal("Conversion error.")
                        }
                        break

                    case "B:":
                        state.b, err = strconv.Atoi(words[2])
                        if err != nil {
                            log.Fatal("Conversion error.")
                        }
                        break
                        
                    case "C:":
                        state.c, err = strconv.Atoi(words[2])
                        if err != nil {
                            log.Fatal("Conversion error.")
                        }
                        break

                    default:
                        log.Fatal("Parsing error.")
                }

            } else if words[0] == "Program:" {
                words = strings.Split(words[1], ",")

                for _, word := range words {
                    num, err := strconv.Atoi(word)
                    if err != nil {
                        log.Fatal("Program conversion error.")
                    }

                    state.prog = append(state.prog, num)
                }

            } // End if.

        } // End if.

    } // End for.

    // Run program.
    solution := ""

    for {
        if state.ip >= len(state.prog) {
            break
        }

        switch state.prog[state.ip] {
            case 0:
                adv(&state)
                break;

            case 1:
                bxl(&state)
                break;

            case 2:
                bst(&state)
                break;

            case 3:
                jnz(&state)
                break;

            case 4:
                bxc(&state)
                break;

            case 5:
                out(&state)
                break;

            case 6:
                bdv(&state)
                break;

            case 7:
                cdv(&state)
                break;

            default:
                log.Fatal("Invalid opcode.")
            
        } // End switch.

    } // End for.

    solution = state.out

    // Print solution.
    fmt.Println("Day 17 Part 1")
    fmt.Println("Filename:", fileName)
    fmt.Println()
    fmt.Println("Solution:", solution)
    fmt.Println()

} // End main.


/**
 * @brief   Returns the value of a combo operand for the current operation 
 *          based on the current instruction pointer of the given state. If 
 *          operand is 4, 5, or 6, then the value is set to the value from 
 *          registers A, B, or C respectively. Values greater than 6 will 
 *          result in a fatal error.
 * 
 * @param   state   Pointer to state of CPU.
 * 
 * @return  int     Value of combo operand.
 */
func combo(state *State) int {

    val := state.prog[state.ip + 1]

    if val > 3 {
        switch val {
            case 4:
                val = state.a
                break

            case 5:
                val = state.b
                break

            case 6:
                val = state.c
                break

            default:
                log.Fatal("Invalid operand.")
        }
    }

    return val

} // End combo.


/**
 * @brief   Divide value in register A by combo operand and store result in 
 *          register A. Division is integral division.
 *
 * @param   state   Pointer to state of CPU.
 */
func adv(state *State) {

    combo := combo(state)
    state.a = state.a / int(math.Pow(2, float64(combo)))
    state.ip += 2

} // End adv.


/**
 * @brief   Calculate the bitwise XOR of register B and the literal operand 
 *          and store the result in register B.
 *
 * @param   state   Pointer to state of CPU.
 */
func bxl(state *State) {

    state.b = state.b ^ state.prog[state.ip + 1]
    state.ip += 2

} // End bxl.


/**
 * @brief   Calculate the value of the combo operand modulo 8 and store the 
 *          result in register B.
 *
 * @param   state   Pointer to state of CPU.
 */
func bst(state *State) {

    combo := combo(state)
    state.b = combo % 8
    state.ip += 2

} // End bst.


/**
 * @brief   Do nothing if value in register A is 0. Otherwise jump to the 
 *          instruction index of the literal operand.
 *
 * @param   state   Pointer to state of CPU.
 */
func jnz(state *State) {

    if state.a != 0 {
        state.ip = state.prog[state.ip + 1]

    } else {
        state.ip += 2
    }

} // End jnz.


/**
 * @brief   Calculate the bitwise XOR of register B and register C and store
 *          the result in register B.
 *
 * @param   state   Pointer to state of CPU.
 */
func bxc(state *State) {

    state.b = state.b ^ state.c
    state.ip += 2

} // End bxc.


/**
 * @brief   Calculate the value of the combo operand modulo 8 and store the 
 *          result in the output stream.
 *
 * @param   state   Pointer to state of CPU.
 */
func out(state *State) {

    combo := combo(state) % 8

    if state.out == "" {
        state.out = strconv.Itoa(combo)

    } else {
        state.out = state.out[:] + "," + strconv.Itoa(combo)
    }

    state.ip += 2

} // End out.


/**
 * @brief   Divide value in register A by combo operand and store result in 
 *          register B. Division is integral division.
 *
 * @param   state   Pointer to state of CPU.
 */
func bdv(state *State) {

    combo := combo(state)
    state.b = state.a / int(math.Pow(2, float64(combo)))
    state.ip += 2

} // End bdv.


/**
 * @brief   Divides value in register A by combo operand and store result in 
 *          register C. Division is integral division.
 *
 * @param   state   Pointer to state of CPU.
 */
func cdv(state *State) {

    combo := combo(state)
    state.c = state.a / int(math.Pow(2, float64(combo)))
    state.ip += 2

} // End adv.



