package main

import (
  "fmt" //nom court pour bibliothèque standard
  "bufio"
  "strings"
  "net"
  "io"
  "strconv"
  "sync"
)

var workersNumber int = 5

func main() {
  fromCollector := make(chan int)
  availableWorkers := make(chan chan int, workersNumber)
  for i := 0; i < workersNumber; i++ {
    workChans := make(chan int)
    go Travailler(workChans, availableWorkers)
    availableWorkers <- workChans
  }
  go Repartir(fromCollector, availableWorkers)
  go Collecter(fromCollector)
}
