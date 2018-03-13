package main

var workersNumber int = 5

/** Fonction principale initialisant le serveur
*/
func main() {
  fromCollector := make(chan int) //canal permettant de communiquer les taches à effectuer
  availableWorkers := make(chan chan int, workersNumber) //canal permettant de stocker les canaux servant à communiquer les taches à chaque travailleur
  for i := 0; i < workersNumber; i++ {
    workChans := make(chan int) //canal servant à communiquer les taches à chaque travailleur
    go Travailler(workChans, availableWorkers)
    availableWorkers <- workChans
  }
  go Repartir(fromCollector, availableWorkers)
  go Collecter(fromCollector)
  for {

  }
}
