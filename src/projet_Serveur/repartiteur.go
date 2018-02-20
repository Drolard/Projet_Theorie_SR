package main

import (
  "fmt"
)

func Repartir(fromCollector chan int, availaibleWorkers chan chan int) {
  for {
    int task = <-fromCollector
    chan int nextWorker = <-availaibleWorkers
    nextWorker <- task
    fmt.Println("Repartiteur : J'ai réparti une tâche portant le numéro "+task)
  }
}
