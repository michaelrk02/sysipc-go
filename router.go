package sysipc

import (
    "path"
)

type Router struct {
    name string
}

func NewRouter(name string) *Router {
    r := new(Router)

    r.name = path.Clean(name)

    return r
}

func (r *Router) Name() string {
    return r.name
}

