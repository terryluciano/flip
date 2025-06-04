package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func randNum() int {
    num := rand.IntN(2)
    return num
}

func main() {

    help := flag.Bool("help", false, "Display help message")
    headsFlag := flag.String("heads", "Heads", "String to represent heads")
    tailsFlag := flag.String("tails", "Tails", "String to represent tails")
    flag.Parse()

    args := flag.Args()

    if *help {
        fmt.Println("Usage: flip [number_of_flips]")
        fmt.Println("       flip -help")
        fmt.Println("")
        fmt.Println("Arguments:")
        fmt.Println("  number_of_flips  Number of coin flips (default is 1)")
        fmt.Println("")
        fmt.Println("Options:")
        fmt.Println("  -help            Display this help message")
        return
    }

    if *headsFlag == "" || *tailsFlag == "" {
        fmt.Println("Error: Heads and Tails strings cannot be empty")
        os.Exit(1)
    }

    x := 1
    
    if len(args) > 0 {
        num, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Printf("Error: '%s' is not a valid number\n", args[0])
            os.Exit(1)
        }
        x = num
    }

    if x <= 0 {
        fmt.Println("Error: Number of flips cannot be zero")
        os.Exit(1)
    }

    if x > 1 {
        fmt.Println()
    }

    heads := 0
    tails := 0

    for i := 0; i < x; i++ {
        num := randNum()

        switch num {
            case 0:
                fmt.Println(*headsFlag)
                heads++
            case 1:
                fmt.Println(*tailsFlag)
                tails++
            default:
                fmt.Println("Error: Invalid number generated")
                os.Exit(1)
        }
    }

    if x == 1 {
        os.Exit(0)
    }

    fmt.Println()
    fmt.Println("Flipping completed.")
    fmt.Println()
    fmt.Println("============================")
    fmt.Println()
    fmt.Printf("%s: %d, %s: %d\n", *headsFlag, heads, *tailsFlag, tails)
    if heads > tails {
        fmt.Printf("%s wins!\n", strings.ToUpper(*headsFlag))
    } else if tails > heads {
        fmt.Printf("%s wins!\n", strings.ToUpper(*tailsFlag))
    } else {
        fmt.Println("It's a TIE!")
    }

    os.Exit(0)
}