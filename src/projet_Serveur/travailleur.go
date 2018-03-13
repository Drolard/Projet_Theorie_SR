package main

import (
  "fmt"
  "strconv"
)

/** Dès qu'une tache se trouve dans le canal workChan, cette fonction l'effectue
* @param workChan, canal contenant la tache à effectuer
* @param availableWorkers, canal contenant les canaux des travailleurs disponibles
*/
func Travailler(workChan chan Task, availaibleWorkers chan chan Task) {
  for {
    task := <-workChan
    for i := 0; i < task.loopNumber; i++ {
      // fmt.Println(i)
    }
    fmt.Println("Travailleur : J'ai travaillé "+strconv.Itoa(task.loopNumber)+" fois")
    availaibleWorkers <- workChan // Ce travailleur se rend disponible après avoir effectuer la tâche
  }
}
