package github.com/risou/modtest-hello

import (
    "errors"
    "fmt"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("modtest-hello.Hello need non-empty string as name")
    }
    return fmt.Sprintf("You inputted name `%s`", name), nil
}
