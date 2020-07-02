package main

import (
    "github.com/michaelrk02/sysipc-go"
    "os"
    "os/signal"
)

var s *sysipc.Server
var sigCh chan os.Signal

func main() {
    s = sysipc.NewServer(sysipc.LocalRouter, "arith", os.Stdout)
    s.Handle("add", addHandler)
    s.Handle("sub", subHandler)
    s.Handle("mul", mulHandler)
    s.Handle("div", divHandler)

    sigCh = make(chan os.Signal, 1)
    go wait()
    signal.Notify(sigCh, os.Interrupt)

    s.Run()
}

func wait() {
    <-sigCh
    s.Stop()
}
