package main

import (
  "fmt"
  "strconv"
)
// "time"

func Travailler(workChan chan int, availaibleWorkers chan chan int) {
  for {
    var taskNumber int = <-workChan
    fmt.Println(taskNumber)
    for i := 0; i < taskNumber; i++ {
      fmt.Println(i)
      // time.Sleep(1)
    }
    fmt.Println("Travailleur : J'ai travaillé "+strconv.Itoa(taskNumber)+" fois")
    availaibleWorkers <- workChan
  }
}
