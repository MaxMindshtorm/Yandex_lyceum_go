package main

import (
    "time"
)

func NextWorkday(start time.Time) time.Time{
    return start + 
}
func main() {
    t1 := time.Now()
    time.Sleep(5 * time.Second)
    t2 := time.Now()
    fmt.Println(TimeDifference(t1, t2))
}