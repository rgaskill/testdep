package main

import "fmt"
import "github.com/rgaskill/ignoreme/something"
import "github.com/rgaskill/ignoreme/nothing"

func main() {
	something.Run()
	nothing.Nothing()
    fmt.Println("hello world")
}