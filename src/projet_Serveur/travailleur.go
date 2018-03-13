package main

import (
  "fmt"
  "strconv"
)
// "time"

func Travailler(workChan chan int, availaibleWorkers chan chan int) {
  for {
    var taskNumber int = <-workChan
    for i := 0; i < taskNumber; i++ {
      fmt.Println(strconv.Itoa(i))
      // time.Sleep(1000000)
    }
    fmt.Println("Travailleur : J'ai travaillé "+strconv.Itoa(taskNumber)+" fois")
    availaibleWorkers <- workChan
  }
}
