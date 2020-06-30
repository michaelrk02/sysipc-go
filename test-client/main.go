package main

import (
    "fmt"
    "github.com/michaelrk02/sysipc-go"
)

func main() {
    c := sysipc.NewClient(sysipc.LocalRouter, "arith")
    for {
        var a, b float64
        fmt.Printf("Enter two numbers: ")
        fmt.Scanf("%f%f", &a, &b)

        fmt.Printf("Calculating ...\n")

        args := make(map[string]interface{})
        args["lhs"] = a
        args["rhs"] = b

        var ret interface{}
        var err error

        fmt.Printf("Addition: ")
        ret, err = c.Call("add", args)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("%f\n", ret.(float64))

        fmt.Printf("Subtraction: ")
        ret, err = c.Call("sub", args)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("%f\n", ret.(float64))

        fmt.Printf("Multiplication: ")
        ret, err = c.Call("mul", args)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("%f\n", ret.(float64))

        fmt.Printf("Division: ")
        ret, err = c.Call("div", args)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("%f\n", ret.(float64))
    }
}

