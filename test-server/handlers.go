package main

import (
    "errors"
    "time"
)

func checkArgs(args map[string]interface{}) error {
    if _, ok := args["lhs"]; !ok {
        return errors.New("expecting 'lhs' argument")
    }
    if _, ok := args["rhs"]; !ok {
        return errors.New("expecting 'rhs' argument")
    }
    d, _ := time.ParseDuration("2s")
    time.Sleep(d)
    return nil
}

func addHandler(args map[string]interface{}) (interface{}, error) {
    if err := checkArgs(args); err != nil {
        return nil, err
    }
    lhs := args["lhs"].(float64)
    rhs := args["rhs"].(float64)
    result := lhs + rhs
    return result, nil
}

func subHandler(args map[string]interface{}) (interface{}, error) {
    if err := checkArgs(args); err != nil {
        return nil, err
    }
    lhs := args["lhs"].(float64)
    rhs := args["rhs"].(float64)
    result := lhs - rhs
    return result, nil
}

func mulHandler(args map[string]interface{}) (interface{}, error) {
    if err := checkArgs(args); err != nil {
        return nil, err
    }
    lhs := args["lhs"].(float64)
    rhs := args["rhs"].(float64)
    result := lhs * rhs
    return result, nil
}

func divHandler(args map[string]interface{}) (interface{}, error) {
    if err := checkArgs(args); err != nil {
        return nil, err
    }
    lhs := args["lhs"].(float64)
    rhs := args["rhs"].(float64)
    if rhs == 0.0 {
        return nil, errors.New("attempted to divide by zero")
    }
    result := lhs / rhs
    return result, nil
}

