package main

import (
  "fmt"
  "bufio"
  "strings"
  "net"
  "strconv"
  "sync"
)

var wg sync.WaitGroup

<<<<<<< HEAD
/** Fonction qui ouvre une connexion sur le port 10000 de la machine et qui va gérer tout les clients se connectant à ce serveur
* @param fromCollector, canal d'une place qui va contenir les entrées valide des clients
*/
func Collecter(fromCollector chan int){
  listener,err:=net.Listen("tcp",":10000") // Ouvre la connexion
  if err!=nil{
=======
func Collecter(fromCollector chan int) {
  listener, err := net.Listen("tcp",":10000")
  if err != nil {
>>>>>>> ac5115a7dca52e8c75f18038e6eb115835f67d16
    fmt.Println("Erreur de Listen")
  }

  for {
<<<<<<< HEAD

    connexionBase, err := listener.Accept() // Attend la connexion d'un client
    if err!=nil{
=======
    connexionBase, err := listener.Accept()
    if err != nil {
>>>>>>> ac5115a7dca52e8c75f18038e6eb115835f67d16
      fmt.Println("Erreur de Connexion")
    }
    wg.Add(1)
    go EcouterClient(connexionBase, fromCollector) // Lance une goRoutine dès qu'un client s'est connecté

  }

<<<<<<< HEAD
//Si le programme sort de la boucle infinie
    fmt.Println("ShutDown du serveur")
}

/* Fontion qui gère un et un seul client en l'écoutant et en mettant ses entrées valides dans un canal
* @param connexion, correspond au client
* @param fromCollector, canal d'une place qui va contenir les entrées valide du client
*/
func EcouterClient(connexion net.Conn, fromCollector chan int){
=======
  wg.Wait()
  fmt.Println("ShutDown du serveur")
}


func EcouterClient(connexion net.Conn, fromCollector chan int) {
>>>>>>> ac5115a7dca52e8c75f18038e6eb115835f67d16
  defer wg.Done()
  var msg int
  reader := bufio.NewReader(connexion) // On associe un reader au client
  for {
<<<<<<< HEAD
    message,err := reader.ReadString('\n') // Dès que ce client entre un message valide (càd qu'il entre un entier)
    if err == nil {
        msg, err = strconv.Atoi(strings.TrimRight(message, "\n")) // Parse le message de string à int en enlevant le retour chariot, cela donne un entier
        if err == nil {
          fromCollector <- msg // Le client attend que le canal fromCollector soit vide, dès qu'il est, y rentre la valeur de son message (en int)
        }
=======
    message, err := reader.ReadString('\n')
    if err == nil {
      msg, err = strconv.Atoi(strings.TrimRight(message, "\n"))
      if err == nil {
        fromCollector <- msg
>>>>>>> ac5115a7dca52e8c75f18038e6eb115835f67d16
      }
    }
  }
}
