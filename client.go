package sysipc

import (
    "errors"
    "path"
)

type Client struct {
    r *Router
    name string

    fd, req, res *fileDispatch
}

func NewClient(r *Router, name string) *Client {
    c := new(Client)

    c.r = r
    c.name = name

    c.fd = newFileDispatch(c.ServerAddress())
    c.req = newFileDispatch(c.ServerAddress() + ".request")
    c.res = newFileDispatch(c.ServerAddress() + ".response")

    return c
}

func (c *Client) Router() *Router {
    return c.r
}

func (c *Client) ServerName() string {
    return c.name
}

func (c *Client) ServerAddress() string {
    return path.Join(c.r.name, c.name)
}

func (c *Client) Call(method string, args map[string]interface{}) (interface{}, error) {
    var err error

    c.fd.lock()
    defer c.fd.unlock()

    req := new(request)
    req.CallId = uint64(rnd.Int63())
    req.Method = method
    req.Args = args
    if err = c.req.send(req, true); err != nil {
        return nil, err
    }

    for c.fd.locked() && !c.res.exists() {
    }
    if !c.fd.locked() {
        return nil, errors.New("server was unlocked by another process")
    }

    res := new(response)
    if err = func() error {
        c.res.lock()
        defer c.res.unlock()

        if err = c.res.receive(res, false); err != nil {
            return err
        }

        if req.CallId != res.CallId {
            return errors.New("server responded with different call ID")
        }

        c.res.remove(false)

        return nil
    }(); err != nil {
        return nil, err
    }

    err = nil
    if res.Err != "" {
        err = errors.New(res.Err)
    }
    return res.Ret, err
}

