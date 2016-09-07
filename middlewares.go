package main

import (
  "log"
  "net/http"
  "runtime"
)

func recovery() func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  return func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    defer func() {
      if err := recover(); err != nil {
        stack := make([]byte, 1024*8)
        stack = stack[:runtime.Stack(stack, true)]
        log.Printf("PANIC: %s\n", err)
        log.Println(string(stack))
        body := []byte("500 Internal Server Error")
        res.WriteHeader(http.StatusInternalServerError)
        res.Write(body)
      }
    }()
    next(res, req)
  }
}

