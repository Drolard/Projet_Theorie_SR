package main

import (
  "fmt"
)

func Travailler(workChan chan int, availaibleWorkers chan chan int) {
  for  {
    var taskNumber int = <-workChan
    for i := 0; i < taskNumber; i++ {

    }
    fmt.Println("Travailleur : J'ai travaillé "+taskNumber+" fois")
    availaibleWorkers <- workChan
  }
}
