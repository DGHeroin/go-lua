package main

import (
    "fmt"
    "net/http"
)

func NewHTTP(address string) error {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("Hello World HTTP Plugin!"))
    })
    go func() {
        if err := http.ListenAndServe(address, nil); err != nil {
            fmt.Println(err)
        }
    }()
    return nil
}
