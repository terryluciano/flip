package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func randNum() int {
    num := rand.IntN(2)
    return num
}

func main() {

    help := flag.Bool("help", false, "Display help message")
    flag.Parse()

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

    x := 1

    args := flag.Args()
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
                fmt.Println("Heads")
                heads++
            case 1:
                fmt.Println("Tails")
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
    fmt.Printf("Heads: %d, Tails: %d\n", heads, tails)
    if heads > tails {
        fmt.Println("HEADS wins!")
    } else if tails > heads {
        fmt.Println("TAILS wins!")
    } else {
        fmt.Println("It's a TIE!")
    }

    os.Exit(0)
}