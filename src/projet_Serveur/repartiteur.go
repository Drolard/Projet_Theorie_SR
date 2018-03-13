package main

import (
  "fmt"
  "strconv"
)

/** Fonction permettant de répartir les tâches à effectuer avec les travailleurs
* @param fromCollector, canal contenant les tâches à répartir
* @param availaibleWorkers, canal contenant les canaux permettant de communiquer une tache à un travailleur disponible
*/
func Repartir(fromCollector chan Task, availaibleWorkers chan chan Task) {
  for {
    task := <-fromCollector // Bloquantloquant si aucune tâche n'est à effectuer
    nextWorker := <-availaibleWorkers // Bloquant si aucun travailleur n'est disponible
    nextWorker <- task
    fmt.Println("Repartiteur : J'ai réparti une tâche portant le numéro "+ strconv.Itoa(task.loopNumber))
  }
}
