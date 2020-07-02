package sysipc

import (
    "errors"
    "io"
    "log"
    "path"
    "sync"
)

type Handler func(map[string]interface{}) (interface{}, error)

type Server struct {
    r *Router
    name string

    l *log.Logger
    handlers map[string]Handler

    fd, req, res *fileDispatch

    running bool
    mu sync.Mutex
}

func NewServer(r *Router, name string, logOutput io.Writer) *Server {
    s := new(Server)

    s.r = r
    s.name = name

    if logOutput != nil {
        s.l = log.New(logOutput, "SysIPC [" + s.Address() + "] ", log.LstdFlags)
    }
    s.handlers = make(map[string]Handler)

    s.fd = newFileDispatch(s.Address())
    s.req = newFileDispatch(s.Address() + ".request")
    s.res = newFileDispatch(s.Address() + ".response")

    s.running = false

    return s
}

func (s *Server) Router() *Router {
    return s.r
}

func (s *Server) Name() string {
    return s.name
}

func (s *Server) Address() string {
    return path.Join(s.r.name, s.name)
}

func (s *Server) Handle(method string, h Handler) error {
    var ok bool

    if _, ok = s.handlers[method]; ok {
        return errors.New("method `" + method + "` is already handled by another routine")
    }

    s.handlers[method] = h

    return nil
}

func (s *Server) Running() bool {
    s.mu.Lock()
    running := s.running
    s.mu.Unlock()
    return running
}

func (s *Server) Run() {
    s.mu.Lock()
    s.running = true
    s.mu.Unlock()

    s.fd.unlock()
    s.req.remove(true)
    s.res.remove(true)
    for s.Running() {
        var err error
        if err = s.intercept(); err != nil {
            if s.l != nil {
                s.l.Println(err)
            }
        }
    }
}

func (s *Server) Stop() {
    s.mu.Lock()
    s.running = false
    s.mu.Unlock()
}

func (s *Server) intercept() error {
    var err error
    var ok bool

    for s.Running() && !s.req.exists() {
    }
    if !s.Running() {
        return errors.New("server stopped")
    }

    res := new(response)
    defer func() {
        if err != nil {
            res.Err = err.Error()
        }
        if err = s.res.send(res, true); err != nil {
            return
        }
    }()

    req := new(request)
    if err = func() error {
        s.req.lock()
        defer s.req.remove(true)

        return s.req.receive(req, false)
    }(); err != nil {
        return err
    }
    defer func() {
        res.CallId = req.CallId
    }()

    var h Handler
    if h, ok = s.handlers[req.Method]; !ok {
        err = errors.New("no handler is associated for method `" + req.Method + "`")
        return err
    }
    res.Ret, err = h(req.Args)

    return err
}

