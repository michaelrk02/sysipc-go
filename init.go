package sysipc

import (
    "math/rand"
    "time"
)

func init() {
    rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

    LocalRouter = NewRouter(".")
}

