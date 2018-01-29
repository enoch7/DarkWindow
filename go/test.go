package main
 
import (
    "log"
 
    gc "github.com/rthornton128/goncurses"
)
 
func main() {
    s, err := gc.Init()
    if err != nil {
        log.Fatal("init:", err)
    }
    defer gc.End()
    s.Move(5, 2)
    s.Println("Hello")
    s.GetChar()
}