package sysipc

import (
    "encoding/json"
    "os"
)

type fileDispatch struct {
    name, lockname string
}

func newFileDispatch(name string) *fileDispatch {
    fd := new(fileDispatch)
    fd.name = name
    fd.lockname = name + ".lock"
    return fd
}

func (fd *fileDispatch) exists() bool {
    var err error
    if _, err = os.Stat(fd.name); err != nil {
        return false
    }
    return true
}

func (fd *fileDispatch) locked() bool {
    var err error
    if _, err = os.Stat(fd.lockname); err != nil {
        return false
    }
    return true
}

func (fd *fileDispatch) lock() error {
    var err error

    for fd.locked() {
    }

    var f *os.File
    if f, err = os.Create(fd.lockname); err != nil {
        return err
    }
    defer f.Close()

    return nil
}

func (fd *fileDispatch) unlock() {
    if fd.locked() {
        os.Remove(fd.lockname)
    }
}

func (fd *fileDispatch) send(in interface{}, lock bool) error {
    var err error

    if lock {
        if err = fd.lock(); err != nil {
            return err
        }
        defer fd.unlock()
    }

    var f *os.File
    if f, err = os.Create(fd.name); err != nil {
        return err
    }
    defer f.Close()

    enc := json.NewEncoder(f)
    if err = enc.Encode(in); err != nil {
        return err
    }

    return nil
}

func (fd *fileDispatch) receive(out interface{}, lock bool) error {
    var err error

    if lock {
        if err = fd.lock(); err != nil {
            return err
        }
        defer fd.unlock()
    }

    var f *os.File
    if f, err = os.Open(fd.name); err != nil {
        return err
    }
    defer f.Close()

    dec := json.NewDecoder(f)
    if err = dec.Decode(out); err != nil {
        return err
    }

    return nil
}

func (fd *fileDispatch) remove(unlock bool) {
    os.Remove(fd.name)
    if unlock {
        fd.unlock()
    }
}

