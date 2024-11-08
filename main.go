package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

var (
    students = make(map[int]Student)
    nextID   = 1
    mu       sync.Mutex // For thread-safe access to `students`
)
