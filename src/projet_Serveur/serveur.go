package main

var workersNumber int = 5
type Task struct {
  loopNumber int //dans notre cas actuel, la tache à effectuer est une nombre de tour de boucle
}

/** Fonction principale initialisant le serveur
*/
func main() {
  fromCollector := make(chan Task) // Canal permettant de communiquer les taches à effectuer
  availableWorkers := make(chan chan Task, workersNumber) // Canal permettant de stocker les canaux servant à communiquer les taches à chaque travailleur
  for i := 0; i < workersNumber; i++ {
    workChans := make(chan Task) // Canal servant à communiquer les taches à chaque travailleur
    go Travailler(workChans, availableWorkers)
    availableWorkers <- workChans
  }
  go Repartir(fromCollector, availableWorkers)
  go Collecter(fromCollector)
  for {

  }
}
