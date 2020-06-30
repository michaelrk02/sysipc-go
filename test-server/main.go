package main

import (
    "github.com/michaelrk02/sysipc-go"
    "os"
)

func main() {
    s := sysipc.NewServer(sysipc.LocalRouter, "arith", os.Stdout)
    s.Handle("add", addHandler)
    s.Handle("sub", subHandler)
    s.Handle("mul", mulHandler)
    s.Handle("div", divHandler)
    s.Run()
}

