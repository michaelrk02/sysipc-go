package main

import (
    "github.com/michaelrk02/sysipc-go"
)

func main() {
    s := sysipc.NewServer(sysipc.LocalRouter, "arith")
    s.Handle("add", addHandler)
    s.Handle("sub", subHandler)
    s.Handle("mul", mulHandler)
    s.Handle("div", divHandler)
    s.Run()
}

