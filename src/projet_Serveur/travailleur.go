package main

import (
  "fmt"
  "strconv"
)

/** Dès qu'une tache se trouve dans le canal workChan, cette fonction l'effectue
* @param workChan, canal contenant la tache à effectuer
* @param availableWorkers, canal contenant les canaux des travailleurs disponibles
*/
func Travailler(workChan chan int, availaibleWorkers chan chan int) {
  for {
<<<<<<< HEAD
    var taskNumber int = <-workChan
    fmt.Println(taskNumber)
    for i := 0; i < taskNumber; i++ {
      fmt.Println(i)
      // time.Sleep(1)
=======
    var taskNumber int = <-workChan //bloquant si aucune tâche n'est communiqué à ce travailleur
    for i := 0; i < taskNumber; i++ {
      // fmt.Println(strconv.Itoa(i))
>>>>>>> ac5115a7dca52e8c75f18038e6eb115835f67d16
    }
    fmt.Println("Travailleur : J'ai travaillé "+strconv.Itoa(taskNumber)+" fois")
    availaibleWorkers <- workChan // ce travailleur se rend disponible après avoir effectuer la tâche
  }
}
