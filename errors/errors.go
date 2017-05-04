package errors

import "fmt"

type Error struct {
  Code int,
  Message string
}

const (
  someError Error = iota
  anotherError
)
