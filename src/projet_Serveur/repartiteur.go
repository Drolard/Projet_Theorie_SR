package main

import (
  "fmt"
  "strconv"
)

func Repartir(fromCollector chan int, availaibleWorkers chan chan int) {
  for {
    task := <-fromCollector
    nextWorker := <-availaibleWorkers
    nextWorker <- task
    fmt.Println("Repartiteur : J'ai réparti une tâche portant le numéro "+ strconv.Itoa(task))
  }
}
